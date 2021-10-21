package state

type Option struct {
	Type  string `hcl:"type,label"`
	Name  string `hcl:"name,label"`
	Value string `hcl:"value"`
}

func (o *Option) Apply(override OptionOverride) bool {
	if o.Name != override.Name || o.Type != override.Type {
		return false
	}
	if override.Value != nil {
		o.Value = *override.Value
	}
	return true
}
