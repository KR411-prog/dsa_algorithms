func isValid(s string) bool {

	size := len(s)

	if size == 0 || size%2 != 0 {
		return false
	}

	check := []rune{}

	for _, c := range s {
		switch c {
		case '[':
			check = append(check, ']')
		case '(':
			check = append(check, ')')
		case '{':
			check = append(check, '}')
		default:
			check_len := len(check)
			if check_len == 0 || c != check[check_len-1]{
				return false
			}

			check = check[:check_len-1]

		}
	}

	return len(check) == 0
}