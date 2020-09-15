package main

import (
	"github.com/joshcarp/cloudfunctionsdemo1"
	"net/http"
)

func main(){
	http.HandleFunc("/", cloudfunctionsdemo1.ServeHTTP)
	http.ListenAndServe(":81", nil)

}