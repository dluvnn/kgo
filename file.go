package kgo

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

// UnmarshalJSONFile ...
func UnmarshalJSONFile(fname string, x interface{}) error {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, x)
}

// FileExists ...
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirExists ...
func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// RemoveComment remove comment from the io.Reader
var (
	removeCommentRegex = regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")
)

// RemoveJSONComment ...
func RemoveJSONComment(s string) string {
	return removeCommentRegex.ReplaceAllString(s, "")
}

// CopyFile ...
func CopyFile(src, dst string) (int64, error) {
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// ReadTextFile ...
func ReadTextFile(fname string) (string, error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return "", err
	}
	return string(data), err
}
