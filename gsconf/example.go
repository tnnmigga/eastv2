package gsconf

import "github.com/tnnmigga/corev2/conf/tables"

type TestConf struct {
	ID   int
	Data string
}

func (t TestConf) GetID() int {
	return t.ID
}

func (t TestConf) Table() string {
	return "example"
}

func checkTestConf(newConfs tables.Tables) error {
	newConfs.Get("aaa", 1)
	return nil
} 