package slice_diff

type Comparable interface {
	Identifier() string
	EqualityHash() string
}
