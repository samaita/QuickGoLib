package main

import (
	"flag"
	"log"
	"strings"
	"time"

	"github.com/samaita/QuickGoLib/file"
	"github.com/samaita/QuickGoLib/sort/merge"
)

func main() {
	var input, output string
	var arr []int64

	flag.StringVar(&input, "input", "usia.txt", "Input File, txt only")
	flag.StringVar(&output, "output", "usia_urut.txt", "Output File")

	flag.Parse()

	t := time.Now()

	// dummy validation of format file
	if !strings.Contains(input, ".txt") {
		log.Fatal("Invalid file format type, .txt only")
	}

	arr = file.ReadFileToArrayInt64(input)

	result := merge.SortV2(arr)

	file.WriteArrayInt64ToFile(output, result)

	log.Fatal("Successfully Done <3 Time Elapsed: ", time.Since(t))
}
