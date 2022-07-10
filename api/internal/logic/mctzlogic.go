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

type MctzLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMctzLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MctzLogic {
	return &MctzLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MctzLogic) Mctz(req *types.Request) (resp *types.MctzRsp, err error) {
	// todo: add your logic here and delete this line
	province := fmt.Sprintf("^%s[1-9]", req.Province)
	mctz := l.svcCtx.Dao.Mctz
	var results []*model.Mctz
	var total int64
	results, total, err = mctz.WithContext(l.ctx).Where(mctz.KCAAA.Regexp(province)).FindByPage(req.Size*(req.Page-1), req.Size)
	if err != nil {
		return nil, err
	}
	var mctzs []types.Mctz
	copier.Copy(&mctzs, &results)
	return &types.MctzRsp{
		Total: int(total),
		Mctzs: mctzs,
	}, nil
}
