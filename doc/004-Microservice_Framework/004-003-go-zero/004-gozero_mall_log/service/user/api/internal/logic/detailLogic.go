package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/004-gozero_mall_log/service/user/api/internal/svc"
	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/004-gozero_mall_log/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, req.UserID)
	if err != nil {
		if err != sqlx.ErrNotFound {
			logx.Errorw("UserModel.FindOneByUserId failed", logx.Field("err", err))
			return &types.DetailResponse{}, errors.New("内部错误")
		}
		return &types.DetailResponse{}, errors.New("用户不存在")

	}

	return &types.DetailResponse{
		Username: user.Username,
		Gender:   int(user.Gender),
	}, nil

}
