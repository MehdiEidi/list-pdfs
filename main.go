package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func getDrives(wg *sync.WaitGroup) []string {
	d := make([]string, 0)

	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			d = append(d, string(drive)+":/")
			wg.Add(1)
			f.Close()
		}
	}

	return d
}

func main() {
	start := time.Now()
	fmt.Println("Started...")

	var wg sync.WaitGroup

	drives := getDrives(&wg)

	f, err := os.Create("list.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, d := range drives {
		go func(drive string) {
			err := filepath.Walk(drive, func(path string, info fs.FileInfo, err error) error {
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

			wg.Done()
		}(d)
	}

	wg.Wait()
	fmt.Println("Finished.")
	duration := time.Since(start)
	fmt.Println("Elapsed Time:", duration)
}
