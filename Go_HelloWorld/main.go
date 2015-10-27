package main

import(
	"fmt"
	"net/http"
	"runtime"
	"os"
)

func indexHandler( w http.ResponseWriter, r *http.Request){
	name, err := os.Hostname()

	if err == nil {
	    fmt.Fprintf(w, "Hello World\nOS: %s\nMachine Name: %s", runtime.GOOS, name)
    }
}

func main(){
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
