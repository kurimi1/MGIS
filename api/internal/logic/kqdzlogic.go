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

type KqdzLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKqdzLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KqdzLogic {
	return &KqdzLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KqdzLogic) Kqdz(req *types.Request) (resp *types.KqdzRsp, err error) {
	// todo: add your logic here and delete this line
	// 根据省份查询
	// 省份取前两位用正则
	province := fmt.Sprintf("^%s[1-9]", req.Province)
	kqdz := l.svcCtx.Dao.Kqdz
	var results []*model.Kqdz
	var total int64
	results, total, err = kqdz.WithContext(l.ctx).Where(kqdz.KCAAA.Regexp(province)).FindByPage(req.Size*(req.Page-1), req.Size)
	if err != nil {
		return nil, err
	}
	var kqdzs []types.Kqdz
	copier.Copy(&kqdzs, &results)
	return &types.KqdzRsp{
		Total: int(total),
		Kqdzs: kqdzs,
	}, nil
}
