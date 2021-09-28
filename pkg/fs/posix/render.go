package posix

func (fs *Posix) render(file string, vars interface{}) (int64, error) {
	//sourceFileStat, err := os.Stat(*s.FromLocal)
	//if err != nil {
	//	return 0, err
	//}
	//
	//if !sourceFileStat.Mode().IsRegular() {
	//	return 0, fmt.Errorf("%s is not a regular file", *s.FromLocal)
	//}
	//
	//source, err := os.Open(*s.FromLocal)
	//if err != nil {
	//	return 0, err
	//}
	//defer source.Close()
	//
	//destination, err := os.Create(s.Path)
	//if err != nil {
	//	return 0, err
	//}
	//defer destination.Close()
	//nBytes, err := io.Copy(destination, source)
	//return nBytes, err
	return 0, nil
}
