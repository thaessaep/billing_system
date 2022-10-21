package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thaessaep/billingSystem/internal/model"
	"github.com/thaessaep/billingSystem/internal/storage"
)

func TestReserveBills_AddReserveBill(t *testing.T) {
	s, teardown := storage.TestStore(t, databaseURL)
	defer teardown("reserve_bills")

	u := &model.User{
		UserId:  3,
		Balance: 60,
	}
	s.User().AddBalance(u)

	rB := &model.ReserveBills{
		OrderId:   1,
		ServiceId: 2,
		Cost:      50,
		User: model.User{
			UserId: 3,
		},
	}

	err := s.ReserveBills().AddReserveBill(rB)

	assert.Nil(t, rB.Success)
	assert.NoError(t, err)
	assert.NotNil(t, rB)
	assert.Equal(t, 10, rB.Balance)

	err = s.ReserveBills().AddReserveBill(rB)
	assert.NoError(t, err)
	assert.NotNil(t, rB.Success, "Success is nil")
	assert.Equal(t, true, *rB.Success)

	rB.Cost = 80
	err = s.ReserveBills().AddReserveBill(rB)
	assert.Error(t, err)

	rB.Cost = 5
	rB.OrderId = -4
	err = s.ReserveBills().AddReserveBill(rB)
	assert.Error(t, err)
}
