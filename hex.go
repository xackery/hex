// Package hex implements hexadecimal encoding and decoding.
package hex

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

var (
	pat = regexp.MustCompile("[0-9a-z]{8}  ([0-9a-z ]+) ")
)

// Dump takes a byte slice and transforms it
func Dump(data []byte) string {
	in := hex.Dump(data)
	out := ""
	matches := pat.FindAllStringSubmatch(in, -1)
	for _, subs := range matches {
		for i, sub := range subs {
			if i == 0 {
				continue
			}

			out += fmt.Sprintf("0x%s,\n", strings.ReplaceAll(strings.TrimSpace(strings.ReplaceAll(sub, "  ", " ")), " ", ", 0x"))
		}
	}

	return out
}
