package main

import (
	"fmt"
	"log"
	"net/http"
)

// ERROR HANDLING
/*****************/

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func ReturnError(w http.ResponseWriter, err error) {
	fmt.Fprint(w, "{\"error\": \"%v\"}", err)
}
