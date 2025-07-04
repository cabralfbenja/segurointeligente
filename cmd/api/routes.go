package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	mux := pat.New()

	mux.Post("/rules", http.HandlerFunc(app.insuranceHandler.Create))
	mux.Get("/rules", http.HandlerFunc(app.insuranceHandler.GetAll))
	mux.Put("/rules/:id", http.HandlerFunc(app.insuranceHandler.Update))

	return mux
}
