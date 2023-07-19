package service

import (
	"github.com/andrescosta/ticketex/func/common"
	"github.com/andrescosta/ticketex/func/ticket/internal/entity"
	"github.com/andrescosta/ticketex/func/ticket/internal/enums"
	"github.com/andrescosta/ticketex/func/ticket/internal/repository"
	"go.jetpack.io/typeid"
)

type ITicketSvc interface {
	GenerateTickets(ticketTrans entity.TicketTrans) (entity.TicketTrans, error)
	UpdateTicketTrans(ticketTrans entity.TicketTrans) error
	GetTicketTrans(ticketTrans entity.TicketTrans) (entity.TicketTrans, error)
}

type TicketSvc struct {
	repo repository.ITicket
}

func Init(config config.Config) (ITicketSvc, error) {
	if repo, err := repository.Init(config); err != nil {
		return nil, err
	} else {
		return &TicketSvc{
			repo: repo,
		}, nil
	}
}

func (r *TicketSvc) GenerateTickets(ticketTrans entity.TicketTrans) (entity.TicketTrans, error) {
	// This method will validate with the external payment platform the TX ID
	ticketTrans.Status = enums.Validated
	ticketTrans.Tickets = make([]entity.Ticket, ticketTrans.Quantity)
	prefix := "t"
	for i := 0; i < int(ticketTrans.Quantity); i++ {
		id, err := typeid.New(prefix)
		if err == nil {
			ticketTrans.Tickets[i] = entity.Ticket{
				Code:         id.String(),
				Adventure_id: ticketTrans.Adventure_id,
				User_id:      ticketTrans.User_id,
				Type:         ticketTrans.Type,
			}
		} else {
			return entity.TicketTrans{}, err
		}
	}
	err := r.repo.NewTicketTrans(ticketTrans)
	if err != nil {
		return entity.TicketTrans{}, err
	} else {
		return ticketTrans, nil
	}
}

func (r *TicketSvc) UpdateTicketTrans(ticketTrans entity.TicketTrans) error {
	return r.repo.UpdateTicketTrans(ticketTrans)
}

func (r *TicketSvc) GetTicketTrans(ticketTrans entity.TicketTrans) (entity.TicketTrans, error) {
	return r.repo.GetTicketTrans(ticketTrans)
}
