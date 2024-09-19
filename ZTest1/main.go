// {1,1,2,3,3,4,5}
// {1,3,5,6,6,7}
// {2,4,6,7}

package main

import "fmt"

// arr1 = N
// arr2 = M

// O(M+N)

func main() {
	arr1 := []int{1, 1, 2, 3, 3, 4, 5}
	arr2 := []int{1, 3, 5, 6, 6, 7}

	duplicate := make(map[int]int)
	for i := 0; i < len(arr1); i++ {
		duplicate[arr1[i]] = 1
	}
	fmt.Println(duplicate)
	for i := 0; i < len(arr2); i++ {
		if _, found := duplicate[arr2[i]]; found {
			duplicate[arr2[i]]++
		} else {
			duplicate[arr2[i]] = len(arr2) * -1
		}
	}
	fmt.Println(duplicate)
	for i, count := range duplicate {
		if count < 2 {
			fmt.Println(i)
		}
	}
}
