package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-collection-rpc/errorx"
	"github.com/xh-polaris/meowchat-collection-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-collection-rpc/pb"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCatLogic {
	return &DeleteCatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCatLogic) DeleteCat(in *pb.DeleteCatReq) (*pb.DeleteCatResp, error) {
	// todo: add your logic here and delete this line
	id, err := strconv.ParseInt(in.CatId, 10, 64)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.CatModel.DeleteNotDelete(l.ctx, id)
	if err != nil {
		return nil, errorx.NoSuchCat
	}
	return &pb.DeleteCatResp{}, nil
}
