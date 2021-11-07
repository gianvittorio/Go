package main

import (
	"bufio"
	"fmt"
	//"go/scanner"
	"io"
	"io/ioutil"
	"log"
	"os"
	//"time"
)

func main() {

	// Open a new file for read only
	f1, err := os.Open("test1.txt")
	PrintFatalError(err)
	defer f1.Close()

	// Create a new file
	f2, err := os.Create("test2.txt")
	PrintFatalError(err)
	defer f2.Close()

	// Open file for read write
	f3, err := os.OpenFile("test3.txt", os.O_APPEND|os.O_RDWR, 0666)
	PrintFatalError(err)
	defer f3.Close()

	// Rename a file
	err = os.Rename("test1.txt", "test1New.txt")
	PrintFatalError(err)

	// Move a file
	err = os.Rename("./test1.txt", "./testfolder/test1.txt")
	PrintFatalError(err)

	// Copy a file
	CopyFile("test3.txt", "./testfolder/test3.txt")

	// Delete a file
	err = os.Remove("test2.txt")
	PrintFatalError(err)

	bytes, err := ioutil.ReadFile("test3.txt")
	fmt.Println(string(bytes))

	scanner := bufio.NewScanner(f3)
	count := 0
	for scanner.Scan() {
		count++
		fmt.Println("Found line: ", count, scanner.Text())
	}

	// Buffered write, efficient store in memory, saves disk I/O
	writeBuffer := bufio.NewWriter(f3)
	for i := 1; i <= 5; i++ {
		writeBuffer.WriteString(fmt.Sprintln("Added line", i))
	}
	writeBuffer.Flush()

	// GenerateFileStatusReport("test3.txt")

	// fileStat1, err := os.Stat("test3.txt")
	// PrintFatalError(err)
	// for {
		
	// }
}

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error happened while processing file", err)
	}
}

func CopyFile(fname1, fname2 string) {
	fold, err := os.Open(fname1)
	PrintFatalError(err)
	defer fold.Close()

	fNew, err := os.Create(fname2)
	PrintFatalError(err)
	defer fNew.Close()

	// Copy bytes from source to destination
	_, err = io.Copy(fNew, fold)
	PrintFatalError(err)

	// Flush file contents to desc
	err = fNew.Sync()
	PrintFatalError(err)
}
