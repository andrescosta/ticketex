package service

import (
	"github.com/andrescosta/ticketex/func/internal/config"
	"github.com/andrescosta/ticketex/func/internal/ticket/entity"
	"github.com/andrescosta/ticketex/func/internal/ticket/enums"
	"github.com/andrescosta/ticketex/func/internal/ticket/repository"
	"go.jetpack.io/typeid"
)

type Ticket struct {
	repo *repository.Ticket
}

func New(config config.Config) (*Ticket, error) {
	repo, err := repository.New(config)
	if err != nil {
		return nil, err
	}
	return &Ticket{
		repo: repo,
	}, nil
}

func (r *Ticket) GenerateTickets(ticketTrans entity.TicketTrans) (entity.TicketTrans, error) {
	// This method will validate with the external payment platform the TX ID
	ticketTrans.Status = enums.Validated
	ticketTrans.Tickets = make([]entity.Ticket, ticketTrans.Quantity)
	prefix := "t"
	for i := 0; i < int(ticketTrans.Quantity); i++ {
		id, err := typeid.New(prefix)
		if err != nil {
			return entity.TicketTrans{}, err
		}
		ticketTrans.Tickets[i] = entity.Ticket{
			Code:        id.String(),
			AdventureID: ticketTrans.AdventureID,
			UserID:      ticketTrans.UserID,
			Type:        ticketTrans.Type,
		}
	}
	err := r.repo.NewTicketTrans(ticketTrans)
	if err != nil {
		return entity.TicketTrans{}, err
	}
	return ticketTrans, nil
}

func (r *Ticket) UpdateTicketTrans(ticketTrans entity.TicketTrans) error {
	return r.repo.UpdateTicketTrans(ticketTrans)
}

func (r *Ticket) GetTicketTrans(ticketTrans entity.TicketTrans) (entity.TicketTrans, error) {
	return r.repo.GetTicketTrans(ticketTrans)
}
