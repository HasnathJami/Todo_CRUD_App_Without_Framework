package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}){
	// if body, err := ioutil.ReadAll(r.Body); err !=nil {
	// 	if err := json.Unmarshal([]byte (body), x); err !=nil {
	// 		return
	// 	}
	// }
    body, err := ioutil.ReadAll(r.Body)
	checkError(err)
	err = json.Unmarshal([]byte (body), x)
	checkError(err)
}

func checkError( err error){
    if err != nil{
        log.Fatal(err)
    }
}