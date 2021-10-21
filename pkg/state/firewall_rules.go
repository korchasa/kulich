package state

type FirewallRules []FirewallRule

func (r *FirewallRules) Apply(override []FirewallRuleOverride) {
	for _, oo := range override {
		applied := false
		for i := range *r {
			applied = applied || (*r)[i].Apply(oo)
		}
		if !applied {
			*r = append(*r, oo.NewRule())
		}
	}
}
