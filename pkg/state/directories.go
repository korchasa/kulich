package state

type Directories []Directory

func (d *Directories) Apply(override []DirectoryOverride) {
	for _, oo := range override {
		applied := false
		for i := range *d {
			applied = applied || (*d)[i].Apply(oo)
		}
		if !applied {
			*d = append(*d, oo.NewDirectory())
		}
	}
}
