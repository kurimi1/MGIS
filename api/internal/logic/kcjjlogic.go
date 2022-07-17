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

type KcjjLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKcjjLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KcjjLogic {
	return &KcjjLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KcjjLogic) Kcjj(req *types.Request) (resp *types.KcjjRsp, err error) {
	// todo: add your logic here and delete this line
	province := fmt.Sprintf("^%s[0-9]", req.Province)
	kcjj := l.svcCtx.Dao.Kcjj
	var results []*model.Kcjj
	var total int64
	results, total, err = kcjj.WithContext(l.ctx).Where(kcjj.KCAAA.Regexp(province)).FindByPage(req.Size*(req.Page-1), req.Size)
	if err != nil {
		return nil, err
	}
	var kcjjs []types.Kcjj
	copier.Copy(&kcjjs, &results)
	return &types.KcjjRsp{
		Total: int(total),
		Kcjjs: kcjjs,
	}, nil
}
