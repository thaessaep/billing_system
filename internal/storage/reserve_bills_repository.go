package storage

import (
	"errors"
	"fmt"

	"github.com/thaessaep/billingSystem/internal/model"
)

type ReserveBillsRepository struct {
	storage *Storage
}

func (rB *ReserveBillsRepository) checkForValid(r *model.ReserveBills) error {
	if r.OrderId < 0 || r.ServiceId < 0 || r.Cost <= 0 {
		return errors.New("Invalid data")
	}
	return nil
}

func (rB *ReserveBillsRepository) AddReserveBill(r *model.ReserveBills) error {
	if err := rB.checkForValid(r); err != nil {
		return err
	}

	var query string

	if exist := rB.checkForExist(r); exist {

		rB.findSuccess(r)
		if success := rB.findSuccess(r); success != nil {
			return errors.New("Already confirmed")
		}

		if r.Success != nil {
			if *r.Success == false {
				rB.storage.User().AddBalance(&model.User{
					UserId:  r.UserId,
					Balance: r.Cost,
				})
			}
		} else {
			r.Success = new(bool)
			*r.Success = true
		}

		query = fmt.Sprintf(
			"UPDATE reserve_bills SET success=%t WHERE order_id=%d AND service_id=%d AND cost=%d AND user_id=%d",
			*r.Success,
			r.OrderId,
			r.ServiceId,
			r.Cost,
			r.UserId,
		)

	} else {
		if err := rB.storage.User().ReserveBalance(&r.User, r.Cost); err != nil {
			return err
		}

		query = fmt.Sprintf(
			"INSERT INTO reserve_bills (order_id, service_id, cost, user_id) VALUES (%d, %d, %d, %d)",
			r.OrderId,
			r.ServiceId,
			r.Cost,
			r.UserId,
		)
	}

	r.Balance, _ = rB.storage.User().FindBalanceById(r.UserId)

	if err := rB.storage.db.QueryRow(query).Err(); err != nil {
		return err
	}

	return nil
}

func (rB *ReserveBillsRepository) checkForExist(r *model.ReserveBills) bool {
	var exist bool

	query := fmt.Sprintf(
		"SELECT EXISTS (SELECT * FROM reserve_bills WHERE order_id=%d AND service_id=%d AND cost=%d AND user_id=%d)",
		r.OrderId,
		r.ServiceId,
		r.Cost,
		r.UserId,
	)

	rB.storage.db.QueryRow(query).Scan(&exist)
	return exist
}

func (rB *ReserveBillsRepository) findSuccess(r *model.ReserveBills) *bool {
	var success *bool

	query := fmt.Sprintf(
		"SELECT success FROM reserve_bills WHERE order_id=%d AND service_id=%d AND cost=%d AND user_id=%d",
		r.OrderId,
		r.ServiceId,
		r.Cost,
		r.UserId,
	)

	rB.storage.db.QueryRow(query).Scan(&success)
	return success
}
