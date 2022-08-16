package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionRepository interface {
	SaveTransaction(transaction Transaction, creditCard CreditCard) error
	GetCreditCart(creditCard CreditCard) (CreditCard, error)
	CreateCardCard(creditCard CreditCard) error
}

type Transaction struct {
	ID string
	Amount float64
	Status string
	Description string
	Store string
	CreditCardId string
	CreatedAt time.Time
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	t.ID = uuid.NewV4().String()
	t.CreatedAt = time.Now()
	return t
}

func (t *Transaction) ProcessAndValidate(creditCart *CreditCard) {
	if t.Amount + creditCart.Balance > creditCart.Limit {
		t.Status = "rejected"
	} else {
		t.Status = "accepted"
		creditCart.Balance = creditCart.Balance + t.Amount
	}
}