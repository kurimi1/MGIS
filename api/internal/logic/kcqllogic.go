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

type KcqlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKcqlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KcqlLogic {
	return &KcqlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KcqlLogic) Kcql(req *types.Request) (resp *types.KcqlRsp, err error) {
	// todo: add your logic here and delete this line
	province := fmt.Sprintf("^%s[1-9]", req.Province)
	kcql := l.svcCtx.Dao.Kcql
	var results []*model.Kcql
	var total int64
	results, total, err = kcql.WithContext(l.ctx).Where(kcql.KCAAA.Regexp(province)).FindByPage(req.Size*(req.Page-1), req.Size)
	if err != nil {
		return nil, err
	}
	var kcqls []types.Kcql
	copier.Copy(&kcqls, &results)
	return &types.KcqlRsp{
		Total: int(total),
		Kcqls: kcqls,
	}, nil
}
