package items

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type RestServer struct {
	storage *LocalStorage
}

// Run starts rest server with several endpoints
func (rS *RestServer) Run() error {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/get-items", rS.getItems).Methods("GET")

	log.Println("Server started")
	return http.ListenAndServe(":8080", rtr)
}

// NewRestServer creates rest server with local storage from file
func NewRestServer() *RestServer {
	storage, err := NewLocalStorage(os.Getenv("ITEMS_STORAGE_FILE"))
	if err != nil {
		log.Fatalf("Cannot create storage for server, err: %s", err)
	}

	return &RestServer{
		storage: storage,
	}
}
