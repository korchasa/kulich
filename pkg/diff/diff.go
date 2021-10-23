package diff

import (
	"fmt"
	"reflect"
)

func interfaceSlice(slice interface{}) ([]Comparable, error) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	if s.IsNil() {
		return nil, nil
	}

	ret := make([]Comparable, s.Len())
	for i := 0; i < s.Len(); i++ {
		//c, ok := s.Index(i).Interface().(Comparable)
		//if !ok {
		//	return nil, fmt.Errorf("can't cast %s to Compatible", s.Index(i).Type())
		//}
		ret[i] = s.Index(i).Interface().(Comparable)
	}

	return ret, nil
}

func Diff(from, to interface{}) (changed, removed []interface{}, err error) {
	fromSlice, err := interfaceSlice(from)
	if err != nil {
		return nil, nil, fmt.Errorf("can't convert `from` slice: %w", err)
	}
	toSlice, err := interfaceSlice(to)
	if err != nil {
		return nil, nil, fmt.Errorf("can't convert `to` slice: %w", err)
	}

	for _, tv := range toSlice {
		found := false
		for fi, fv := range fromSlice {
			if fv != nil && fv.Identifier() == tv.Identifier() {
				if fv.EqualityHash() != tv.EqualityHash() {
					changed = append(changed, tv)
				}
				fromSlice[fi] = nil
				found = true
			}
		}
		if !found {
			changed = append(changed, tv)
		}
	}

	for _, v := range fromSlice {
		if v != nil {
			removed = append(removed, v)
		}
	}

	return
}
