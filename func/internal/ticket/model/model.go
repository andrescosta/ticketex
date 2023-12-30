package model

import "github.com/andrescosta/ticketex/func/internal/ticket/enum"

type TicketTrans struct {
	AdventureID    string                 `json:"adventure_id"`
	UserID         string                 `json:"user_id"`
	Type           string                 `json:"type"`
	Quantity       uint                   `json:"quantity"`
	Status         enum.TransactionStatus `json:"status"`
	CreditCardTXID string                 `json:"credit_card_tx_id"`
	Tickets        []Ticket               `json:"tickets"`
}

type Ticket struct {
	Code string `json:"code"`
}
