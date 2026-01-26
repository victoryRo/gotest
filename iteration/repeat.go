// Package iteration ...
package iteration

import (
	"strings"
)

func Repeat(val string, quantity int) string {
	var res strings.Builder
	for range quantity {
		res.WriteString(val)
	}

	return res.String()
}
