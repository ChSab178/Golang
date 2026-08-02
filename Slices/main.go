package main

import (
	"fmt"
	"slices"
)

func main() {
	// Initialized a slice:

	var arr []int
	fmt.Println(arr) // Stores 'nil' when it is initialized not 'null'.

	fmt.Println(len(arr))
	fmt.Println(append(arr, 1))

	// 2nd way to create the slices:

	var nums = make([]int, 2, 5) // It creates the num slice of int dtype and stores 0's in it for 2 times as we give it the capacity of 5.

	// Now as we create the slice if we want to add the value based on index in the slice as in array we write:
	nums[0] = 1 // Make sure you have enough sized given in the make function to access the following index to add values.

	fmt.Println(nums)
	fmt.Println(nums)
	fmt.Println(cap(nums))       // Cap stands for capacity. For the slices they doesn't matter cuz the slices are dynamic so we can store as much values as we want.
	fmt.Println(append(nums, 1)) // It takes slice and value and add the given value at the end of given array.

	// How to copy a slice to another one:

	var numsCopy = make([]int, 3, 5)
	// Most of the people think that way:
	// numsCopy=nums But it makes the actual refernce not the copy if you change anything in numsCopy it will also changed in nums slice.
	copy(numsCopy, nums) // Make sure copy arr has size given in the make function more than 0. If the src slice has more size than the destination slice size so the destination no of size values from src will be stored in destination slice.
	fmt.Println(numsCopy)

	// Slice operator:
	var num1 = []int{1, 2, 3}
	fmt.Println(num1[0:2]) // It starts from the 0 index to val of slice - 1. It doesn't include the last value. As in this case in 0 index : 1, 1 index:2 but it will not going to give value of 2 index.
	fmt.Println(num1[:])   // If we don't give it starts from 0 index value and if we don't give it last index it will include the last value too.
	fmt.Println(num1[:2])
	fmt.Println(num1[1:])

	// How to compare the 2 slices:
	// We use a package knows as slices.
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 2, 3}
	fmt.Println(slices.Equal(slice1, slice2)) // .Equal is the method in slices package that gives true:1 of false:0 based on the slices.

	// How to make 2d slices:
	var slice3=[][]int{{1,2,3},{1,2,3}}
	fmt.Println(slice3)

}
