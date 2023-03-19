package user

import (
	"context"
	"database/sql"
)

type UserService interface {
	Load(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewUserService(db *sql.DB, repository UserRepository) UserService {
	return &userService{db: db, repository: repository}
}

type userService struct {
	db         *sql.DB
	repository UserRepository
}

func (s *userService) Load(ctx context.Context, id string) (*User, error) {
	return s.repository.Load(ctx, id)
}
func (s *userService) Create(ctx context.Context, user *User) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return -1, nil
	}
	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.repository.Create(ctx, user)
	if err != nil {
		er := tx.Rollback()
		if er != nil {
			return -1, er
		}
		return -1, err
	}
	err = tx.Commit()
	return res, err
}
func (s *userService) Update(ctx context.Context, user *User) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return -1, nil
	}
	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.repository.Update(ctx, user)
	if err != nil {
		er := tx.Rollback()
		if er != nil {
			return -1, er
		}
		return -1, err
	}
	err = tx.Commit()
	return res, err
}
func (s *userService) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return -1, nil
	}
	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.repository.Patch(ctx, user)
	if err != nil {
		er := tx.Rollback()
		if er != nil {
			return -1, er
		}
		return -1, err
	}
	err = tx.Commit()
	return res, err
}
func (s *userService) Delete(ctx context.Context, id string) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return -1, nil
	}
	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.repository.Delete(ctx, id)
	if err != nil {
		er := tx.Rollback()
		if er != nil {
			return -1, er
		}
		return -1, err
	}
	err = tx.Commit()
	return res, err
}
