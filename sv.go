package kgo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LocalRedirect redirects current path to new path
func LocalRedirect(w http.ResponseWriter, r *http.Request, newPath string) {
	if q := r.URL.RawQuery; q != "" {
		newPath += "?" + q
	}
	w.Header().Set("Location", newPath)
	w.WriteHeader(http.StatusMovedPermanently)
}

func RedirectHandler(newPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, newPath, http.StatusSeeOther)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request, newPath string) {
	http.Redirect(w, r, newPath, http.StatusSeeOther)
}

func ReadBody(r *http.Request) ([]byte, error) {
	data, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	return data, err
}

// JSONParse parses the request body as JSON object
func JSONParse(r *http.Request, x interface{}) error {
	data, err := ReadBody(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, x)
}

// JSONReply writes reponse with the body is JSON of variable x
func JSONReply(w http.ResponseWriter, x interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(x)
}

func GetJSON(url string, out interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("server respond status code " + strconv.Itoa(resp.StatusCode))
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.Unmarshal(data, out)
}

// PostJson sends a POST request and return error if the status code is not equal 200
func PostJSON(url string, x interface{}) error {
	data, err := json.Marshal(x)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("server respond status code " + strconv.Itoa(resp.StatusCode))
	}
	return nil
}

// func NotFound(w http.ResponseWriter, r *http.Request, reason ...interface{}) {
// 	http.NotFound(w, r)
// }

func AccessDenied(w http.ResponseWriter, r *http.Request, reason ...interface{}) {
	http.Error(w, "access denied", http.StatusForbidden)
}

func ExtractIP(address string) string {
	host, _, err := net.SplitHostPort(address)
	if err != nil {
		return ""
	}
	if host == "::1" || host == "localhost" {
		return "127.0.0.1"
	}
	return host
}

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return ExtractIP(forwarded)
	}
	return ExtractIP(r.RemoteAddr)
}

var slist []*http.Server

// Wait ...
func Wait(shutdownFunc ...func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("\rtry to shutdown server ...")
	for _, s := range slist {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		s.Shutdown(ctx)
	}
	for _, f := range shutdownFunc {
		f()
	}

	fmt.Println("bye bye!")
	os.Exit(0)
}

// Run ...
func Run(addr string, handler http.Handler) {
	server := new(http.Server)
	server.Addr = addr
	if server.Addr == "" {
		panic("lack backend field")
	}
	server.Handler = handler
	fmt.Println("start listen on", addr)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	// wait(server)
	slist = append(slist, server)
}

// GetObjectID ...
func GetObjectID(x url.Values) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(x.Get("_id"))
}
