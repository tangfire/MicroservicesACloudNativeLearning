package logic

import (
	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/003-gozero_mall_cache/service/user/api/internal/svc"
	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/003-gozero_mall_cache/service/user/api/internal/types"
	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/003-gozero_mall_cache/service/user/model"
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
	fmt.Printf("req:%#v\n", req)
	// 把用户的注册信息保存到数据库中
	// 查询username是否已经被注册
	u, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
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
		return nil, err
	}
	return &types.SignupResponse{Message: "success"}, nil
}
