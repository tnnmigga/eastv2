package gsconf

func init() {
	RegisterTable(parse[TestConf])
	RegisterCheck(checkTestConf)
}
