package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func open(name string) (*os.File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f, err
}

func readFile(name string) ([]string, error) {
	f, err := open(name)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(f)
	log.Printf("%v", f)
	var lines []string
	for {
		line, err := reader.ReadString('\n')
		log.Printf("%v", line)

		if err != io.EOF {
			break
		}

		if err != nil {
			log.Printf("%v", err)
			return nil, err
		}

		lines = append(lines, string(line))
	}
	return lines, nil
}

func scan() {
	f, err := open("RP.fhx")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func util() {
	content, err := ioutil.ReadFile("RP.fhx")
	if err != nil {
		log.Println(err)
		return
	}
	str := strings.Split(string(content), "\n")
	fmt.Printf("%v\n", str)
}

func stream() {
	f, err := os.Open("RP.fhx")
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	defer f.Close()
	if err != nil {
		log.Printf("%v\n", err)
	}

	const BufferSize = 128

	b1 := make([]byte, BufferSize)
	for {
		n, err := f.Read(b1)

		if err != nil {
			if err != io.EOF {
				log.Printf("%v\n", err)
			}
			break
		}

		fmt.Println("Bytes Read:", n)
		fmt.Println("Content", string(b1[:n]))
	}

}

func scanLine() {
	f, err := os.Open("RP.fhx")
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// for _, each_ln := range text {
	// 	fmt.Println(each_ln)
	// }

	fmt.Printf("%v\n", len(text))
}

func main() {
	// util()
	// stream()
	scanLine()
}
