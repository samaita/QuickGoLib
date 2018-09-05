package file

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileToArrayInt64(filepath string) []int64 {
	var arr []int64

	// read file
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// validate empty space
		if strings.TrimSpace(scanner.Text()) == "" {
			// skip empty space
			continue
		}

		// validate if line is a number
		v, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatalln("File corrupted, contain:", scanner.Text())
		}

		// assign as array for input
		arr = append(arr, v)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return arr
}

func WriteArrayInt64ToFile(filepath string, arr []int64) error {

	// create file
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range arr {
		fmt.Fprintln(w, line)
	}

	return w.Flush()
}
