package emailprovider

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEmailProviderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteEmailProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEmailProviderLogic {
	return &DeleteEmailProviderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteEmailProviderLogic) DeleteEmailProvider(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.McmsRpc.DeleteEmailProvider(l.ctx, &mcms.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
