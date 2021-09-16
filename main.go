package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func getDrives() []string {
	d := make([]string, 0)

	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			d = append(d, string(drive))
			f.Close()
		}
	}

	return d
}

func main() {
	drives := []string{"g:/Books/"}

	for _, d := range drives {
		// start walk from each d
		// for each pdf, add its name to list.txt
		err := filepath.Walk(d, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
			}

			if !info.IsDir() {
				if ext := filepath.Ext(path); ext == ".pdf" {
					fmt.Println(filepath.Base(path))
				}

			}

			return nil
		})

		if err != nil {
			fmt.Println(err)
		}
	}
}
