package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	unixTimeStamp := "1432572732"

	unixIntValue, err := strconv.ParseInt(unixTimeStamp, 10, 64)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	timeStamp := time.Unix(unixIntValue, 0)

	result := fmt.Sprint(timeStamp)
	fmt.Printf("Result value is %v and type is %T\n", result, result)
	fmt.Printf("Now in UTC is %v\n", time.Now().UTC())
	fmt.Printf("Now in Local is %v\n", time.Now().Local())
}
