package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cabralfbenja/segurointeligente/internal/db"
	"github.com/cabralfbenja/segurointeligente/internal/handler"
	"github.com/cabralfbenja/segurointeligente/internal/repository"
	"github.com/cabralfbenja/segurointeligente/internal/service"
)

type application struct {
	insuranceHandler *handler.InsuranceHandler
}

func main() {
	sqlDB, err := db.NewDB()
	if err != nil {
		log.Fatalf("error conectando a la base de datos: %v", err)
	}
	defer sqlDB.Close()

	ruleRepo := repository.NewMySQLInsuranceRepository(sqlDB)
	ruleService := service.NewRuleService(ruleRepo)
	ruleHandler := handler.NewInsuranceHandler(ruleService)

	app := &application{
		insuranceHandler: ruleHandler,
	}

	r := app.routes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor escuchando en puerto %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("error iniciando servidor: %v", err)
	}
}
