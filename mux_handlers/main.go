package main

import (
	"fmt"
	"net/http"
)





func main(){
	// implements the httpHandler interface
	api := &api{addr: ":8080"}

	// initialize the ServeMux
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr: api.addr,
		Handler: mux,
	}


	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port 8080")
}