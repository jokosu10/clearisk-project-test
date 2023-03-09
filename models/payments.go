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
