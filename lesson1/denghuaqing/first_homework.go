package  main

import (
	"net/http"
	"fmt"
)

func funderl(r http.ResponseWriter,w *http.Request)  {
	fmt.Fprintf(r,"hello %s",w.URL.Query().Get("user"))
}
func main()  {
	http.HandleFunc("/",funderl)
	http.ListenAndServe(":8800",nil)
}
