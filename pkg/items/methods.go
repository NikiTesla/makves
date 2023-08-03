package items

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// getItems returns response with json formatted record from local storage
func (rS *RestServer) getItems(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot parse id from url query, err: %s\n", err), http.StatusBadRequest)
		return
	}

	record, err := rS.storage.getItem(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot find such record in storage, err: %s", err), http.StatusNotFound)
		return
	}

	output, err := json.Marshal(record)
	if err != nil {
		log.Printf("cannot marshal record in json format, err: %s\n", err)
	}

	w.Write(output)
}
