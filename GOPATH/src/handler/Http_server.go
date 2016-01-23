package handler

import (
	"net/http"
	"html"
	"fmt"
	"strings"
)

func HandleHTTP() {
	http.HandleFunc("/test",func(w http.ResponseWriter, r *http.Request){
		result := "Your Data:\n"
		data := r.URL.RawQuery
		splitted := strings.Split(data,"&")
		for i := range splitted {
			result = result + splitted[i] + "\n"
		}
		fmt.Fprint(w,html.EscapeString(result))
		fmt.Println("Request for /test with data " + r.URL.RawQuery)
	})
	http.ListenAndServe(":8080",nil)
	for {}
}