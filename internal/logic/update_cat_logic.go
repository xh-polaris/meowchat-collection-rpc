package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-collection-rpc/errorx"

	"github.com/xh-polaris/meowchat-collection-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-collection-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCatLogic {
	return &UpdateCatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// int64, err := strconv.ParseInt(string, 10, 64)
func (l *UpdateCatLogic) UpdateCat(in *pb.UpdateCatReq) (*pb.UpdateCatResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.CatModel.Update(l.ctx, svc.TransformModelCat(in.Cat))
	if err != nil {
		return nil, errorx.UpdateError
	}
	return &pb.UpdateCatResp{}, nil
}
