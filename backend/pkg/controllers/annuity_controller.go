package controllers

import (
	"net/http"
)

type Logger interface {
	Info(v ...any)
	Debug(v ...any)
	Warn(v ...any)
	Error(v ...any)
}

func CalculateAnnuityController(logger Logger) http.Handler {
	logger.Info("Hello from logging handler: %s", "Hi")
	fn := func(w http.ResponseWriter, r *http.Request) {
		calculateAnnuityHandler(w, r, logger)
	}
	return http.HandlerFunc(fn)
}

func calculateAnnuityHandler(w http.ResponseWriter, r *http.Request, logger Logger) {
	if r.Method != http.MethodPost {
		logger.Error("CalculateAnnuityHandler error: want post methode got %s", r.Method)
		http.Error(w, "Method Not Allowed", 404)
		return
	}

	// get json
	// logger.Info("Remote address: %s", r.RemoteAddr)
	// logger.Info("HEADER:")
	// for key, value := range r.Header {
	// 	logger.Info("%s: %s", key, value)
	// }
}
