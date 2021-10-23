package state

type State struct {
	Name         string        `hcl:"name,label"`
	Config       Config        `hcl:"config,block"`
	System       System        `hcl:"system,block"`
	Applications []Application `hcl:"application,block"`
}
