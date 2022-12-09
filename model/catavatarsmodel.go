package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CatAvatarsModel = (*customCatAvatarsModel)(nil)

type (
	// CatAvatarsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCatAvatarsModel.
	CatAvatarsModel interface {
		catAvatarsModel
	}

	customCatAvatarsModel struct {
		*defaultCatAvatarsModel
	}
)

// NewCatAvatarsModel returns a model for the database table.
func NewCatAvatarsModel(conn sqlx.SqlConn, c cache.CacheConf) CatAvatarsModel {
	return &customCatAvatarsModel{
		defaultCatAvatarsModel: newCatAvatarsModel(conn, c),
	}
}
