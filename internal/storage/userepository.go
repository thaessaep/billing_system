package storage

import (
	"errors"
	"fmt"

	"github.com/thaessaep/billingSystem/internal/model"
)

type UserRepository struct {
	storage *Storage
}

func (ur *UserRepository) AddBalance(u *model.User) error {
	var query string

	if u.Balance < 0 {
		return errors.New("Balance is invalid")
	}
	balance, err := ur.FindBalanceById(u.UserId)
	if err != nil {
		query = fmt.Sprintf("INSERT INTO users (user_id, balance) VALUES (%d, %d) RETURNING user_id", u.UserId, u.Balance)
	} else {
		query = fmt.Sprintf("UPDATE users SET balance=%d WHERE user_id=%d RETURNING user_id", u.Balance+balance, u.UserId)
	}

	if err := ur.storage.db.QueryRow(query).Scan(&u.UserId); err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) ReserveBalance(u *model.User, cost int) error {
	var query string

	balance, err := ur.FindBalanceById(u.UserId)

	if err != nil {
		return err
	} else if balance < cost {
		return errors.New(fmt.Sprintf("Balance less cost, Balance = %d, Cost = %d", balance, cost))
	} else {
		query = fmt.Sprintf("UPDATE users SET balance=%d WHERE user_id=%d RETURNING balance", balance-cost, u.UserId)
	}

	if err := ur.storage.db.QueryRow(query).Scan(&u.Balance); err != nil {
		return err
	}

	return nil
}

// Get user balance by id
func (ur *UserRepository) FindById(id int) (*model.User, error) {
	u := &model.User{}
	query := "SELECT user_id, balance FROM users WHERE user_id=$1"
	if err := ur.storage.db.QueryRow(query, id).Scan(&u.UserId, &u.Balance); err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *UserRepository) FindBalanceById(id int) (int, error) {
	u, err := ur.FindById(id)
	if err != nil {
		return 0, err
	}

	return u.Balance, nil
}
