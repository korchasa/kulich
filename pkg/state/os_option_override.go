package state

type OptionOverride struct {
	Type  string  `hcl:"type,label"`
	Name  string  `hcl:"name,label"`
	Value *string `hcl:"value"`
}

func (o *OptionOverride) NewOption() Option {
	return Option{
		Type:  o.Type,
		Name:  o.Name,
		Value: StringDeref(o.Value),
	}
}
