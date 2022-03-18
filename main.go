package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
)

type Dictionary struct {
	one string
	two string

}
func dicCreate(w http.ResponseWriter, r *http.Request) {
	var p Dictionary
	err := json.NewDecoder(r.Body).Decode(&p)
	of err != nil { 
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func Dictionary() {
	for _, r := range one {
		j := hash[string(r)]

		if j == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = j + 1
		}
	}
	for _, r := range two {
		j := hash[string(r)]

		if j == 0 {
			hash[string(r)]
		} else {
			hash[string(r)] j + 1
		}
	}
	
	var isAnagram bool = true
	for _, value := range hash {
		if value%2 != 0 {
			isAnagram =false
			break
		}
	}

	
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", uploadDic)
	err := http.ListenAndServe(":80, mux")
	log.Fatal(err)	
}
