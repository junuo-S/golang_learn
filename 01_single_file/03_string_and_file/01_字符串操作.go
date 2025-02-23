package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contains
	// strings.Contains(s, substr string) bool
	fmt.Println(strings.Contains("hello world", "world")) // true
	fmt.Println(strings.Contains("hello world", "abc"))   // false

	// Join
	// strings.Join(a []string, sep string) string
	slice1 := []string{"abc", "def", "ghi"}
	fmt.Println(strings.Join(slice1, ",")) // abc,def,ghi

	// Index
	// strings.Index(s, sep string) int
	fmt.Println(strings.Index("hello world", "world")) // 6
	fmt.Println(strings.Index("hello world", "abc"))   // -1

	// Repeat
	// strings.Repeat(s string, count int) string
	fmt.Println(strings.Repeat("hello", 3)) // hellohellohello

	// Split
	// strings.Split(s, sep string) []string
	fmt.Println(strings.Split("hello,world", ",")) // [hello world]

	// Trim 去掉字符串首尾的指定字符
	// strings.Trim(s string, cutset string) string
	fmt.Println(strings.Trim(" hello world ", " ")) // hello world

	// Fields 去掉字符串首尾的空格，然后以空格分割字符串
	// strings.Fields(s string) []string
	fmt.Println(strings.Fields(" hello world ")) // [hello world]
}
