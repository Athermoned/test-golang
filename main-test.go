package main

import (
	/* "encoding/json" */
	"fmt"
)

func main() {
	dct := dictionary()
	checkAnagram(dct)

}

func dictionary() map[string][]string {
	dictionary := map[string][]string{
		"frstWrd": {"abba"},
		"scndWrd": {"baab"},
		"thrdWrd": {"biba"},
		"frthWrd": {"boba"},
	}
	/* bs, _ := json.Marshal(dictionary["frstWrd"])
	fmt.Println(string(bs)) */
	return dictionary

}
func checkAnagram(dictionary map[string][]string) {
	hash := make(map[string]int)
	for _, r := range dictionary["frstWrd"] {
		j := hash[string(r)]
		if j == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = j + 1
		}
	}
	for _, r := range dictionary["scndWrd"] {
		j := hash[string(r)]
		if j == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = j + 1
		}
	}
	var checkBool bool = true
	for _, value := range hash {
		if value%2 != 0 {
			checkBool = false
			break
		}
	}
	if checkBool {
		fmt.Println("Эти строки анаграмы")
	} else {
		fmt.Println("Эти строки не анаграмы")
	}
}
