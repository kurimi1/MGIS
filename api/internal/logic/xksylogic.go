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

type XksyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewXksyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *XksyLogic {
	return &XksyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *XksyLogic) Xksy(req *types.Request) (resp *types.XksyRsp, err error) {
	// todo: add your logic here and delete this line
	province := fmt.Sprintf("^%s[1-9]", req.Province)
	xksy := l.svcCtx.Dao.Xksy
	var results []*model.Xksy
	var total int64
	results, total, err = xksy.WithContext(l.ctx).Where(xksy.KCAAA.Regexp(province)).FindByPage(req.Size*(req.Page-1), req.Size)
	if err != nil {
		return nil, err
	}
	var xksys []types.Xksy
	copier.Copy(&xksys, &results)
	return &types.XksyRsp{
		Total: int(total),
		Xksys: xksys,
	}, nil
}
