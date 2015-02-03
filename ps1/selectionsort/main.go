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

type assignments_t uint64
type comparisons_t uint64

func selectionsort(data []int) ([]int, assignments_t, comparisons_t) {
	var assignments assignments_t
	var comparisons comparisons_t

	/// DO NOT replace this with the XOR method. It is slow.
	swap := func(a *int, b *int) {
		temp := *a
		*a = *b
		*b = temp
	}

	// i is the current position, before which all elements are sorted
	for i := 0; i != len(data); i++ {
		nextLowest := i ///< the index of the lowest value in the array, which needs to be 'selected' and swapped with data[i]
		// j is the position of the iterator, which finds the next lowest value in the array
		for j := i; j != len(data); j++ {
			comparisons++
			if data[j] < data[nextLowest] {
				nextLowest = j
			}
		}

		assignments += 2
		swap(&data[i], &data[nextLowest])
	}

	return data, assignments, comparisons
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
	data, assignments, comparisons := selectionsort(data)
	fmt.Println("sorted:")
	fmt.Println(data)
	fmt.Printf("assignments: %d\n", assignments)
	fmt.Printf("comparisons: %d\n", comparisons)
}
