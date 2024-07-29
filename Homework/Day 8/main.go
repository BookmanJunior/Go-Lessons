package main

import "fmt"

func main() {
	nums := []int{5, 3, 4, 10, 99, 65, 32, 55, 43, 44, 1}
	nums4 := []int{1, 0, 1, 2, 0, 1, 3}
	// nums2 := []int{5, -3, 4, -10, -99, 65, -32, -55, 43, 44}
	// nums3 := append(nums, nums2...)
	// fmt.Println(nums3)

	// fmt.Println(FindMax(nums))
	// fmt.Println(FindMin(nums))
	// fmt.Println(LenOfPositiveNums(nums2))
	// fmt.Println(Sum(nums))
	// fmt.Println(MeanOfNums(nums))
	// fmt.Println(FilterItems(nums, 3))
	// fmt.Println(Map(nums, 2))
	// fmt.Println(FindAllIndexes(nums, 3))
	// c := make([]int, len(nums))
	// copy(c, nums)
	// fmt.Println(c)
	// SwapMaxAndMin(nums)
	// fmt.Println(nums)
	// Reverse(nums)
	// fmt.Println(nums)
	// fmt.Println(isPalindrome([]int{9, 5, 2, 2, 5, 9}))
	// fmt.Println(isPalindrome([]int{9, 5, 2, 2, 5, 6}))
	// fmt.Println(moveZeroes(nums4))
	fmt.Println(intersectionOfTwoArrays(nums, nums4))

}

func FindMax(arr []int) int {
	var max int = arr[0]

	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

func FindMin(arr []int) int {
	var min int = arr[0]

	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}
	return min
}

func LenOfPositiveNums(arr []int) int {
	var len int

	for _, num := range arr {
		if num > 0 {
			len++
		}
	}
	return len
}

func Sum(arr []int) int {
	var sum int

	for _, num := range arr {
		sum += num
	}
	return sum
}

func MeanOfNums(arr []int) int {
	return Sum(arr) / len(arr)
}

func FilterItems(arr []int, item int) []int {
	filteredArr := []int{}

	for _, v := range arr {
		if item != v {
			filteredArr = append(filteredArr, v)
		}
	}
	return filteredArr
}

func Map(arr []int, num int) []int {
	newArr := []int{}

	for i := 0; i < len(arr); i++ {
		newArr = append(newArr, arr[i]*num)
	}
	return newArr
}

func FindAllIndexes(arr []int, num int) []int {
	indexes := []int{}

	for i, v := range arr {
		if v == num {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

func SwapMaxAndMin(arr []int) []int {
	var maxIndex int
	var minIndex int

	for i := range arr {
		if arr[maxIndex] < arr[i] {
			maxIndex = i
		}
		if arr[minIndex] > arr[i] {
			minIndex = i
		}
	}

	arr[minIndex], arr[maxIndex] = arr[maxIndex], arr[minIndex]
	return arr
}

func Reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func isPalindrome(arr []int) bool {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

func moveZeroes(arr []int) []int {
	newArr := make([]int, 0)
	zeroes := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zeroes = append(zeroes, arr[i])
		} else {
			newArr = append(newArr, arr[i])
		}
	}
	newArr = append(newArr, zeroes...)
	return newArr
}

func intersectionOfTwoArrays(arr, arr2 []int) []int {
	var intersection []int
	intersected := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr2); j++ {
			_, exists := intersected[arr[i]]
			if exists {
				break
			}
			if arr[i] == arr2[j] {
				intersected[arr[i]] = arr[i]
				intersection = append(intersection, arr[i])

			}
		}
	}
	return intersection
}
