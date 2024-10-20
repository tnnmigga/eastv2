package account

import (
	"eastv2/define"
	"eastv2/pb"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tnnmigga/corev2/infra/mdb"
	"github.com/tnnmigga/corev2/infra/rdb"
	"github.com/tnnmigga/corev2/log"
	"github.com/tnnmigga/corev2/message"
	"github.com/tnnmigga/corev2/utils"
	"golang.org/x/crypto/argon2"
)

func (m *account) register() {
	message.Response(m, onTokenAuthReq)
	message.Response(m, onAuthOrCreateAccountReq)
}

func onTokenAuthReq(body *pb.TokenAuthReq, response func(*pb.TokenAuthResp, error)) {
	uid, err := rdb.Default().GET(tokenKey(body.Token)).Value().Uint64()
	if errors.Is(err, redis.Nil) {
		response(&pb.TokenAuthResp{Code: pb.PARAM_ERROR}, nil)
	}
	resp, err := message.RequestAny[pb.UserLoginResp](define.SERV_GAME, &pb.UserLoginReq{UserID: uid})
	if err != nil {
		log.Error(err)
		response(&pb.TokenAuthResp{Code: pb.SERVER_ERROR}, nil)
		return
	}
	response(&pb.TokenAuthResp{
		Code:    pb.SUCCESS,
		UserID:  uid,
		SeverID: resp.ServerID,
	}, nil)
}

func onAuthOrCreateAccountReq(body *pb.AuthOrCreateAccountReq, response func(*pb.AuthOrCreateAccountResp, error)) {
	if len(body.Account) == 0 {
		response(&pb.AuthOrCreateAccountResp{Code: pb.PARAM_ERROR}, nil)
		return
	}
	var data []*Account
	err := mdb.Default().Table("account").Find(&data, "account = ?", body.Account).Error
	if err != nil {
		log.Error(err)
		response(&pb.AuthOrCreateAccountResp{Code: pb.SERVER_ERROR}, err)
		return
	}
	salt := []byte(body.Account)
	hash := string(argon2.IDKey([]byte(body.Password), salt, 1, 64*1024, 4, 32))
	var account *Account
	if len(data) == 0 {
		account = &Account{Account: body.Account, PwdDigest: hash}
		err = mdb.Default().Table("account").Create(account).Error
		if err != nil {
			log.Error(err)
			response(&pb.AuthOrCreateAccountResp{Code: pb.SERVER_ERROR}, err)
			return
		}
	} else {
		account = data[0]
	}
	if hash != account.PwdDigest {
		response(&pb.AuthOrCreateAccountResp{Code: pb.PARAM_ERROR}, nil)
		return
	}
	for i := 0; i < 10; i++ {
		token := utils.RandomString(32)
		key := tokenKey(token)
		ok, err := rdb.Default().SETNX(key, data[0].ID).Value().Int()
		if err != nil {
			log.Error(err)
			response(&pb.AuthOrCreateAccountResp{Code: pb.SERVER_ERROR}, err)
			return
		}
		if ok != 1 {
			continue
		}
		rdb.Default().EXPIRE(key, 30*24*time.Hour)
		response(&pb.AuthOrCreateAccountResp{Code: pb.SUCCESS, Token: token}, nil)
		return
	}
	log.Errorf("onAccountAuthReq generate token faild")
}
