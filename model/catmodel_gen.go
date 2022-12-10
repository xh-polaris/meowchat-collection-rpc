// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	catFieldNames          = builder.RawFieldNames(&Cat{})
	catRows                = strings.Join(catFieldNames, ",")
	catRowsExpectAutoSet   = strings.Join(stringx.Remove(catFieldNames, "`id`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), ",")
	catRowsWithPlaceHolder = strings.Join(stringx.Remove(catFieldNames, "`id`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), "=?,") + "=?"

	cacheMeowchatCollectionRpcCatIdPrefix = "cache:meowchatCollectionRpc:cat:id:"
)

type (
	catModel interface {
		Insert(ctx context.Context, data *Cat) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Cat, error)
		Update(ctx context.Context, data *Cat) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCatModel struct {
		sqlc.CachedConn
		table string
	}

	Cat struct {
		Id           int64     `db:"id"`
		CreateAt     time.Time `db:"create_at"` // 创建时间
		DeleteAt     time.Time `db:"delete_at"`
		Age          string    `db:"age"`
		CommunityId  string    `db:"community_id"`
		Color        string    `db:"color"`
		Details      string    `db:"details"`
		Name         string    `db:"name"`
		Popularity   int64     `db:"popularity"`
		Sex          string    `db:"sex"`
		Status       int64     `db:"status"`
		Area         string    `db:"area"`
		IsSnipped    int64     `db:"is_snipped"`
		IsSterilized int64     `db:"is_sterilized"`
		IsDelete     int64     `db:"is_delete"`
		Avatars      string    `db:"avatars"` // 图片
	}
)

func newCatModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultCatModel {
	return &defaultCatModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`cat`",
	}
}

func (m *defaultCatModel) Delete(ctx context.Context, id int64) error {
	meowchatCollectionRpcCatIdKey := fmt.Sprintf("%s%v", cacheMeowchatCollectionRpcCatIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, meowchatCollectionRpcCatIdKey)
	return err
}

func (m *defaultCatModel) FindOne(ctx context.Context, id int64) (*Cat, error) {
	meowchatCollectionRpcCatIdKey := fmt.Sprintf("%s%v", cacheMeowchatCollectionRpcCatIdPrefix, id)
	var resp Cat
	err := m.QueryRowCtx(ctx, &resp, meowchatCollectionRpcCatIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", catRows, m.table)
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

func (m *defaultCatModel) Insert(ctx context.Context, data *Cat) (sql.Result, error) {
	meowchatCollectionRpcCatIdKey := fmt.Sprintf("%s%v", cacheMeowchatCollectionRpcCatIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, catRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeleteAt, data.Age, data.CommunityId, data.Color, data.Details, data.Name, data.Popularity, data.Sex, data.Status, data.Area, data.IsSnipped, data.IsSterilized, data.IsDelete, data.Avatars)
	}, meowchatCollectionRpcCatIdKey)
	return ret, err
}

func (m *defaultCatModel) Update(ctx context.Context, data *Cat) error {
	meowchatCollectionRpcCatIdKey := fmt.Sprintf("%s%v", cacheMeowchatCollectionRpcCatIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, catRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeleteAt, data.Age, data.CommunityId, data.Color, data.Details, data.Name, data.Popularity, data.Sex, data.Status, data.Area, data.IsSnipped, data.IsSterilized, data.IsDelete, data.Avatars, data.Id)
	}, meowchatCollectionRpcCatIdKey)
	return err
}

func (m *defaultCatModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMeowchatCollectionRpcCatIdPrefix, primary)
}

func (m *defaultCatModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", catRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultCatModel) tableName() string {
	return m.table
}
