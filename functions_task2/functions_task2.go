package main

import "fmt"

func longestString(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}

	longest := strs[0]
	for _, s := range strs {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}

func main() {
	str1 := "короткая"
	str2 := "строка подлиннее"
	str3 := "самая длинная строка из всех"
	str4 := "еще одна"

	fmt.Println("Ищем самую длинную строку среди:")
	fmt.Printf("- \"%s\"\n- \"%s\"\n- \"%s\"\n- \"%s\"\n", str1, str2, str3, str4)

	result := longestString(str1, str2, str3, str4)
	fmt.Printf("\nСамая длинная строка: \"%s\"\n", result)

	result2 := longestString("Go", "Python", "JavaScript")
	fmt.Printf("\nСамая длинная из 'Go', 'Python', 'JavaScript': \"%s\"\n", result2)
}
