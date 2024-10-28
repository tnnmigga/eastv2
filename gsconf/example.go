package gsconf

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

func checkTestConf(newConfs configs) error {
	return nil
} 