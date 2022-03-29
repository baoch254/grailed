package userstorage

import (
	"context"

	"grailed/libs/gr-api/data-access-user/usermodel"
	common "grailed/libs/gr-api/shared-common"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
