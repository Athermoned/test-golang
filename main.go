package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
)

func checkAnagram(w1, w2 string) bool {
	wa := [2]string{w1, w2}
	for i := range wa {
		sa := strings.Split(strings.ToLower(wa[i]), "")
		sort.Strings(sa)
		wa[i] = strings.Join(sa, "")
	}
	return wa[0] == wa[1]
}

func getWords(key string) []string {
	var ret []string
	for _, v := range words {
		if checkAnagram(key, v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// global
var (
	words []string
	cntr  int
)

func main() {

	cntr = 0

	http.HandleFunc("/load", func(w http.ResponseWriter, r *http.Request) {
		cntr++
		log.Println(r.Method, r.URL)
		if r.Method != http.MethodPost {
			log.Println("Unallowed method!")
			return
		}
		buf := make([]byte, r.ContentLength) // No length limit
		_, err := r.Body.Read(buf)
		if err != nil && err != io.EOF { // .Read returns err = io.EOF even if succeed
			log.Println("Body.Read err:", err)
			return
		}
		err = json.Unmarshal(buf, &words)
		if err != nil {
			log.Println("json.Unmarshall err:", err)
			return
		}
		log.Println("Words uploaded:", words)
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		cntr++
		log.Println(r.Method, r.URL)
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			return
		}
		if len(r.Form["word"]) == 0 {
			return
		}
		ww := getWords(r.Form["word"][0])
		data, err := json.Marshal(ww)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = w.Write(data)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Anagrams:", ww)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		cntr++
		log.Println(r.Method, r.URL)
		data := map[string]interface{}{
			"name":     "dictionary",
			"status":   "ok",
			"counter":  cntr,
			"dictSize": len(words),
		}
		bdata, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = w.Write(bdata)
		if err != nil {
			log.Println(err)
			return
		}
	})

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
