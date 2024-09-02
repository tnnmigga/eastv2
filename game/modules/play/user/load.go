package user

import (
	"eastv2/game/modules/play/user/model"

	"github.com/tnnmigga/corev2/module"
	"github.com/tnnmigga/corev2/utils"
)

func LoadAsync(uid uint64, cb func(*model.Model, error)) {
	if p, ok := manager.cache[uid]; ok {
		cb(p, nil)
		return
	}
	manager.waiting[uid] = append(manager.waiting[uid], cb)
	cbs := manager.waiting[uid]
	if len(cbs) > 1 {
		return
	}
	module.Async(manager, func() (int, error) {
		return 0, nil
	}, func(n int, err error) {
		cbs := manager.waiting[uid]
		for _, cb := range cbs {
			func() {
				defer utils.RecoverPanic()
				cb(&model.Model{}, err)
			}()
		}
	})
}
