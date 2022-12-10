package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-collection-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-collection-rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
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
	conn, err := l.svcCtx.CatModel.Insert(l.ctx, TransformModelCat(in.Cat))
	if err != nil {
		return nil, err
	}
	id, err := conn.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &pb.AddCatResp{CatId: strconv.FormatInt(id, 10)}, nil
}
