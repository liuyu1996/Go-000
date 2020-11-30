package main

import (
	"Go-000/Week02/controller"
	"net/http"
)

func main()  {
	http.HandleFunc("/",controller.Controller)
	http.ListenAndServe(":8080", nil)
}
