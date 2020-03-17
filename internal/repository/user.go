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
	u.updated_at,
	ut.id AS user_type_id,
	ut.name AS user_type_name,
	ut.disabled AS user_type_disabled,
	ut.created_at AS user_type_created_at,
	ut.updated_at AS user_type_updated_at
FROM users u
INNER JOIN user_types ut ON ut.id = u.user_type_id
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
		userTypeBuilder := model.NewUserTypeBuilder()

		ID := sql.NullInt64{}
		username := sql.NullString{}
		disabled := sql.NullBool{}
		createdAt := sql.NullTime{}
		updatedAt := sql.NullTime{}
		userTypeID := sql.NullInt64{}
		userTypeName := sql.NullString{}
		userTypeDisabled := sql.NullBool{}
		userTypeCreatedAt := sql.NullTime{}
		userTypeUpdatedAt := sql.NullTime{}

		err = rows.Scan(&ID, &username, &disabled, &createdAt, &updatedAt, &userTypeID, &userTypeName, &userTypeDisabled, &userTypeCreatedAt, &userTypeUpdatedAt)

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

		if userTypeID.Valid {
			userTypeBuilder.WithID(userTypeID.Int64)
		}

		if userTypeName.Valid {
			userTypeBuilder.WithName(userTypeName.String)
		}

		if userTypeDisabled.Valid {
			userTypeBuilder.WithDisabled(userTypeDisabled.Bool)
		}

		if userTypeCreatedAt.Valid {
			userTypeBuilder.WithCreatedAt(userTypeCreatedAt.Time)
		}

		if userTypeUpdatedAt.Valid {
			userTypeBuilder.WithUpdatedAt(userTypeUpdatedAt.Time)
		}

		userBuilder.WithUserType(*userTypeBuilder.Build())

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
	query := `INSERT INTO users (username, user_type_id, disabled, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`

	res, err := r.db.Exec(query, user.Username, user.UserType.ID, user.Disabled, user.CreatedAt, user.UpdatedAt)

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
		user_type_id = ?,
		disabled = ?,
		updated_at = ?
	WHERE id = ?`

	_, err := r.db.Exec(query, user.Username, user.UserType.ID, user.Disabled, user.UpdatedAt, user.ID)

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
