package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-collection-rpc/errorx"

	"github.com/xh-polaris/meowchat-collection-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-collection-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetManyCatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetManyCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetManyCatLogic {
	return &GetManyCatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetManyCatLogic) GetManyCat(in *pb.GetManyCatReq) (*pb.GetManyCatResp, error) {
	// todo: add your logic here and delete this line
	conn, err := l.svcCtx.CatModel.FindManyByCommunityIdNotDelete(l.ctx, in.CommunityId, in.Skip, in.Count)
	if err != nil {
		return nil, errorx.GetManyCatError
	}
	var Cat []*pb.Cat
	for _, val := range *conn {
		Cat = append(Cat, svc.TransformPdCat(&val))
	}
	return &pb.GetManyCatResp{Cats: Cat}, nil
}
