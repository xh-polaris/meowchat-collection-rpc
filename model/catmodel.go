package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var (
	_                                           CatModel = (*customCatModel)(nil)
	cacheMeowchatCollectionRpcCommunityIdPrefix          = "cache:meowchatCollectionRpc:cat:CommunityId:"
)

type (
	// CatModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCatModel.
	CatModel interface {
		catModel
		DeleteNotDelete(ctx context.Context, id int64) error
		FindOneNotDelete(ctx context.Context, id int64) (*Cat, error)
		FindManyByCommunityIdNotDelete(ctx context.Context, CommunityId string, skip int64, count int64) (*[]Cat, error)
	}

	customCatModel struct {
		*defaultCatModel
	}
)

// NewCatModel returns a model for the database table.
func NewCatModel(conn sqlx.SqlConn, c cache.CacheConf) CatModel {
	return &customCatModel{
		defaultCatModel: newCatModel(conn, c),
	}
}

func (m *defaultCatModel) DeleteNotDelete(ctx context.Context, id int64) error {
	meowchatCollectionRpcCatIdKey := fmt.Sprintf("%s%v", cacheMeowchatCollectionRpcCatIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set is_delete= 1 delete_at =? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, time.Now(), id)
	}, meowchatCollectionRpcCatIdKey)
	return err
}

func (m *defaultCatModel) FindOneNotDelete(ctx context.Context, id int64) (*Cat, error) {
	meowchatCollectionRpcCatIdKey := fmt.Sprintf("%s%v", cacheMeowchatCollectionRpcCatIdPrefix, id)
	var resp Cat
	err := m.QueryRowCtx(ctx, &resp, meowchatCollectionRpcCatIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and is_delete=0 limit 1", catRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCatModel) FindManyByCommunityIdNotDelete(ctx context.Context, CommunityId string, skip int64, count int64) (*[]Cat, error) {
	cacheMeowchatCollectionRpcCommunityId := fmt.Sprintf("%s%s:%v:%v", cacheMeowchatCollectionRpcCommunityIdPrefix, CommunityId, skip, count)
	var resp []Cat
	err := m.QueryRowCtx(ctx, &resp, cacheMeowchatCollectionRpcCommunityId, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where  community_id = ? and is_delete=0 limit %v,%v", catRows, m.table, skip, count)
		return conn.QueryRowCtx(ctx, v, query, CommunityId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
