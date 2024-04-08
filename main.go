package main

import (
	"net/http"
)
func main(){
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	} )
	http.HandleFunc("/world" , func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("world"))
	})
	http.ListenAndServe(":8000" , nil)

}