package main

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}

	ret := 0x0

	for i := 0; i < len(password); i++ {
		if isSmall(password[i]) {
			ret = ret | 0x0001
		} else if isBig(password[i]) {
			ret = ret | 0x0010
		} else if isNumber(password[i]) {
			ret = ret | 0x0100
		} else if isSpecial(password[i]) {
			ret = ret | 0x1000
		}

		if i > 0 {
			if password[i] == password[i-1] {
				return false
			}
		}
	}

	return ret == 0x1111
}

func isSmall(c uint8) bool {
	return c >= 'a' && c <= 'z'
}

func isBig(c uint8) bool {
	return c >= 'A' && c <= 'Z'
}

func isNumber(c uint8) bool {
	return c >= '0' && c <= '9'
}

func isSpecial(c uint8) bool {
	ret := []uint8{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '+'}
	for _, u := range ret {
		if c == u {
			return true
		}
	}

	return false
}
