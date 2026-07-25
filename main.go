package main

import (
	"fmt"
	"log"
	"net/http"
)





func main() {
	static_dir := "./static"
	port_num := 8080
	fileServer := http.FileServer(http.Dir(static_dir))


	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)


	fmt.Printf("Starting server at port: %d\n", port_num)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port_num), nil); err != nil {
		log.Fatal(err)
	}
}




// Request: what the client sends to the server.
// ResponseWriter: what the server sends back to the user
func helloHandler(w http.ResponseWriter, r *http.Request){
	// Handling some edge cases 
	if r.URL.Path != "/hello" { // wrong url (path)
		http.Error(w, "404 not found", http.StatusNotFound)
		return 
	}	

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return 
	}



	fmt.Fprintf(w, "hello!")
}



func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return 
	}

	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Addresss = %s\n", address)


}
