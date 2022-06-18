package main

import (
	"fmt"
	"sort"
)

var myInt int
var myInt16 int16
var myInt32 int32
var myInt64 int64
var myUint uint // unsigned int only positive numbers
var myFloat float32
var myFloat64 float64
var myString string // Strings are immutable
var myBool bool     // bool values

type Car struct {
	price float64
	Model string
}

func main() {

	var animals []string
	animals = append(animals,"dog")

	// Are they sorted?
	fmt.Println("Slice is sorted?",sort.StringsAreSorted(animals))
	sort.Strings(animals) // Sorts the strings

	for index,x := range animals {
		fmt.Println(x,index)
	}

	var myStrings [3]string // Its an array
	myStrings[0] = "a"
	myStrings[1] = "b"
	myStrings[2] = "c"

	var myCar Car

	myCar.price = 100
	myCar.Model = "Tesla"

	// myCar2 := Car{
	// 	price:10,
	// 	Model:"BMW",
	// }
	
	x := 10
	myPointer := &x
	*myPointer = 15
	fmt.Println(x," ",myPointer)
}

func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a 
}

func maps() {
	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3

	delete(intMap, "two")

	el, ok := intMap["one"]

	intMap["three"] = 3 // Overwrite the default value

	// Check if the map contains the specified value
	if ok {
		fmt.Println(el,"is in map")
	} else {
		fmt.Println(el,"is not in map")
	}

	for key, val := range intMap {
		fmt.Println(key,val)
	}

}

func sumMany(nums ...int) int {
	total := 0

	for _,value := range nums {
		total += value

	}
	return total
}