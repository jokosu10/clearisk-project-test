package models

type SumOfBalancePaymentsByPeriod struct {
	SumOfBalanceOfPaymentByPeriod string
}

type Sumbalanceperiode_response_single struct {
	Code    int                          `json:"code"`
	Status  string                       `json:"status"`
	Message string                       `json:"message"`
	Data    SumOfBalancePaymentsByPeriod `json:"data"`
}

type SumOfBalancePaymentsByStatus struct {
	Status       string
	STATUS_COUNT string
}
type Sumbalancestatus_response_single struct {
	Code    int                            `json:"code"`
	Status  string                         `json:"status"`
	Message string                         `json:"message"`
	Data    []SumOfBalancePaymentsByStatus `json:"data"`
}
