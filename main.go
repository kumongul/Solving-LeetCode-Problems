package main

import "fmt"

func main() {
	nums := []int{10, 20, 30}
	fmt.Println(containsDuplicate(nums))
}

// 1
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		need := target - num

		if j, ok := hashMap[need]; ok {
			return []int{i, j}
		}
		hashMap[num] = i
	}
	return nil
}

// 2
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}
	return false
}

// 3
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)

	for _, ch := range s {
		m[ch]++
	}
	for _, ch := range t {
		m[ch]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
