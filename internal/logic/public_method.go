package logic

import (
	"encoding/json"
	"github.com/xh-polaris/meowchat-collection-rpc/model"
	"github.com/xh-polaris/meowchat-collection-rpc/pb"
	"strconv"
	"time"
)

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
func TransformPbCat(Cat *model.Cat) *pb.Cat {
	id := strconv.FormatInt(Cat.Id, 10)
	var s []string
	json.Unmarshal([]byte(Cat.Avatars), &s)
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
		Avatars:      s,
	}
}
func TransformModelCat(Cat *pb.Cat) *model.Cat {
	id, _ := strconv.ParseInt(Cat.Id, 10, 64)
	str, _ := json.Marshal(Cat.Avatars)
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
		Avatars:      string(str),
	}
}
