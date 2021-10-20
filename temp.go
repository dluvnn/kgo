package kgo

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// TempMan ...
type TempMan struct {
	TPL *template.Template
}

func loadTemplate(tpl *template.Template, filename, id string) error {
	if !FileExists(filename) {
		return nil
	}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	t := tpl.New(id)
	s := fmt.Sprintf("{{ define \"%s\"}}", id) + string(b) + "{{ end }}"
	_, err = t.Parse(s)
	return err
}

// Build ...
func (tm *TempMan) Build(tempPath string) error {
	files := []string{}
	if DirExists(tempPath) {
		err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if path == tempPath {
				return nil
			}
			rel, err := filepath.Rel(tempPath, path)
			if err != nil {
				return err
			}
			if strings.ToLower(filepath.Ext(rel)) == ".html" {
				id := filepath.Join(tempPath, rel)
				files = append(files, id)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	temp := template.New("").Funcs(FMap)
	for _, file := range files {
		elen := len(filepath.Ext(file))
		s, err := filepath.Rel(tempPath, file)
		if err != nil {
			return err
		}
		id := strings.ReplaceAll(s[:len(s)-elen], "\\", "/")
		err = loadTemplate(temp, file, id)
		if err != nil {
			return err
		}
	}
	tm.TPL = temp
	return nil
}

// HasTemplate ...
func (tm *TempMan) HasTemplate(name string) bool {
	return tm.TPL.Lookup(name) != nil
}

// Render ...
func (tm *TempMan) Render(w io.Writer, name string, data interface{}) error {
	return tm.TPL.ExecuteTemplate(w, name, data)
}

// ToString ...
func (tm *TempMan) ToString(name string, data interface{}) (string, error) {
	var tpl bytes.Buffer
	err := tm.TPL.ExecuteTemplate(&tpl, name, data)
	if err != nil {
		return "", err
	}
	return tpl.String(), nil
}
