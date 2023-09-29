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
	fmt.Fprintf(w, "Hello, %s\n", id)
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
		fmt.Fprintf(w, "User %s is existed!\n", id)
	} else {
		validIDs[id] = true
		fmt.Fprintf(w, "User %s is added successfully!\n", id)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if contain(id, validIDs) {
		delete(validIDs, id)
		fmt.Fprintf(w, "User %s is deleted successfully!\n", id)
	} else {
		fmt.Fprintf(w, "There is no user ID %s!\n", id)
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
