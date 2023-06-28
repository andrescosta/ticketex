package model

import "github.com/andrescosta/ticketex/func/ticket/internal/enums"

type TicketTrans struct {
	Adventure_id      string                  `json:"adventure_id"`
	User_id           string                  `json:"user_id"`
	Type              string                  `json:"type"`
	Quantity          uint                    `json:"quantity"`
	Status            enums.TransactionStatus `json:"status"`
	Credit_Card_TX_ID string                  `json:"credit_card_tx_id"`
	Tickets           []Ticket                `json:"tickets"`
}

type Ticket struct {
	Code string `json:"code"`
}
