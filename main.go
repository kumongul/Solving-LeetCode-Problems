package main

import (
	"fmt"
	"unicode"
)

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
		if m[ch] < 0 {
			return false
		}
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

// 4
func validPalindrome(s string) bool {
	var clean []rune

	for _, ch := range s {
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			continue
		}
		ch = unicode.ToLower(ch)

		clean = append(clean, ch)
	}

	left, right := 0, len(clean)-1

	for left < right {
		if clean[left] != clean[right] {
			return false
		}
		left++
		right--
	}

	return true
}
