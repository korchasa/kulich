package diff

type Comparable interface {
	Identifier() string
	EqualityHash() string
}
