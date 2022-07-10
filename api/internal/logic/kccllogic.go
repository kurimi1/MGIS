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

type KcclLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKcclLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KcclLogic {
	return &KcclLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KcclLogic) Kccl(req *types.Request) (resp *types.KcclRsp, err error) {
	// todo: add your logic here and delete this line
	province := fmt.Sprintf("^%s[1-9]", req.Province)
	kccl := l.svcCtx.Dao.Kccl
	var results []*model.Kccl
	var total int64
	results, total, err = kccl.WithContext(l.ctx).Where(kccl.KCAAA.Regexp(province)).FindByPage(req.Size*(req.Page-1), req.Size)
	if err != nil {
		return nil, err
	}
	var kqdzs []types.Kccl
	copier.Copy(&kqdzs, &results)
	return &types.KcclRsp{
		Total: int(total),
		Kccls: kqdzs,
	}, nil
}
