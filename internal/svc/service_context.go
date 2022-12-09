package svc

import (
	"github.com/xh-polaris/meowchat-collection-rpc/internal/config"
	"github.com/xh-polaris/meowchat-collection-rpc/model"
	"github.com/xh-polaris/meowchat-collection-rpc/pb"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
	"time"
)

type ServiceContext struct {
	Config          config.Config
	CatModel        model.CatModel
	CatAvatarsModel model.CatAvatarsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:          c,
		CatModel:        model.NewCatModel(conn, c.CacheRedis),
		CatAvatarsModel: model.NewCatAvatarsModel(conn, c.CacheRedis),
	}
}
func BoolToInt(a bool) int64 {
	if a {
		return 1
	}
	return 0
}
func IntToBool(a int64) bool {
	if a == 1 {
		return true
	}
	return false
}

// int64, err := strconv.ParseInt(string, 10, 64)
func TransformPdCat(Cat *model.Cat) *pb.Cat {
	id := strconv.FormatInt(Cat.Id, 10)
	return &pb.Cat{
		Id:           id,
		CreateAt:     Cat.CreateAt.Unix(),
		Age:          Cat.Age,
		CommunityId:  Cat.CommunityId,
		Color:        Cat.Color,
		Details:      Cat.Details,
		Name:         Cat.Name,
		Popularity:   Cat.Popularity,
		Sex:          Cat.Sex,
		Status:       Cat.Status,
		Area:         Cat.Area,
		IsSnipped:    IntToBool(Cat.IsSnipped),
		IsSterilized: IntToBool(Cat.IsSterilized),
	}
}
func TransformModelCat(Cat *pb.Cat) *model.Cat {
	id, _ := strconv.ParseInt(Cat.Id, 10, 64)
	return &model.Cat{
		Id:           id,
		CreateAt:     time.Unix(Cat.CreateAt, 0),
		Age:          Cat.Age,
		CommunityId:  Cat.CommunityId,
		Color:        Cat.Color,
		Details:      Cat.Details,
		Name:         Cat.Name,
		Popularity:   Cat.Popularity,
		Sex:          Cat.Sex,
		Status:       Cat.Status,
		Area:         Cat.Area,
		IsSnipped:    BoolToInt(Cat.IsSnipped),
		IsSterilized: BoolToInt(Cat.IsSterilized),
	}
}
