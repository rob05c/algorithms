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

func mergesort(data []int) []int {
	/// DO NOT replace this with the XOR method. It is slow.
	swap := func(a *int, b *int) {
		temp := *a
		*a = *b
		*b = temp
	}

	merge := func(a []int, b []int) []int {
		buffer := make([]int, len(a)+len(b), len(a)+len(b))
		ai, bi := 0, 0
		for ai < len(a) && bi < len(b) {
			if a[ai] < b[bi] {
				buffer[ai+bi] = a[ai]
				ai++
			} else {
				buffer[ai+bi] = b[bi]
				bi++
			}
		}

		if ai != len(a) {
			copy(buffer[ai+bi:], a[ai:])
		} else if bi != len(b) {
			copy(buffer[ai+bi:], b[bi:])
		}

		return buffer
	}

	if len(data) > 2 {
		split := len(data) / 2
		/// These recursive mergesorts can be safely, trivially run in parallel with goroutines, for log n speedup for T∞ = log n processors.
		left := mergesort(data[:split])
		right := mergesort(data[split:])
		return merge(left, right) /// Less trivially, the merge may be parallelised for greater parallelism which scales with manycore architectures.
	}

	if len(data) == 2 {
		if data[0] > data[1] {
			swap(&data[0], &data[1])
		}
		return data
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
	data = mergesort(data)

	fmt.Println(data)
}
