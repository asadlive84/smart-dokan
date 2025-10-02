package db

import (
	"context"
	"log"
	"smart-dokan/usersvc/internal/application/core/domain"

	"github.com/google/uuid"
)

func (a *Adapter) Insert(user *domain.User) (uuid.UUID, error) {

	ctx := context.Background()

	var id uuid.UUID

	tx, err := a.db.Begin(ctx)
	if err != nil {
		return id, err
	}

	defer func() {
		if err != nil {
		_:
			tx.Rollback(ctx)
		}
	}()

	stmt, err := tx.Prepare(ctx, "insert_user", "INSERT INTO users(first_name, last_name, email, password, phone_number) VALUES ($1,$2,$3,$4,$5) RETURNING id")

	if err != nil {
		log.Printf("prepare can't inserted! %+v ", err)
		return id, err
	}

	err = tx.QueryRow(ctx, stmt.Name, user.FirstName, user.LastName, user.Email, user.Password, user.PhoneNumber).Scan(&id)

	if err != nil {
		log.Printf("user can't inserted! %+v ", err)
		return id, err
	}

	if err := tx.Commit(ctx); err != nil {
		log.Printf("user commit failed! %+v ", err)

		return id, err
	}

	return id, nil
}
