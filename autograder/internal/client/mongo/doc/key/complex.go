package key

import "strings"

// Complex returns complex key based on subkeys combined with ".".
//
//	For example:
//	Complex("user", "passport", "first_name") // "user.passport.name"
func Complex(k1, k2 string, kk ...string) string {
	var sb strings.Builder
	sb.WriteString(k1)
	sb.WriteString(".")
	sb.WriteString(k2)
	if len(kk) > 0 {
		for _, k := range kk {
			sb.WriteString(".")
			sb.WriteString(k)
		}
	}
	return sb.String()
}
