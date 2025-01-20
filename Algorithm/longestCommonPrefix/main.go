package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	if len(strs) == 0 {
		return ""
	}

	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 && strs[i][:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			break
		}
	}
	return prefix
}
func main() {
	fmt.Println(longestCommonPrefix([]string{"meles", "mesay", "moon"}))
}
