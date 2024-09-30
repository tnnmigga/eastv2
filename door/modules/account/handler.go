package account

import (
	"eastv2/pb"

	"github.com/tnnmigga/corev2/log"
	"github.com/tnnmigga/corev2/message"
)

func (m *account) initHandler() {
	message.Response(m, m.onTokenAuthReq)
}

func (m *account) onTokenAuthReq(req *pb.TokenAuthReq, resp func(any, error)) {
	log.Debugf("onTokenAuthReq: %v", req)
	resp(&pb.TokenAuthResp{
		Code:    pb.SUCCESS,
		UserID:  1,
		SeverID: 1001,
	}, nil)
}
