package repository

import (
	"context"
	"gh5-backend/pkg/ctxval"
	res "gh5-backend/pkg/utils/response"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	conn       *gorm.DB
	entity     T
	entityName string
}

func NewRepository[T any](conn *gorm.DB, entity T, entityName string) *Repository[T] {
	return &Repository[T]{
		conn,
		entity,
		entityName,
	}
}

func (r *Repository[T]) getConn() *gorm.DB {
	return r.conn
}

func (m *Repository[T]) checkTrx(ctx context.Context) {
	trx := ctxval.GetTrxValue(ctx)
	if trx != nil {
		m.conn = trx.Db
	}
	m.conn = m.conn.WithContext(ctx)
}

func (m *Repository[T]) maskError(err error) error {
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		if pqErr, ok := err.(*pgconn.PgError); ok {
			if pqErr.Code == "23505" {
				return res.ErrorBuilder(&res.ErrorConstant.DuplicateEntity, err)
			}
		}

		return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
	}

	return nil
}

func (r *Repository[T]) Create(ctx context.Context, data T) (T, error) {
	r.checkTrx(ctx)
	query := r.conn.Model(r.entity)
	err := query.Create(&data).Error
	return data, r.maskError(err)
}

func (r *Repository[T]) Delete(ctx context.Context, entity *T) error {
	r.checkTrx(ctx)
	err := r.conn.Model(r.entity).Delete(entity).Error
	return r.maskError(err)
}

func (r *Repository[T]) CountByID(ctx context.Context, id any) (int64, error) {
	r.checkTrx(ctx)
	var total int64
	query := r.conn.Model(r.entity)
	err := query.Model(new(T)).Where("id = ?", id).Count(&total).Error
	return total, r.maskError(err)
}

func (r *Repository[T]) FindByID(ctx context.Context, id any) (*T, error) {
	r.checkTrx(ctx)
	query := r.conn.Model(r.entity)
	result := new(T)
	err := query.Where("id", id).First(result).Error
	if err != nil {
		return nil, r.maskError(err)
	}
	return result, nil
}

func (r *Repository[T]) Find(ctx context.Context) ([]T, error) {
	r.checkTrx(ctx)
	query := r.conn.Model(r.entity)
	result := new([]T)
	err := query.Find(result).Error
	if err != nil {
		return nil, r.maskError(err)
	}
	return *result, nil
}
