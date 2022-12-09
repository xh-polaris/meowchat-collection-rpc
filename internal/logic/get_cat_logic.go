package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-collection-rpc/errorx"
	"strconv"

	"github.com/xh-polaris/meowchat-collection-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-collection-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCatLogic {
	return &GetCatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// int64, err := strconv.ParseInt(string, 10, 64)
func (l *GetCatLogic) GetCat(in *pb.GetCatReq) (*pb.GetCatResp, error) {
	// todo: add your logic here and delete this line
	id, err := strconv.ParseInt(in.CatId, 10, 64)
	if err != nil {
		return nil, err
	}
	conn, err := l.svcCtx.CatModel.FindOneNotDelete(l.ctx, id)
	if err != nil {
		return nil, errorx.NoSuchCat
	}
	return &pb.GetCatResp{Cat: svc.TransformPdCat(conn)}, nil
}
