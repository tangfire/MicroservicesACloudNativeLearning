package logic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

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

	// 根据userID查询数据库返回用户信息
	one, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, in.UserID)
	// 1. 查询数据库失败
	// 2. 根据userID查不到用户
	if errors.Is(err, sqlx.ErrNotFound) {
		return nil, errors.New("user not found")
	}
	if err != nil {
		logx.Errorw("user.rpc.GetUser FindOneByUserId failed", logx.Field("err", err))
		return nil, errors.New("查询失败")
	}

	return &user.GetUserResp{
		UserID:   one.UserId,
		Username: one.Username,
		Gender:   one.Gender,
	}, nil
}
