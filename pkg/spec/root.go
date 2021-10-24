package spec

type Root struct {
	Spec Spec `hcl:"spec,block"`
}
