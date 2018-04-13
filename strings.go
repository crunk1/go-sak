package sak

func StrOr(s string, ss ...string) string {
	if s != "" || len(ss) == 0 {
		return s
	}
	for _, s := range ss {
		if s != "" {
			return s
		}
	}
	return ""
}