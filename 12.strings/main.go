package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "hello world"
	str2 := "world"
	fmt.Println(strings.Contains(str1, str2))
	fmt.Println(strings.Compare(str1, str2))
	fmt.Println(strings.HasSuffix("abcd", "cd"))
	fmt.Println(strings.Split(str1, " "))
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "_"))
	fmt.Println(strings.Trim("             Hello", " "))

}
