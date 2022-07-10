package logic

import (
	"context"
	"fmt"

	"LGIS/api/internal/svc"
	"LGIS/api/internal/types"
	"LGIS/model"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type KcjsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKcjsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KcjsLogic {
	return &KcjsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KcjsLogic) Kcjs(req *types.Request) (resp *types.KcjsRsp, err error) {
	// todo: add your logic here and delete this line
	province := fmt.Sprintf("^%s[1-9]", req.Province)
	kcjs := l.svcCtx.Dao.Kcj
	var results []*model.Kcj
	var total int64
	results, total, err = kcjs.WithContext(l.ctx).Where(kcjs.KCAAA.Regexp(province)).FindByPage(req.Size*(req.Page-1), req.Size)
	if err != nil {
		return nil, err
	}
	var kcjss []types.Kcj
	copier.Copy(&kcjss, &results)
	return &types.KcjsRsp{
		Total: int(total),
		Kcjs:  kcjss,
	}, nil
}
