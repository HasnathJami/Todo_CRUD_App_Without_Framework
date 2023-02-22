package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HasnathJami/todo-crud/routers"
)


func main() { 
    
	fmt.Println("Server Has Started Successfully")
	routers.Router()
	log.Fatal(http.ListenAndServe(":4041",nil))
	//fmt.Println(tasks)
}