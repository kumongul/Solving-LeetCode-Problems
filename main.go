package main

import (
	"fmt"
	"unicode"
)

func main() {
	nums := []int{1, 0, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
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

// 5
func moveZeroes(nums []int) {
	write := 0

	for read := 0; read < len(nums); read++ {
		if nums[read] != 0 {
			nums[write] = nums[read]
			write++
		}
	}
	for i := write; i < len(nums); i++ {
		nums[i] = 0
	}
}

// 6
func validParentheses(s string) bool {
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack []rune

	for _, ch := range s {

		// 1. если ОТКРЫВАЮЩАЯ → push
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
			continue
		}

		if ch == ')' || ch == ']' || ch == '}' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if pairs[top] == ch {
				stack = stack[:len(stack)-1]
				continue
			}
			return false
		}
	}

	return len(stack) == 0
}
