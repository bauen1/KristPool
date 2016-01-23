package handler

import (
	"net/http"
	//"html"
	"fmt"
	"strings"
	"os"
	"io/ioutil"
)

func HandleHTTP() {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		if _, err := os.Stat("html/" + r.URL.Path); err == nil {

		} else {
			if _, err := os.Stat("html/errors/404.html"); err == nil {
				dat,_ := ioutil.ReadFile("html/errors/404.html")
				fmt.Fprint(w,strings.Replace(string(dat),"$URL",r.URL.Path,-1))
			} else {
				fmt.Fprint(w,"404.html not found! Seriously, we lost our 404 page...")
			}
		}
	})
	http.ListenAndServe(":8080",nil)
}