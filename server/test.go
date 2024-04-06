package main

import (
	"fmt"
	"time"
)

func main() {
	t, _ := time.Parse("2006-01-02 15:04:05", "2022-12-01 14:00:00")
	fmt.Println("time is ", t)

	var x byte = '1'
	fmt.Println(x)
	fmt.Println(`"` + "1111")

}
