package api

import (
	"fmt"
	"net/http"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {

	mux := http.NewServeMux()

	handleRoutes(mux)

	return http.ListenAndServe(s.listenAddr, mux)
}

func handleRoutes(mux *http.ServeMux) {
	//Main Page
	mux.Handle("GET /", http.FileServer(http.Dir("./static")))

	v1 := "/api/v1"
	//Budget Routes
	budgetGroup := fmt.Sprintf("%s/budget", v1)

	mux.HandleFunc("POST "+budgetGroup, createNewBudget)

	//Plaid Routes
}

func createNewBudget(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)

}
