package strings

func IsBlank(s *string) bool {
	return s == nil || len(*s) == 0
}

func IsNotBlank(s *string) bool {
	return !IsBlank(s)
}
