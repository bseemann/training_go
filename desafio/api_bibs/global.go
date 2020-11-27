package main

const Reject = "reject"
const Accept = "accept"
const ContentTypeJSON = "application/json; charset=UTF-8"

type Payment struct {
	Name 		string		`json:"name"`
	InvoiceId 	int			`json:"invoiceId"`
	InvoiceTotal 	string		`json:"invoiceTotal"`
	Currency 		string		`json:"currency"`
}

type RefundPayment struct {
	Name 			string		`json:"name"`
	Email			string		`json:"email"`
	Address1		string		`json:"address1"`
	Address2		string		`json:"address2"`
	City			string		`json:"city"`
	State 			string		`json:"state"`
	PostCode		string		`json:"postcode"`
	Country			string		`json:"country"`
	Phone			string		`json:"phone"`
	TransactionID	string		`json:"transactionIdToRefund"`
	RefundAmount	string		`json:"refundAmount"`
	Currency		string		`json:"currencyCode"`
}

type SuccessResponse struct {
	Status			string		`json:"status"`
	TransactionId  	string		`json:"transaction_id"`
	Fee				string		`json:"fee"`
}

type DeclinedResponse struct {
	Status			string		`json:"status"`
	DeclineReason	string		`json:"declinereason"`
}
