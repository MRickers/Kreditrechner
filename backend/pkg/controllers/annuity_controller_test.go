package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/stretchr/testify/assert"
)

type DummyLogger struct {
	Message string
}

func (l *DummyLogger) Info(v ...any) {
	strConvert := fmt.Sprintf("%s", v[0])
	l.Message += strConvert
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
		SetJSON(gofight.D{
			"creditsum":                  350000,
			"initial_repayment_rate":     2,
			"interest_rate":              3,
			"unscheduled_repayment_rate": 2,
			"runtime":                    10,
		}).
		SetDebug(true).
		Run(http.DefaultServeMux, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			var response AnnuityResponse
			err := json.Unmarshal(r.Body.Bytes(), &response)

			assert.Equal(t, nil, err)
			assert.Equal(t, response.Results[4].Residual_dept, uint(293154))
			assert.Equal(t, response.Results[4].Interest, uint(8795))
			assert.Equal(t, response.Results[4].Repayment, uint(8705))
			assert.Equal(t, response.Results[4].Unscheduled_repayment, uint(5863))
			assert.Equal(t, response.Results[4].Annuity, uint(17500))
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
