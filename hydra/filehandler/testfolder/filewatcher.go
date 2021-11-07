package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	WatchFile("testfile.txt")
}

func WatchFile(fname string) {
	fileStat1, err := os.Stat(fname)
	PrintFatalError(err)
	for {
		time.Sleep(1 * time.Second)
		fileStat2, err := os.Stat(fname)
		PrintFatalError(err)
		if fileStat1.ModTime() != fileStat2.ModTime() {
			fmt.Println("File was modified at", fileStat2.ModTime())
			fileStat1, err = os.Stat(fname)
			PrintFatalError(err)
		}
	}
}

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error happened while processing file", err)
	}
}
