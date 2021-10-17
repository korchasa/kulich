package state

func BoolRef(b bool) *bool {
	return &b
}

func BoolRefValueString(b *bool) string {
	if b == nil {
		return "nil"
	}
	if *b {
		return "true"
	} else {
		return "false"
	}
}

func StringRefValueString(s *string) string {
	if s == nil {
		return "nil"
	}
	return *s
}

func StringRef(s string) *string {
	return &s
}
