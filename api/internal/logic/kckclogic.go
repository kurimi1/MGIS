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

type KckcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKckcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KckcLogic {
	return &KckcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KckcLogic) Kckc(req *types.Request) (resp *types.KckcRsp, err error) {
	// todo: add your logic here and delete this line
	province := fmt.Sprintf("^%s[1-9]", req.Province)
	kckc := l.svcCtx.Dao.Kckc
	var results []*model.Kckc
	var total int64
	results, total, err = kckc.WithContext(l.ctx).Where(kckc.KCAAA.Regexp(province)).FindByPage(req.Size*(req.Page-1), req.Size)
	if err != nil {
		return nil, err
	}
	var kckcs []types.Kckc
	copier.Copy(&kckcs, &results)
	return &types.KckcRsp{
		Total: int(total),
		Kckcs: kckcs,
	}, nil
}
