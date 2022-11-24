package annuity

import (
	"fmt"
	"math"
)

type Request struct {
	Creditsum                  uint
	Runtime                    uint
	Interest_rate              float64
	Initial_repayment_rate     float64
	Unscheduled_repayment_rate float64
}

type Result struct {
	Year                  uint
	Residual_dept         uint
	Interest              uint
	Repayment             uint
	Unscheduled_repayment uint
	Annuity               uint
}

type Response struct {
	Results []Result
}

type Observer interface {
	Update(response Response)
}

type Observerable interface {
	AddObserver(observer Observer)
	NotifyObservers(response Response)
}

type AnnuityCalculator struct {
	observers []Observer
}

func (ac *AnnuityCalculator) Calculate(request Request) (Response, error) {
	response, err := CalculateAnnuity(request)

	if err != nil {
		return Response{}, err
	}
	ac.NotifyObservers(response)
	return response, nil
}

func (ac *AnnuityCalculator) AddObserver(observer Observer) {
	ac.observers = append(ac.observers, observer)
}

func (ac *AnnuityCalculator) NotifyObservers(response Response) {
	for _, observer := range ac.observers {
		observer.Update(response)
	}
}

func NewRequest() Request {
	return Request{100000, 10, 1, 1, 1}
}

func CalculateAnnuity(request Request) (Response, error) {
	err := sanityChecks(request)

	if err != nil {
		return Response{}, err
	}

	old_result := initResult(request)
	response := Response{
		Results: []Result{old_result},
	}
	interest_rate := request.Interest_rate / 100
	unscheduled_repayment_rate := request.Unscheduled_repayment_rate / 100

	for year := 0; year < int(request.Runtime); year++ {
		next_result := calulateResultForNextYear(old_result, unscheduled_repayment_rate, interest_rate)
		response.Results = append(response.Results, next_result)
		old_result = next_result
	}

	return response, nil
}

func sanityChecks(request Request) error {
	const RUNTIME_UPPER_BOUND uint = 100

	if request.Runtime > RUNTIME_UPPER_BOUND {
		return fmt.Errorf("runtime exceeds upper bound %d", request.Runtime)
	}

	return nil
}

func initResult(request Request) Result {
	result := Result{}
	result.Annuity = calculateAnnuity(request)
	result.Residual_dept = request.Creditsum
	result.Interest = uint(math.Round(float64(request.Creditsum) * request.Interest_rate / 100))
	result.Repayment = uint(result.Annuity - result.Interest)
	result.Unscheduled_repayment = uint(math.Round(float64(request.Creditsum) * request.Unscheduled_repayment_rate / 100))
	result.Year = 1

	return result
}

func calculateAnnuity(request Request) uint {
	interest_factor := request.Interest_rate / 100
	var annuity = math.Round(float64(request.Creditsum) * (request.Initial_repayment_rate / 100))
	annuity += math.Round(float64(request.Creditsum) * interest_factor)
	return uint(annuity)
}

func calulateResultForNextYear(
	old_result Result,
	unscheduled_repayment_rate float64,
	interest_rate float64) Result {
	var result Result

	result.Annuity = old_result.Annuity
	result.Residual_dept = old_result.Residual_dept -
		old_result.Repayment -
		old_result.Unscheduled_repayment
	result.Interest = uint(math.Round(float64(result.Residual_dept) * interest_rate))
	result.Repayment = result.Annuity - result.Interest
	result.Unscheduled_repayment = uint(math.Round(float64(result.Residual_dept) * unscheduled_repayment_rate))

	result.Year = old_result.Year + 1
	return result
}
