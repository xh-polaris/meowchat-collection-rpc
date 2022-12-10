package logic

import (
	"context"
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
	conn, err := l.svcCtx.CatModel.FindManyByCommunityIdNotDelete(l.ctx, in.CommunityId, in.Skip, in.Count)
	if err != nil {
		return nil, err
	}
	var Cat []*pb.Cat
	for _, val := range conn {
		Cat = append(Cat, TransformPbCat(val))
	}
	return &pb.GetManyCatResp{Cats: Cat}, nil
}
