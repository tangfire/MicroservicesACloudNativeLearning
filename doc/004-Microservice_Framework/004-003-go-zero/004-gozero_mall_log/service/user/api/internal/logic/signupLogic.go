package logic

import (
	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/004-gozero_mall_log/service/user/api/internal/svc"
	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/004-gozero_mall_log/service/user/api/internal/types"
	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/004-gozero_mall_log/service/user/model"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

var secret = []byte("fireshine")

type SignupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignupLogic {
	return &SignupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignupLogic) Signup(req *types.SignupRequest) (resp *types.SignupResponse, err error) {
	// todo: add your logic here and delete this line
	// 参数校验
	if req.RePassword != req.Password {
		return nil, errors.New("两次输入的密码不一致")
	}
	logx.Debugv(req) // json.Marshal(req)
	logx.Debugf("req:%#v\n", req)
	// 把用户的注册信息保存到数据库中
	// 查询username是否已经被注册
	u, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		logx.Errorw("user_signup_UserModel.FindOneByUsername failed",
			logx.Field("err", err))
		fmt.Printf("FindOneByUsername err:%v\n", err)
		return nil, errors.New("内部失败")
	}

	if u != nil {
		return nil, errors.New("用户名已存在")
	}

	h := md5.New()
	h.Write([]byte(req.Password))
	h.Write(secret)
	passwordStr := hex.EncodeToString(h.Sum(nil))

	user := &model.User{
		UserId:   time.Now().Unix(),
		Username: req.Username,
		Password: passwordStr,
		Gender:   int64(req.Gender),
	}
	if _, err := l.svcCtx.UserModel.Insert(context.Background(), user); err != nil {
		logx.Errorf("user_signup_UserModel.Insert failed,err:%v\n", err)
		return nil, err
	}
	return &types.SignupResponse{Message: "success"}, nil
}
