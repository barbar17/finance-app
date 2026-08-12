package main

import (
	"html/template"
	"os"
	"path/filepath"
)

func loadWebPages() (*template.Template, error) {
	tmpl := template.New("")

	err := filepath.Walk("web/pages", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Ext(path) != ".html" {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		name, err := filepath.Rel("web/pages", path)
		if err != nil {
			return err
		}

		name = filepath.ToSlash(name)

		_, err = tmpl.New(name).Parse(string(content))
		return err
	})

	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
