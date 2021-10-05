package variadic

import "strings"

// JoinV
func JoinV(sep string, str ...string) string {
	return strings.Join(str, sep)
}
