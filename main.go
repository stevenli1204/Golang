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
	//mux := http.NewServeMux()
	//AddUserFunc := http.HandlerFunc(AddUser)

	http.HandleFunc("/dyn-user", HelloUser)
	http.HandleFunc("/check-userid", CheckUserID)
	http.HandleFunc("/add-user", ValidateUserID(AddUser))
	//mux.Handler("/add-user", ValidateUserID(AddUserFunc))
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
		fmt.Fprintf(w, "User %s is deleted successfully!", id)
	} else {
		fmt.Fprintf(w, "There is no user ID %s!", id)
	}
}

func contain(str string, m map[string]bool) bool {
	_, ok := m[str]
	return ok
}

/*
func ValidateUserID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			fmt.Fprintln(w, "Please input a valid ID!")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
*/

func ValidateUserID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			fmt.Fprintln(w, "Please input a valid ID!")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next(w, r)
	}
}
