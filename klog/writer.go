package klog

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"time"
)

// RespondLogger ...
type RespondLogger interface {
	Print(*http.Request, time.Time)
	Header() http.Header
	Write(data []byte) (int, error)
	WriteHeader(status int)
}

// LogWriter ...
type LogWriter struct {
	W      http.ResponseWriter
	sz     int
	Status int
}

// Print ...
func (w *LogWriter) Print(r *http.Request, tstart time.Time) {
	if w.Status == http.StatusSwitchingProtocols {
		return
	}
	log.Printf(
		"%.6fs %03d %s %4s %8.1fKB %s\n",
		time.Since(tstart).Seconds(),
		w.Status,
		r.Header.Get("CF-Connecting-IP"),
		r.Method,
		float64(w.sz)/1024.0,
		r.RequestURI,
	)
}

// Header ...
func (w *LogWriter) Header() http.Header {
	return w.W.Header()
}

// Write ...
func (w *LogWriter) Write(data []byte) (int, error) {
	s, err := w.W.Write(data)
	w.sz = w.sz + s
	return s, err
}

// WriteHeader ...
func (w *LogWriter) WriteHeader(status int) {
	w.Status = status
	w.W.WriteHeader(status)
}

// HijackLogger ...
type HijackLogger struct {
	LogWriter
}

// Hijack ...
func (hl *HijackLogger) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h := hl.LogWriter.W.(http.Hijacker)
	conn, rw, err := h.Hijack()
	if err == nil && hl.LogWriter.Status == 200 {
		// The status will be StatusSwitchingProtocols if there was no error and
		// WriteHeader has not been called yet
		hl.LogWriter.Status = http.StatusSwitchingProtocols
	}
	return conn, rw, err
}
