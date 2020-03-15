package repository

import (
	"database/sql"
	"fmt"

	"github.com/comfortablynumb/goginrestapi/internal/apperror"
	"github.com/comfortablynumb/goginrestapi/internal/context"
	"github.com/comfortablynumb/goginrestapi/internal/model"
	"github.com/comfortablynumb/goginrestapi/internal/repository/utils"
	"github.com/rs/zerolog"
)

// Constants

const (
	UserRepositorySourceName = "UserRepository"
)

// Interfaces

type UserRepository interface {
	Find(ctx *context.RequestContext, filters *utils.UserFindFilters, options *utils.UserFindOptions) ([]*model.User, *apperror.AppError)
	FindOneByUsername(ctx *context.RequestContext, username string) (*model.User, *apperror.AppError)
	Create(ctx *context.RequestContext, user *model.User) *apperror.AppError
	Update(ctx *context.RequestContext, user *model.User) *apperror.AppError
	Delete(ctx *context.RequestContext, user *model.User) *apperror.AppError
}

// Structs

type userRepository struct {
	db     *sql.DB
	logger *zerolog.Logger
}

func (r *userRepository) Find(ctx *context.RequestContext, filters *utils.UserFindFilters, options *utils.UserFindOptions) ([]*model.User, *apperror.AppError) {
	query := `SELECT
	u.id,
	u.username,
	u.disabled,
	u.created_at,
	u.updated_at
FROM users u
WHERE 1 = 1 `
	bindings := make([]interface{}, 0)

	if filters.GetUsername() != nil {
		query += "AND u.username = ? "
		bindings = append(bindings, filters.GetUsernameValue())
	}

	if options.GetSortBy() != nil && options.GetSortDir() != nil {
		query += fmt.Sprintf("ORDER BY %s %s", options.GetSortByValue(), options.GetSortDirValue())
	}

	if options.GetOffset() != nil && options.GetLimit() != nil {
		query += fmt.Sprintf("LIMIT %d, %d", options.GetOffsetValue(), options.GetLimitValue())
	}

	rows, err := r.db.Query(query, bindings...)

	if err != nil {
		return nil, apperror.NewDbAppError(ctx, err, UserRepositorySourceName)
	}

	res := make([]*model.User, 0)

	for rows.Next() {
		userBuilder := model.NewUserBuilder()

		ID := sql.NullInt64{}
		username := sql.NullString{}
		disabled := sql.NullBool{}
		createdAt := sql.NullTime{}
		updatedAt := sql.NullTime{}

		err = rows.Scan(&ID, &username, &disabled, &createdAt, &updatedAt)

		if err != nil {
			return nil, apperror.NewDbAppError(ctx, err, UserRepositorySourceName)
		}

		if ID.Valid {
			userBuilder.WithID(ID.Int64)
		}

		if username.Valid {
			userBuilder.WithUsername(username.String)
		}

		if disabled.Valid {
			userBuilder.WithDisabled(disabled.Bool)
		}

		if createdAt.Valid {
			userBuilder.WithCreatedAt(createdAt.Time)
		}

		if updatedAt.Valid {
			userBuilder.WithUpdatedAt(updatedAt.Time)
		}

		res = append(res, userBuilder.Build())
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.NewDbAppError(ctx, err, UserRepositorySourceName)
	}

	return res, nil
}

func (r *userRepository) FindOneByUsername(ctx *context.RequestContext, username string) (*model.User, *apperror.AppError) {
	res, err := r.Find(ctx, utils.NewUserFindFilters().WithUsername(username), utils.NewUserFindOptions().WithLimit(0, 1))

	if err != nil {
		return nil, apperror.NewDbAppError(ctx, err, UserRepositorySourceName)
	}

	if len(res) > 0 {
		return res[0], nil
	}

	return nil, nil
}

func (r *userRepository) Create(ctx *context.RequestContext, user *model.User) *apperror.AppError {
	query := `INSERT INTO users (username, disabled, created_at, updated_at) VALUES (?, ?, ?, ?)`

	res, err := r.db.Exec(query, user.Username, user.Disabled, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return apperror.NewDbAppError(ctx, err, UserRepositorySourceName)
	}

	lastInsertId, err := res.LastInsertId()

	if err != nil {
		return apperror.NewDbAppError(ctx, err, UserRepositorySourceName)
	}

	user.ID = lastInsertId

	return nil
}

func (r *userRepository) Update(ctx *context.RequestContext, user *model.User) *apperror.AppError {
	query := `UPDATE users
	SET username = ?,
		disabled = ?,
		updated_at = ?
	WHERE id = ?`

	_, err := r.db.Exec(query, user.Username, user.Disabled, user.UpdatedAt, user.ID)

	if err != nil {
		return apperror.NewDbAppError(ctx, err, UserRepositorySourceName)
	}

	return nil
}

func (r *userRepository) Delete(ctx *context.RequestContext, user *model.User) *apperror.AppError {
	query := `DELETE FROM users
	WHERE id = ?`

	_, err := r.db.Exec(query, user.ID)

	if err != nil {
		return apperror.NewDbAppError(ctx, err, UserRepositorySourceName)
	}

	return nil
}

// Static functions

func NewUserRepository(db *sql.DB, logger *zerolog.Logger) UserRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}
