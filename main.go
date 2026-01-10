package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request){    // w :- response r :- request
    if r.URL.Path != "/hello" {
		http.Error(w , "404 not found" , http.StatusNotFound)
		// syntax: func Error(w ResponseWriter, error string, code int) , here statusnot found is 404
        return
	}

    if r.Method != "GET"{  // we only want user to GET not to POST,UPDATE or DELETE
		http.Error(w , "Method is not supported", http.StatusNotFound)
		return // return, require because to stop executio otherwise hello and error is both given 
	}
	fmt.Fprintf(w , "HELLO!")
	// syntax func fmt.Fprintf(w io.Writer, format string, a ...any) (n int, err error)
	// Fprintf is required becoz we need to write in http response not to terminal

}

func formHandler(w http.ResponseWriter ,r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w,"ParseForm() error: %v", err)
        return 
	}
	// ParseForm() : it reads incoming http request data and store in go DS 
	//%v it is a format verb that returns the error , if parsing succeeds then err=nil.

	fmt.Fprintf(w,"POST request successful\n")
	name := r.FormValue("name") // r.FormValue takes form value from request (client)
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name=%s\n" , name)
	fmt.Fprintf(w, "Address=%s\n" , address)

}



func main() {
	fileServer := http.FileServer(http.Dir("./static")) 
	// creates a handler that serves static files
	// handler is a code that runs when web request come to ur GO server.  hadles request and response
    // http.FileServer is a built-in Go handler that serves files from your computer to a web browser.
	
	
	//routing
	http.Handle("/",fileServer) //  if not request matched serve files from ./static i.e index.html
	http.HandleFunc("/form", formHandler) 
	http.HandleFunc("/hello", helloHandler)
    //When a request comes for this path (URL), if matches , run this handler.
	// syntax: http.Handle(path string, handler http.Handler) :- register handler to given URL path
 
	// syntax: http.HandleFunc(path string, handlerFunc func(ResponseWriter, *Request)) :- register func asa handler

	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080",nil); err != nil {
    log.Fatal(err)
	}
    
	// Listen: Open a network TCP port (8080) and wait for incoming requests , Serve = Handle incoming requests and send responses.
    // nil means use default HTTP router (http.DefaultServeMux)” above handlers are register to this router;
// 1: OS opens port 8080
// Server waits for requests
// Browser hits localhost:8080
// Go routes request to handler
// Handler sends response
// Browser receives response

// why error is present? :- Port already in use,Permission denied,Network failure etc.



}

// run: go run main.go
// http://localhost:8080/ show index.html 
