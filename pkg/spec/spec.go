package spec

type Spec struct {
	Name         string        `hcl:"name,label"`
	Config       Config        `hcl:"config,block"`
	System       System        `hcl:"system,block"`
	Applications []Application `hcl:"application,block"`
}

func (s Spec) Diff(_ *Spec) *Diff {
	//return &Diff{
	//	System: s.System.Diff(to.System),
	//	Applications: func() []BlockDiff {
	//		//from := s.Applications
	//		//for fi, fv := range from {
	//		//	created := make([]Application, 0)
	//		//	for ti, tv := range to.Applications {
	//		//
	//		//	}
	//		//}
	//		//return BlockDiff{
	//		//}
	//	}(),
	//}
	return &Diff{}
}

type Diff struct {
	System       BlockDiff
	Applications []BlockDiff
}

type BlockDiff struct {
	Name          string
	OsOptions     OsOptionsDiff
	Users         UsersDiff
	Packages      PackagesDiff
	Directories   DirectoriesDiff
	Files         FilesDiff
	Services      ServicesDiff
	FirewallRules FirewallRulesDiff
}
