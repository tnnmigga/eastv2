package gsconf

import "github.com/tnnmigga/corev2/conf/tables"

func init() {
	tables.Register(tables.Default[TestConf])
	tables.Check(checkTestConf)
}
