package logic

import (
	"context"
	"fmt"

	"LGIS/api/internal/svc"
	"LGIS/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	// 根据省份和矿种查询
	// 省份取前两位用正则
	province := fmt.Sprintf("^%s[1-9]", req.Province)
	kcdj := l.svcCtx.Dao.Kcdj
	// 判断mine是否为空
	results, total, err := kcdj.WithContext(l.ctx).Where(kcdj.KCAAA.Regexp(province), kcdj.KCC.Eq(req.Mine)).FindByPage(req.Size*(req.Page-1), req.Size)
	if err != nil {
		return nil, err
	}
	var kcdjs []types.Kcdj
	copier.Copy(&kcdjs, &results)

	return &types.Response{
		Total: int(total),
		Kcdjs: kcdjs,
	}, nil
}
