package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ariefzainuri96/go-api-blogging/cmd/api/request/xendit"
)

func (app *application) fvaCreatedUpdated(w http.ResponseWriter, r *http.Request) {
	var data xendit.FVACreatedUpdatedReq
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	jsonData, _ := data.Marshal()

	log.Println(string(jsonData))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("fva-created-updated"))
}

func (app *application) fvaPaid(w http.ResponseWriter, r *http.Request) {
	var data xendit.FVAPaidReq
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	jsonData, _ := data.Marshal()

	log.Println(string(jsonData))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("fva-paid"))
}

func (app *application) invoice(w http.ResponseWriter, r *http.Request) {
	var data xendit.InvoiceReq
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = app.store.Order.UpdateStatusOrder(r.Context(), data.ID, data.Status)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("invoice"))
}

func (app *application) XenditCallbackRouter() *http.ServeMux {
	xenditRouter := http.NewServeMux()

	xenditRouter.HandleFunc("POST /fva-created-updated", app.fvaCreatedUpdated)
	xenditRouter.HandleFunc("POST /fva-paid", app.fvaPaid)
	xenditRouter.HandleFunc("POST /invoice", app.invoice)

	// Catch-all route for undefined paths
	xenditRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})

	return xenditRouter
}
