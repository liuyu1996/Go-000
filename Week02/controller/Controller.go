package controller

import (
	"Go-000/Week02/service"
	"fmt"
	"net/http"
)

func Controller(w http.ResponseWriter, r * http.Request)  {
	result, err := service.Service()
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.Write([]byte(result))
	}
	w.Write([]byte("ok"))
}