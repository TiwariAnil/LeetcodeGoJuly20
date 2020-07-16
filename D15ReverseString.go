package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {

	s = strings.TrimSpace(s)
	newS := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			newS = newS + " "
			for s[i] == ' ' && i < len(s) {
				i++
			}
			i--
		} else {
			newS = newS + string(s[i])
		}
	}
	var sArr []string
	temp := ""
	for i := 0; i < len(newS); i++ {
		if newS[i] == ' ' {
			if temp != "" {
				sArr = append(sArr, temp)
				temp = ""
			} else {
				for newS[i] == ' ' && i < len(newS) {
					i++
				}
				i--
			}
		} else {
			temp = temp + string(newS[i])
		}
	}
	if temp != "" {
		sArr = append(sArr, temp)
	}
	res := ""
	for i := len(sArr) - 1; i >= 0; i-- {
		res = res + sArr[i]
		if i > 0 {
			res = res + " "
		}
	}
	return res
}

func main() {
	fmt.Println(reverseWords("     the sky      is blue     "))
}
