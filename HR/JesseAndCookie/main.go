package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'cookies' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY A
 */

func sweetness(a, b int32) int32 {
	return a + (2 * b)
}

func add(A []int32, v int32) []int32 {
	return append(A, v)
}

func print(A []int32) {
	fmt.Printf("A: %v\n", A)
}

func removeValue(A []int32, v int32) []int32 {
	var result []int32
	for i, a := range A {
		if v == a {
			return append(A[:i], A[i+1:]...)
		}
	}
	return result
}

func checkDone(A []int32, k int32) bool {
	if min(A) >= k {
		return true
	} else {
		return false
	}
}

func min(array []int32) int32 {
	if len(array) == 0 {
		panic("empty array")
	}

	var min int32 = array[0]
	for _, value := range array {
		if min > value {
			min = value
		}
	}
	return min
}

func cookies(k int32, A []int32) int32 {
	result := 0
	print(A)
	if min(A) >= k {
		return 0
	}
	for len(A) > 1 {
		min1 := min(A)
		A = removeValue(A, min1)
		print(A)
		min2 := min(A)
		A = removeValue(A, min2)
		print(A)
		sw := sweetness(min1, min2)
		A = add(A, sw)
		result++

		if min(A) >= k {
			return int32(result)
		}
	}
	if min(A) >= k {
		return int32(result)
	} else {
		return -1
	}
}

func main() {
	// open file
	// fmt.Println(os.Getwd())
	file, err := os.Open("input03.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("output01.txt"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	ATemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	fmt.Println(len(ATemp))
	var A []int32

	for i := 0; i < int(n); i++ {
		AItemTemp, err := strconv.ParseInt(ATemp[i], 10, 64)
		checkError(err)
		AItem := int32(AItemTemp)
		A = append(A, AItem)
	}

	result := cookies(k, A)

	fmt.Printf("%d\n", result)

	// writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
