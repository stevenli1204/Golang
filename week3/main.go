package main

import (
	"fmt"
	"net/http"
)

var validIDs = map[string]bool{
	"11111": true,
	"22222": true,
	"33333": true,
}

func main() {
	http.HandleFunc("/dyn-user", HelloUser)
	http.HandleFunc("/check-userid", CheckUserID)
	http.HandleFunc("/add-user", AddUser)
	http.HandleFunc("/delete-user", DeleteUser)
	http.ListenAndServe(":8080", nil)
}

func HelloUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Fprintln(w, "Hello, %s", id)
}

func CheckUserID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if contain(id, validIDs) {
		fmt.Fprintln(w, "Hello, dear user!")
		return
	}
	w.WriteHeader(http.StatusForbidden)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if contain(id, validIDs) {
		fmt.Fprintln(w, "User %s is existed!", id)
	} else {
		validIDs[id] = true
		fmt.Fprintln(w, "User %s is added successfully!", id)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if contain(id, validIDs) {
		delete(validIDs, id)
		fmt.Fprintln(w, "User %s is deleted successfully!", id)
	} else {
		fmt.Fprintln(w, "There is no user ID %s!", id)
	}
}

// Don't need to use range here
func contain(str string, m map[string]bool) bool {
	if _, ok := m[str]; ok {
		return true
	} else {
		return false
	}
}
