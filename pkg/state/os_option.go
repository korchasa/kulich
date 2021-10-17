package state

type Option struct {
	Name  string `hcl:"name,label"`
	Value string `hcl:"value"`
}

type OptionOverride struct {
	Name  string  `hcl:"name,label"`
	Value *string `hcl:"value"`
}
