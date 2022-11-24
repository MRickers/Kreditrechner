package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"kreditrechner/pkg/annuity"
)

type Controller interface {
	Invoke(w http.ResponseWriter, r *http.Request)
}

type Logger interface {
	Info(v ...any)
	Debug(v ...any)
	Warn(v ...any)
	Error(v ...any)
}

type HttpError struct {
	Message string
	Code    int
}

func (e *HttpError) Error() string {
	return e.Message
}

type AnnuityRequest struct {
	Creditsum                  uint    `json:"creditsum"`
	Runtime                    uint    `json:"runtime"`
	Interest_rate              float64 `json:"interest_rate"`
	Initial_repayment_rate     float64 `json:"initial_repayment_rate"`
	Unscheduled_repayment_rate float64 `json:"unscheduled_repayment_rate"`
}

type AnnuityResult struct {
	Year                  uint `json:"year"`
	Residual_dept         uint `json:"residual_dept"`
	Interest              uint `json:"interest"`
	Repayment             uint `json:"repayment"`
	Unscheduled_repayment uint `json:"unscheduled_repayment"`
	Annuity               uint `json:"annuity"`
}

type AnnuityResponse struct {
	Results []AnnuityResult
}

type AnnuityController struct {
	calulator annuity.AnnuityCalculator
	logger    Logger
}

func (ac *AnnuityController) Invoke(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ac.logger.Error("CalculateAnnuityHandler error: want post methode got %s", r.Method)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// get json body
	body, err := readHttpRequestBody(r, ac.logger)

	if err != nil {
		code := err.(*HttpError).Code
		http.Error(w, err.Error(), code)
	}

	// convert from json to annuity.Request
	var annuity_request AnnuityRequest

	err = json.Unmarshal(body, &annuity_request)

	if err != nil {
		ac.logger.Error("Could not unmarshal body")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	request := convertRequest(annuity_request)

	response, err := ac.calulator.Calculate(request)
	if err != nil {
		ac.logger.Error("Calculating request failed: %s", err)
		return
	}
	annuity_response := convertResponse(response)

	json_response, err := json.Marshal(annuity_response)

	if err != nil {
		ac.logger.Error("Marshal json failed: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", string(json_response))
}

func InvokeController(controller Controller) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		controller.Invoke(w, r)
	}
	return http.HandlerFunc(fn)
}

func readHttpRequestBody(r *http.Request, logger Logger) ([]byte, error) {
	if r.ContentLength <= 0 {
		logger.Error("No body available")
		return nil, &HttpError{"Bad Request", http.StatusBadRequest}
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		logger.Error("Could not read request body: %s", err)
		return nil, &HttpError{"Internal Server Error", http.StatusInternalServerError}
	}
	return body, nil
}

func requestAnnuity(request AnnuityRequest) (AnnuityResponse, error) {
	annuity_request := annuity.NewRequest()
	annuity_request.Creditsum = request.Creditsum
	annuity_request.Initial_repayment_rate = request.Initial_repayment_rate
	annuity_request.Interest_rate = request.Interest_rate
	annuity_request.Runtime = request.Runtime
	annuity_request.Unscheduled_repayment_rate = request.Unscheduled_repayment_rate

	response, err := annuity.CalculateAnnuity(annuity_request)

	if err != nil {
		return AnnuityResponse{}, err
	}
	return convertResponse(response), nil
}

func convertRequest(request AnnuityRequest) annuity.Request {
	annuity_request := annuity.NewRequest()
	annuity_request.Creditsum = request.Creditsum
	annuity_request.Initial_repayment_rate = request.Initial_repayment_rate
	annuity_request.Interest_rate = request.Interest_rate
	annuity_request.Runtime = request.Runtime
	annuity_request.Unscheduled_repayment_rate = request.Unscheduled_repayment_rate

	return annuity_request
}
func convertResponse(response annuity.Response) AnnuityResponse {
	var annuity_response AnnuityResponse

	for _, result := range response.Results {
		var annuity_result AnnuityResult

		annuity_result.Annuity = result.Annuity
		annuity_result.Interest = result.Interest
		annuity_result.Repayment = result.Repayment
		annuity_result.Residual_dept = result.Residual_dept
		annuity_result.Unscheduled_repayment = result.Unscheduled_repayment
		annuity_result.Year = result.Year

		annuity_response.Results = append(annuity_response.Results, annuity_result)
	}
	return annuity_response
}
