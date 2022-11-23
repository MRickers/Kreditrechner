package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/stretchr/testify/assert"
)

type DummyLogger struct {
	message string
}

func (l *DummyLogger) Info(v ...any) {
	strConvert := fmt.Sprintf("%s", v)
	l.message += strConvert
}

func (l *DummyLogger) Debug(v ...any) {}
func (l *DummyLogger) Warn(v ...any)  {}
func (l *DummyLogger) Error(v ...any) {}

func HTTPCalculateAnnuityHandler(logger *DummyLogger) {
	annuityHandler := CalculateAnnuityController(logger)
	http.HandleFunc("/test", annuityHandler.ServeHTTP)
}

func TestCalculateAnnuityController(t *testing.T) {
	logger := DummyLogger{}
	HTTPCalculateAnnuityHandler(&logger)

	r := gofight.New()

	r.POST("/test").
		// trun on the debug mode.
		SetDebug(true).
		Run(http.DefaultServeMux, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
