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

type KttzLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKttzLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KttzLogic {
	return &KttzLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KttzLogic) Kttz(req *types.Request) (resp *types.KttzRsp, err error) {
	// todo: add your logic here and delete this line
	province := fmt.Sprintf("^%s[1-9]", req.Province)
	kttz := l.svcCtx.Dao.Kttz
	var results []*model.Kttz
	var total int64
	results, total, err = kttz.WithContext(l.ctx).Where(kttz.KCAAA.Regexp(province)).FindByPage(req.Size*(req.Page-1), req.Size)
	if err != nil {
		return nil, err
	}
	var kttzs []types.Kttz
	copier.Copy(&kttzs, &results)
	return &types.KttzRsp{
		Total: int(total),
		Kttzs: kttzs,
	}, nil
}
