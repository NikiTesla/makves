package items

import (
	"fmt"
	"net/http"
	"strconv"
)

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

	w.Write([]byte(record))
}
