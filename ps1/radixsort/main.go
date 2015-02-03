package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var inputfile = flag.String("i", "", "The data input file to read from")

func getdata(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		file.Close() /// \todo handle error
	}()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // scan the header line
	scanner.Scan() // scan the length of the file
	len, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}

	data := make([]int, len, len)

	for i := 0; scanner.Scan(); i++ {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		data[i] = val
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

type assignments_t int

/// This is a subset of a true radix sort, for input values from [0..99].
/// A true radix sort for arbitrary integers is significantly slower, having to check many more digits.
func restrictedRadixSort(data []int) []int, assignments_t {
	buckets := make([]int, 100, 100)
	var assignments assignments_t

	for _, val := range data {
		buckets[val]++
	}

	pos := 0
	for value, numberOfElements := range buckets {
		for i := 0; i != numberOfElements; i++ {
			data[pos] = value
			pos++
		}
	}

	return data
}

func main() {
	flag.Parse()

	if *inputfile == "" {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		return
	}

	data, err := getdata(*inputfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)

	fmt.Println("sorting...")
	data = restrictedRadixSort(data)

	fmt.Println(data)
}
