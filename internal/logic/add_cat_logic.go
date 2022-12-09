package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-collection-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-collection-rpc/model"
	"github.com/xh-polaris/meowchat-collection-rpc/pb"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCatLogic {
	return &AddCatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *AddCatLogic) AddCat(in *pb.AddCatReq) (*pb.AddCatResp, error) {
	// todo: add your logic here and delete this line
	conn, err := l.svcCtx.CatModel.Insert(l.ctx, &model.Cat{
		Age:          in.Cat.Age,
		CommunityId:  in.Cat.CommunityId,
		Color:        in.Cat.Color,
		Details:      in.Cat.Details,
		Name:         in.Cat.Name,
		Popularity:   in.Cat.Popularity,
		Sex:          in.Cat.Sex,
		Status:       in.Cat.Status,
		Area:         in.Cat.Area,
		IsSnipped:    svc.BoolToInt(in.Cat.IsSnipped),
		IsSterilized: svc.BoolToInt(in.Cat.IsSterilized),
		IsDelete:     0,
	})
	if err != nil {
		return nil, err
	}
	id, err := conn.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &pb.AddCatResp{CatId: strconv.FormatInt(id, 10)}, nil
}
