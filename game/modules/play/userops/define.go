package userops

import "fmt"

const ()

func groupKey(uid uint64) string {
	return fmt.Sprintf("mongo.userdata.%d", uid)
}
