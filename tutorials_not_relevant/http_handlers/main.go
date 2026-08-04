package main

import (
	"log"
	"net/http"
)


type server struct {
	addr string 

}


// signature: parameter type, return type and receiver form
// method name: or function name


func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	// w.Write([]byte("Hello from the server"))
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("index page")) // test with curl http://localhost:8080/
			return 
		case "/users":
			w.Write([]byte("users page")) // test with curl http://localhost:8080/users
			return 
		}
	default:
		w.Write([]byte("404 page")) // test with curl -X POST http://localhost:8080/
		return 
	}
}


func main(){
	s := &server{
		addr: ":8080",
	}
	if err := http.ListenAndServe(":8080", s); err != nil {
		log.Fatal(err)
	}


}