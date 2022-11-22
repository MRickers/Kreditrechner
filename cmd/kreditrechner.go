package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"kreditrechner/pkg/annuity"
)

func main() {
	var creditsum, runtime uint
	var interest_rate, initial_repayment_rate, unscheduled_repayment float64

	flag.UintVar(&creditsum, "c", 100000, "Creditsum to calculate for")
	flag.UintVar(&runtime, "r", 10, "Credit runtime [a]")
	flag.Float64Var(&interest_rate, "i", 1.0, "Interest rate [%]")
	flag.Float64Var(&initial_repayment_rate, "repay", 1.0, "Initial repayment rate [%]")
	flag.Float64Var(&unscheduled_repayment, "un-repay", 1.0, "Unscheduled repayment per year [%]")

	flag.Parse()

	if len(os.Args) == 1 {
		flag.PrintDefaults()
		return
	}

	request := annuity.NewRequest()
	request.Creditsum = creditsum
	request.Initial_repayment_rate = initial_repayment_rate
	request.Interest_rate = interest_rate
	request.Runtime = runtime
	request.Unscheduled_repayment_rate = unscheduled_repayment

	response, err := annuity.CalculateAnnuity(request)

	if err != nil {
		fmt.Println(err)
		return
	}

	printAnnuityResult(response)
}

func printAnnuityResult(response annuity.Response) {
	year_h := fmt.Sprintf("|%4v|", "Year")
	residual_dept := fmt.Sprintf("%17v|", "Residual dept [€]")
	interest := fmt.Sprintf("%15v|", "Interest [€]")
	repayment := fmt.Sprintf("%15v|", "Repayment [€]")
	unscheduled_repayment := fmt.Sprintf("%22v|", "Unsched. repayment [€]")
	annuity := fmt.Sprintf("%15v|", "Annuity [€]")

	header := year_h + residual_dept + interest + repayment + unscheduled_repayment + annuity

	fmt.Println(header)
	line := "|"
	line += strings.Repeat("=", 93)
	line += "|"
	fmt.Println(line)

	for _, result := range response.Results {
		row := fmt.Sprintf("|%4v|%17v|%15v|%15v|%22v|%15v|", result.Year, result.Residual_dept, result.Interest, result.Repayment, result.Unscheduled_repayment, result.Annuity)
		fmt.Println(row)
	}
}
