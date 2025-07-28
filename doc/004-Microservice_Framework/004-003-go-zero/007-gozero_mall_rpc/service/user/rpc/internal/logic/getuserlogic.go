package logic

import (
	"context"

	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/007-gozero_mall_rpc/service/user/rpc/internal/svc"
	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/007-gozero_mall_rpc/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserResp{}, nil
}
