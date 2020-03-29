package repository

import (
	"database/sql"

	"github.com/comfortablynumb/goginrestapi/internal/apperror"
	"github.com/comfortablynumb/goginrestapi/internal/config"
	"github.com/comfortablynumb/goginrestapi/internal/context"
	"github.com/comfortablynumb/goginrestapi/internal/model"
	"github.com/comfortablynumb/goginrestapi/internal/repository/utils"
	"github.com/huandu/go-sqlbuilder"
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

	countOptions.WithCount(true)

	query, bindings := r.createSelectQuery(filters, &countOptions)

	row := r.db.QueryRow(query, bindings...)
	count := int64(0)

	err := row.Scan(&count)

	if err != nil {
		return count, apperror.NewDbAppError(ctx, err, UserTypeRepositorySourceName)
	}

	return count, nil
}

func (r *userTypeRepository) Find(ctx *context.RequestContext, filters *utils.UserTypeFindFilters, options *utils.UserTypeFindOptions) ([]*model.UserType, *apperror.AppError) {
	query, bindings := r.createSelectQuery(filters, options)

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
		utils.NewUserTypeFindFilters().WithNameValue(name),
		utils.NewUserTypeFindOptions().WithOffsetValue(0).WithLimitValue(1),
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
	qb := sqlbuilder.NewInsertBuilder()

	qb.InsertInto("user_types").
		Cols("name", "disabled", "created_at", "updated_at").
		Values(userType.Name, userType.Disabled, userType.CreatedAt, userType.UpdatedAt)

	query, bindings := qb.Build()

	res, err := r.db.Exec(query, bindings...)

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
	qb := sqlbuilder.NewUpdateBuilder()

	qb.Update("user_types").
		Set(qb.Assign("name", userType.Name)).
		Set(qb.Assign("disabled", userType.Disabled)).
		Set(qb.Assign("created_at", userType.CreatedAt)).
		Set(qb.Assign("updated_at", userType.UpdatedAt)).
		Where(qb.Equal("id", userType.ID))

	query, bindings := qb.Build()

	_, err := r.db.Exec(query, bindings...)

	if err != nil {
		return apperror.NewDbAppError(ctx, err, UserTypeRepositorySourceName)
	}

	return nil
}

func (r *userTypeRepository) Delete(ctx *context.RequestContext, userType *model.UserType) *apperror.AppError {
	qb := sqlbuilder.NewDeleteBuilder()

	qb.DeleteFrom("user_types").
		Where(qb.Equal("id", userType.ID))

	query, bindings := qb.Build()

	_, err := r.db.Exec(query, bindings...)

	if err != nil {
		return apperror.NewDbAppError(ctx, err, UserTypeRepositorySourceName)
	}

	return nil
}

func (r *userTypeRepository) createSelectQuery(filters *utils.UserTypeFindFilters, options *utils.UserTypeFindOptions) (string, []interface{}) {
	sb := sqlbuilder.NewSelectBuilder()

	if options.IsCount() {
		sb.Select("COUNT(u.id)")
	} else {
		sb.Select(
			"u.id",
			"u.name",
			"u.disabled",
			"u.created_at",
			"u.updated_at",
		)
	}

	sb.From(sb.As("user_types", "u"))

	if filters.GetName() != nil {
		sb.Where(sb.Equal("u.name", filters.GetNameValue()))
	}

	if !options.IsCount() {
		if options.GetSortBy() != nil && options.GetSortDir() != nil {
			sb.OrderBy(options.GetSortByValue())

			if options.IsAsc() {
				sb.Asc()
			} else {
				sb.Desc()
			}
		}

		if options.GetOffset() != nil && options.GetLimit() != nil {
			sb.Offset(options.GetOffsetValue()).Limit(options.GetLimitValue())
		}
	}

	return sb.Build()
}

// Static functions

func NewUserTypeRepository(appConfig config.AppConfig, db *sql.DB, logger *zerolog.Logger) UserTypeRepository {
	return &userTypeRepository{
		appConfig: appConfig,
		db:        db,
		logger:    logger,
	}
}
