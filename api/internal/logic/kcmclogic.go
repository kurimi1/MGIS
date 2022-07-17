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

type KcmcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKcmcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KcmcLogic {
	return &KcmcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KcmcLogic) Kcmc(req *types.Request) (resp *types.KcmcRsp, err error) {
	// todo: add your logic here and delete this line
	province := fmt.Sprintf("^%s[0-9]", req.Province)
	kcmc := l.svcCtx.Dao.Kcmc
	var results []*model.Kcmc
	var total int64
	results, total, err = kcmc.WithContext(l.ctx).Where(kcmc.KCAAA.Regexp(province)).FindByPage(req.Size*(req.Page-1), req.Size)
	if err != nil {
		return nil, err
	}
	var kcmcs []types.Kcmc
	copier.Copy(&kcmcs, &results)
	return &types.KcmcRsp{
		Total: int(total),
		Kcmcs: kcmcs,
	}, nil
}
