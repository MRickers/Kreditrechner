package main

import (
	"fmt"
	"net/http"
	"os"

	"kreditrechner/pkg/annuity"
	"kreditrechner/pkg/controllers"
	"kreditrechner/pkg/views"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.Out = os.Stdout

	annuity_calculator := annuity.AnnuityCalculator{}
	console_view := views.ConsoleView{}

	annuity_calculator.AddObserver(&console_view)

	annuity_controller := controllers.NewAnnuityController(logger, annuity_calculator)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculateAnnuity", annuity_controller.Invoke)
	mux.Handle("/", http.FileServer(http.Dir("./service/static")))

	addr := ":5000"
	logger.Info(fmt.Sprintf("Listening on %s", addr))
	logger.Error(http.ListenAndServe(addr, mux))
}
