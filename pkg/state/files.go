package state

type Files []File

func (fs *Files) Apply(override []FileOverride) {
	for _, fo := range override {
		applied := false
		for i := range *fs {
			applied = applied || (*fs)[i].Apply(fo)
		}
		if !applied {
			*fs = append(*fs, fo.NewFile())
		}
	}
}
