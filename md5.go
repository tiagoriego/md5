package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	switch os.Args[1] {
	case "-f":
		f, err := os.Open(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			err := f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)

		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if lines != nil {
			h := md5.New()
			line := strings.Join(lines, "\n")
			io.WriteString(h, line)
			fmt.Printf("%x\n", h.Sum(nil))
		}

	case "-t":
		h := md5.New()
		io.WriteString(h, os.Args[2])
		fmt.Printf("%x\n", h.Sum(nil))
	}
}
