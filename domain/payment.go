package domain

import "time"

type Payment struct {
	ID            string    `json:"id"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	DateOfPayment time.Time `json:"dateOfPayment"`
	IdUser        string    `json:"idUser"`
}
