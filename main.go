package main

import (
	"fmt"
	"net/http"

	"github.com/rahul-golang/exercise4/handlers"
	"github.com/rahul-golang/exercise4/services"
)

func main() {

	service := services.NewAppServiceImpl()
	handler := handlers.NewAppHandlers(service)

	http.HandleFunc("/", handler.ServeApplication)
	http.HandleFunc("/date-converter", handler.GetFormatedDateTime)
	fmt.Println(http.ListenAndServe(":8000", nil))
}
