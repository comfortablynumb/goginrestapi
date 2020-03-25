package repository

import (
	"database/sql"
	"fmt"

	"github.com/comfortablynumb/goginrestapi/internal/apperror"
	"github.com/comfortablynumb/goginrestapi/internal/config"
	"github.com/comfortablynumb/goginrestapi/internal/context"
	"github.com/comfortablynumb/goginrestapi/internal/model"
	"github.com/comfortablynumb/goginrestapi/internal/repository/utils"
	"github.com/rs/zerolog"
)

// Constants

const (
	UserTypeRepositorySourceName = "UserTypeRepository"
)

// Interfaces

type UserTypeRepository interface {
	Count(ctx *context.RequestContext, filters *utils.UserTypeFindFilters, options *utils.UserTypeFindOptions) (int64, *apperror.AppError)
	Find(ctx *context.RequestContext, filters *utils.UserTypeFindFilters, options *utils.UserTypeFindOptions) ([]*model.UserType, *apperror.AppError)
	FindOneByName(ctx *context.RequestContext, name string) (*model.UserType, *apperror.AppError)
	Create(ctx *context.RequestContext, user *model.UserType) *apperror.AppError
	Update(ctx *context.RequestContext, user *model.UserType) *apperror.AppError
	Delete(ctx *context.RequestContext, user *model.UserType) *apperror.AppError
}

// Structs

type userTypeRepository struct {
	appConfig config.AppConfig
	db        *sql.DB
	logger    *zerolog.Logger
}

func (r *userTypeRepository) Count(ctx *context.RequestContext, filters *utils.UserTypeFindFilters, options *utils.UserTypeFindOptions) (int64, *apperror.AppError) {
	countOptions := *options

	countOptions.IsCount = true

	query, bindings := r.createQuery(filters, &countOptions)

	row := r.db.QueryRow(query, bindings...)
	count := int64(0)

	err := row.Scan(&count)

	if err != nil {
		return count, apperror.NewDbAppError(ctx, err, UserTypeRepositorySourceName)
	}

	return count, nil
}

func (r *userTypeRepository) Find(ctx *context.RequestContext, filters *utils.UserTypeFindFilters, options *utils.UserTypeFindOptions) ([]*model.UserType, *apperror.AppError) {
	query, bindings := r.createQuery(filters, options)

	rows, err := r.db.Query(query, bindings...)

	if err != nil {
		return nil, apperror.NewDbAppError(ctx, err, UserTypeRepositorySourceName)
	}

	res := make([]*model.UserType, 0)

	for rows.Next() {
		builder := model.NewUserTypeBuilder()

		ID := sql.NullInt64{}
		name := sql.NullString{}
		disabled := sql.NullBool{}
		createdAt := sql.NullTime{}
		updatedAt := sql.NullTime{}

		err = rows.Scan(&ID, &name, &disabled, &createdAt, &updatedAt)

		if err != nil {
			return nil, apperror.NewDbAppError(ctx, err, UserTypeRepositorySourceName)
		}

		if ID.Valid {
			builder.WithID(ID.Int64)
		}

		if name.Valid {
			builder.WithName(name.String)
		}

		if disabled.Valid {
			builder.WithDisabled(disabled.Bool)
		}

		if createdAt.Valid {
			builder.WithCreatedAt(createdAt.Time)
		}

		if updatedAt.Valid {
			builder.WithUpdatedAt(updatedAt.Time)
		}

		res = append(res, builder.Build())
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.NewDbAppError(ctx, err, UserTypeRepositorySourceName)
	}

	return res, nil
}

func (r *userTypeRepository) FindOneByName(ctx *context.RequestContext, name string) (*model.UserType, *apperror.AppError) {
	if name == "" {
		return nil, nil
	}

	res, err := r.Find(
		ctx,
		utils.NewUserTypeFindFiltersBuilder().WithNameValue(name).Build(),
		utils.NewUserTypeFindOptionsBuilder(r.appConfig.DefaultLimit).WithLimitValue(0, 1).Build(),
	)

	if err != nil {
		return nil, apperror.NewDbAppError(ctx, err, UserTypeRepositorySourceName)
	}

	if len(res) > 0 {
		return res[0], nil
	}

	return nil, nil
}

func (r *userTypeRepository) Create(ctx *context.RequestContext, userType *model.UserType) *apperror.AppError {
	query := `INSERT INTO user_types (name, disabled, created_at, updated_at) VALUES (?, ?, ?, ?)`

	res, err := r.db.Exec(query, userType.Name, userType.Disabled, userType.CreatedAt, userType.UpdatedAt)

	if err != nil {
		return apperror.NewDbAppError(ctx, err, UserTypeRepositorySourceName)
	}

	lastInsertId, err := res.LastInsertId()

	if err != nil {
		return apperror.NewDbAppError(ctx, err, UserTypeRepositorySourceName)
	}

	userType.ID = lastInsertId

	return nil
}

func (r *userTypeRepository) Update(ctx *context.RequestContext, userType *model.UserType) *apperror.AppError {
	query := `UPDATE user_types
	SET name = ?,
		disabled = ?,
		updated_at = ?
	WHERE id = ?`

	_, err := r.db.Exec(query, userType.Name, userType.Disabled, userType.UpdatedAt, userType.ID)

	if err != nil {
		return apperror.NewDbAppError(ctx, err, UserTypeRepositorySourceName)
	}

	return nil
}

func (r *userTypeRepository) Delete(ctx *context.RequestContext, userType *model.UserType) *apperror.AppError {
	query := `DELETE FROM user_types
	WHERE id = ?`

	_, err := r.db.Exec(query, userType.ID)

	if err != nil {
		return apperror.NewDbAppError(ctx, err, UserTypeRepositorySourceName)
	}

	return nil
}

func (r *userTypeRepository) createQuery(filters *utils.UserTypeFindFilters, options *utils.UserTypeFindOptions) (string, []interface{}) {
	query := "SELECT "

	if options.IsCount {
		query += "COUNT(u.id)"
	} else {
		query += `u.id,
		u.name,
		u.disabled,
		u.created_at,
		u.updated_at`
	}

	query += `
FROM user_types u
WHERE 1 = 1 `
	bindings := make([]interface{}, 0)

	if filters.GetName() != nil {
		query += "AND u.name = ? "
		bindings = append(bindings, filters.GetNameValue())
	}

	if !options.IsCount {
		if options.SortBy != nil && options.SortDir != nil {
			query += fmt.Sprintf("ORDER BY %s %s ", *options.SortBy, *options.SortDir)
		}

		if options.Offset != nil && options.Limit != nil {
			query += fmt.Sprintf("LIMIT %d, %d ", *options.Offset, *options.Limit)
		}
	}

	return query, bindings
}

// Static functions

func NewUserTypeRepository(appConfig config.AppConfig, db *sql.DB, logger *zerolog.Logger) UserTypeRepository {
	return &userTypeRepository{
		appConfig: appConfig,
		db:        db,
		logger:    logger,
	}
}
