package state

type Options []Option

func (o *Options) Apply(override []OptionOverride) {
	for _, oo := range override {
		applied := false
		for i := range *o {
			applied = applied || (*o)[i].Apply(oo)
		}
		if !applied {
			*o = append(*o, oo.NewOption())
		}
	}
}
