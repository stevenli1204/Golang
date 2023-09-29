package main

import (
	"fmt"
	"net/http"
)

var ids = []string{"11111", "22222", "33333"}

func main() {
	http.HandleFunc("/dyn-user", HelloUser)
	http.HandleFunc("/check-userid", CheckUserID)
	http.ListenAndServe(":8080", nil)
}

func HelloUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Fprintln(w, "Hello, %s", id)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func CheckUserID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	switch contains(ids, id) {
	case true:
		fmt.Fprintln(w, "Hello, dear user!")
	case false:
		fmt.Fprintln(w, "Error 403")
	}
}
