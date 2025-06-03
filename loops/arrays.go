package main

import "fmt"

func function3() {
	var arr [2]string
	arr[0] = "harry"
	arr[1] = "poter"
	fmt.Println(arr)

	//short way to initialized
	newarr := []string{"harray", "poter"}
	fmt.Println(newarr)
	var permises []string = newarr[0:2]
	fmt.Println(permises)

	// sice1 := arr[0:1]
	// sice2 := arr[1:2]
	var slice []int
	fmt.Printf("length is %d cap is %d", len(slice), len(slice))
	var name string
	name = "rashi"
	fmt.Println(name)

	// fmt.Println(sice1, sice2)

	// sice1[0] = "rashi"
	// fmt.Println(arr)
	// fmt.Printf("lemgth of sice is %d and capacity is %d", len(sice1), cap(sice1))

}
