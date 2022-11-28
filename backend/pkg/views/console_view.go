package views

import (
	"fmt"
	"strings"

	"kreditrechner/pkg/annuity"
)

type ConsoleView struct{}

func (v *ConsoleView) Update(request annuity.Request, response annuity.Response) {
	printAnnuityResult(request, response)
}

func printAnnuityResult(request annuity.Request, response annuity.Response) {
	line := "\n\r"
	line += fmt.Sprintf("Interest rate: %.2f %%\n\r", request.Interest_rate)
	line += fmt.Sprintf("Initial repayment rate: %.2f %%\n\r", request.Initial_repayment_rate)
	line += fmt.Sprintf("Unscheduled repayment rate: %.2f %%\n\r", request.Unscheduled_repayment_rate)

	fmt.Println(line)

	year_h := fmt.Sprintf("|%4v|", "Year")
	residual_dept := fmt.Sprintf("%17v|", "Residual dept [€]")
	interest := fmt.Sprintf("%15v|", "Interest [€]")
	repayment := fmt.Sprintf("%15v|", "Repayment [€]")
	unscheduled_repayment := fmt.Sprintf("%22v|", "Unsched. repayment [€]")
	annuity := fmt.Sprintf("%15v|", "Annuity [€]")

	header := year_h + residual_dept + interest + repayment + unscheduled_repayment + annuity

	fmt.Println(header)
	line = "|"
	line += strings.Repeat("=", 93)
	line += "|"
	fmt.Println(line)

	for _, result := range response.Results {
		row := fmt.Sprintf("|%4v|%17v|%15v|%15v|%22v|%15v|", result.Year, result.Residual_dept, result.Interest, result.Repayment, result.Unscheduled_repayment, result.Annuity)
		fmt.Println(row)
	}
	line = "\n\r"
	line += strings.Repeat("=", 40)
	line += "\n\r"
	line += fmt.Sprintf("Interest sum: %d €\n\r", response.Interest_sum)
	line += fmt.Sprintf("Repayment sum: %d €\n\r", response.Repayment_sum)
	line += fmt.Sprintf("Unscheduled repayment sum: %d €\n\r", response.Unscheduled_repayment_sum)
	line += fmt.Sprintf("Total payment: %d €\n\r", response.Annuity_unsched_sum)
	line += strings.Repeat("=", 40)
	fmt.Println(line)
}
