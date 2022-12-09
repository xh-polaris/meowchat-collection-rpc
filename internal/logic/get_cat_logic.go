package logic

import (
	"context"
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
	id, err := strconv.ParseInt(in.CommunityId, 10, 64)
	if err != nil {
		return nil, err
	}
	conn, err := l.svcCtx.CatModel.FindOne(l.ctx, id)
	if conn.DeleteAt > conn.CreateAt {

	}
	return &pb.GetCatResp{}, nil
}
