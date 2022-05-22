package lee

func filterFlags(str string) string {
	for i, ch := range str {
		if ch == ' ' || ch == ';' {
			return str[:i]

		}
	}
	return str
}
