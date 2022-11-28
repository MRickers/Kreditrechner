package annuity

import "testing"

func TestCalculatingAnnuity1(t *testing.T) {
	request := NewRequest()
	request.Creditsum = 350000
	request.Initial_repayment_rate = 2
	request.Unscheduled_repayment_rate = 2
	request.Runtime = 15
	request.Interest_rate = 3

	response, err := CalculateAnnuity(request)

	if err != nil {
		t.Fatalf("Calculating Annuity failed: %s", err)
	}

	if response.Results[0].Annuity != 17500 {
		t.Fatalf("Calculating Annuity failed: %d != %d", response.Results[0].Annuity, 17500)
	}
	if response.Results[0].Interest != 10500 {
		t.Fatalf("Calculating Interest failed: %d != %d", response.Results[0].Interest, 10500)
	}
	if response.Results[0].Repayment != 7000 {
		t.Fatalf("Calculating Repayment failed: %d != %d", response.Results[0].Repayment, 7000)
	}
	if response.Results[0].Unscheduled_repayment != 7000 {
		t.Fatalf("Calculating Unscheduled_Repayment failed: %d != %d", response.Results[0].Unscheduled_repayment, 7000)
	}
	if response.Results[0].Residual_dept != 350000 {
		t.Fatalf("Calculating Residual_dept failed: %d != %d", response.Results[0].Residual_dept, 350000)
	}
}

func TestCalculatingAnnuity2(t *testing.T) {
	request := NewRequest()
	request.Creditsum = 350000
	request.Initial_repayment_rate = 2
	request.Unscheduled_repayment_rate = 2
	request.Runtime = 15
	request.Interest_rate = 3

	response, err := CalculateAnnuity(request)

	if err != nil {
		t.Fatalf("Calculating Annuity failed: %s", err)
	}

	if response.Results[10].Annuity != 17500 {
		t.Fatalf("Calculating Annuity failed: %d != %d", response.Results[10].Annuity, 17500)
	}
	if response.Results[10].Interest != 6106 {
		t.Fatalf("Calculating Interest failed: %d != %d", response.Results[10].Interest, 6106)
	}
	if response.Results[10].Repayment != 11394 {
		t.Fatalf("Calculating Repayment failed: %d != %d", response.Results[10].Repayment, 11394)
	}
	if response.Results[10].Unscheduled_repayment != 4071 {
		t.Fatalf("Calculating Unscheduled_Repayment failed: %d != %d", response.Results[10].Unscheduled_repayment, 4071)
	}
	if response.Results[10].Residual_dept != 203529 {
		t.Fatalf("Calculating Residual_dept failed: %d != %d", response.Results[10].Residual_dept, 203529)
	}
}

func TestCalculatingAnnuity3(t *testing.T) {
	request := NewRequest()
	request.Creditsum = 350000
	request.Initial_repayment_rate = 1.8
	request.Unscheduled_repayment_rate = 1.6
	request.Runtime = 15
	request.Interest_rate = 2.1

	response, err := CalculateAnnuity(request)

	if err != nil {
		t.Fatalf("Calculating Annuity failed: %s", err)
	}

	if response.Results[13].Annuity != 13650 {
		t.Fatalf("Calculating Annuity failed: %d != %d", response.Results[13].Annuity, 13650)
	}
	if response.Results[13].Interest != 4002 {
		t.Fatalf("Calculating Interest failed: %d != %d", response.Results[13].Interest, 4002)
	}
	if response.Results[13].Repayment != 9648 {
		t.Fatalf("Calculating Repayment failed: %d != %d", response.Results[13].Repayment, 9648)
	}
	if response.Results[13].Unscheduled_repayment != 3049 {
		t.Fatalf("Calculating Unscheduled_Repayment failed: %d != %d", response.Results[13].Unscheduled_repayment, 3049)
	}
	if response.Results[13].Residual_dept != 190572 {
		t.Fatalf("Calculating Residual_dept failed: %d != %d", response.Results[13].Residual_dept, 190572)
	}
}

func TestCalculatingAnnuityInvalidRuntime(t *testing.T) {
	request := NewRequest()
	request.Runtime = 101

	_, err := CalculateAnnuity(request)

	if err == nil {
		t.Fatalf("Calculating Annuity failed: want error got none")
	}
}

func TestCalculatingAnnuityResidualOverflow(t *testing.T) {
	request := NewRequest()
	request.Runtime = 30
	request.Creditsum = 100000
	request.Interest_rate = 1
	request.Initial_repayment_rate = 3
	request.Unscheduled_repayment_rate = 2

	result, err := CalculateAnnuity(request)

	if err != nil {
		t.Fatalf("Calculating Annuity failed: want error got none")
	}
	if len(result.Results) > 23 {
		t.Fatalf("Invalid array length: want %d got %d", 23, len(result.Results))
	}
}

func TestCalculatingAnnuitySums(t *testing.T) {
	request := NewRequest()
	request.Runtime = 14
	request.Creditsum = 270000
	request.Interest_rate = 4.34
	request.Initial_repayment_rate = 2
	request.Unscheduled_repayment_rate = 2

	result, err := CalculateAnnuity(request)

	if err != nil {
		t.Fatalf("Calculating Annuity failed: want error got none")
	}
	if result.Interest_sum != 121193 {
		t.Fatalf("invalid interest sum: want %d got %d", 121193, result.Interest_sum)
	}
	if result.Repayment_sum != 135577 {
		t.Fatalf("invalid repayment sum: want %d got %d", 135577, result.Repayment_sum)
	}
	if result.Unscheduled_repayment_sum != 55852 {
		t.Fatalf("invalid unscheduled repayment sum: want %d got %d", 55852, result.Unscheduled_repayment_sum)
	}
	if result.Annuity_sum != 256770 {
		t.Fatalf("invalid annuity sum: want %d got %d", 256770, result.Annuity_sum)
	}
	if result.Annuity_unsched_sum != 312622 {
		t.Fatalf("invalid annuity and unscheduled repayment sum: want %d got %d", 312622, result.Annuity_unsched_sum)
	}
}
