package ruchki

type FsCommon struct {
	
}

func (fs *FsCommon) Setup(c *FsDriverConfig) error {
	panic("implement me")
}

func (fs *FsCommon) ApplyFile(f *File) error {
	panic("implement me")
}

func (fs *FsCommon) ApplyDir(d *Directory) error {
	panic("implement me")
}
