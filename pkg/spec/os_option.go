package spec

import (
	"fmt"
)

type OsOption struct {
	Type  string
	Name  string
	Value string
}

func (o OsOption) Identifier() string {
	return fmt.Sprintf("%s|%s", o.Type, o.Name)
}

func (o OsOption) EqualityHash() string {
	return fmt.Sprintf("%s|%s", o.Identifier(), o.Value)
}

type OsOptionsDiff struct {
	Changed []OsOption
	Removed []OsOption
}

//func DiffOsOptions(from []OsOption, to []OsOption) (created []OsOption, modified []OsOption, removed []OsOption) {
//	f := make([]OsOption, len(from))
//	copy(from, f)
//	t := make([]OsOption, len(to))
//	copy(to, t)
//	created := make([]OsOption, 0)
//	modified := make([]OsOption, 0)
//	for fi, fv := range f {
//		found := false
//		for _, tv := range t {
//			if fv.Same(&tv) {
//				found = true
//				if !fv.Equal(&tv) {
//					modified = append(modified, tv)
//				}
//			}
//		}
//		if !found {
//			created = append(created, tv)
//		}
//	}
//}
