package main

import (
	"fmt"
	"sort"
)

func main ()  {
	var intSlice =[]int{3,4,1,2,7,9,8,6,5,10}
	sort.Ints(intSlice)
	fmt.Println(intSlice)
}