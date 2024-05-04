package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/login", handleLoginRequest)
// 	http.HandleFunc("/logout", handleLogoutRequest)
// 	fs := http.FileServer(http.Dir("static"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))
// 	fmt.Println("listening at port 8080 ....")
// 	http.ListenAndServe(":8080", nil)
// }

// func handleLoginRequest(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "hello from the login endpoint. You visited %s\n", r.URL.Path)
// }

// func handleLogoutRequest(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "hello from the logout endpoint. You visited %s\n", r.URL.Path)
// }
