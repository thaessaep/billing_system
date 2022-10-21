package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thaessaep/billingSystem/internal/model"
	"github.com/thaessaep/billingSystem/internal/storage"
)

func TestUserReposiroty_Create(t *testing.T) {
	s, teardown := storage.TestStore(t, databaseURL)
	defer teardown("users")

	u := &model.User{
		UserId: 2,
	}
	err := s.User().AddBalance(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserReposiroty_FindById(t *testing.T) {
	s, teardown := storage.TestStore(t, databaseURL)
	defer teardown("users")

	id := 2

	_, err := s.User().FindById(id)
	assert.Error(t, err)

	s.User().AddBalance(&model.User{
		UserId:  2,
		Balance: 15,
	})

	user, err := s.User().FindById(id)

	assert.NoError(t, err)
	assert.Equal(t, 15, user.Balance)
}

func TestUserReposiroty_FindBalanceById(t *testing.T) {
	s, teardown := storage.TestStore(t, databaseURL)
	defer teardown("users")

	_, err := s.User().FindBalanceById(20)
	assert.Error(t, err)

	u := &model.User{
		UserId:  2,
		Balance: 15,
	}
	s.User().AddBalance(u)

	balance, err := s.User().FindBalanceById(u.UserId)

	assert.NoError(t, err)
	assert.Equal(t, u.Balance, balance)

	u.Balance = 20
	s.User().AddBalance(u)

	balance, err = s.User().FindBalanceById(u.UserId)

	assert.NoError(t, err)
	assert.Equal(t, 35, balance)
}
