package userops

import (
	"context"
	"eastv2/game/modules/play/userops/userdata"
	"time"

	"github.com/tnnmigga/corev2/conc"
	"github.com/tnnmigga/corev2/infra/mgdb"
	"github.com/tnnmigga/corev2/log"
	"github.com/tnnmigga/corev2/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func LoadAsync(uid uint64, cb func(*userdata.Entity, error)) {
	if p, ok := manager.cache[uid]; ok {
		cb(p, nil)
		return
	}
	manager.waiting[uid] = append(manager.waiting[uid], cb)
	cbs := manager.waiting[uid]
	if len(cbs) > 1 {
		return
	}
	conc.Async(manager, func() (*userdata.Entity, error) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		result := mgdb.Default().Collection("userdata").FindOne(ctx, bson.M{"_id": uid})
		var user userdata.Entity
		err := result.Decode(&user)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}, func(user *userdata.Entity, err error) {
		if err != nil {
			log.Errorf("LoadAsync load error %v", err)
		} else {
			manager.cache[uid] = user
		}
		log.Debugf("load user %d %v", uid, err)
		cbs := manager.waiting[uid]
		for _, cb := range cbs {
			func() {
				defer utils.RecoverPanic()
				cb(user, err)
			}()
		}
	}, groupKey(uid))
}
