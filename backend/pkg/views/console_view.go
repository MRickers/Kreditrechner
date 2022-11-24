package views

import (
	"fmt"
	"strings"

	"kreditrechner/pkg/annuity"
)

type ConsoleView struct{}

func (v *ConsoleView) Update(response annuity.Response) {
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
