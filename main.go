package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func getDrives() []string {
	d := make([]string, 0)

	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			d = append(d, string(drive)+":/")
			f.Close()
		}
	}

	return d
}

func main() {
	drives := getDrives()

	f, err := os.Create("list.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, d := range drives {
		err := filepath.Walk(d, func(path string, info fs.FileInfo, err error) error {
			if !info.IsDir() {
				if ext := filepath.Ext(path); ext == ".pdf" {
					_, err = f.WriteString(path + "\n")
					if err != nil {
						log.Fatal(err)
					}
				}
			}

			return nil
		})

		if err != nil {
			log.Println(err)
		}
	}
}
