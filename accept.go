package accept

import (
	"strconv"
	"strings"
)

// Encoding checks if the given encoding val is acceptable
// according to the Accept-Encoding header.
// See https://www.rfc-editor.org/rfc/rfc7231#section-5.3.4 .
//
// If header is empty/nonexistent, the RFC says, "the user agent has
// no preferences regarding content-codings", and also "it does not imply
// that the user agent will be able to correctly process all encodings", so
// we assume it will not, and return false.
// Empty val also means that this check fails.
func Encoding(header, val string) bool {
	if header == "" || val == "" {
		return false
	}

	for _, s := range strings.Split(header, ",") {
		s, w, _ := strings.Cut(s, ";")
		s = strings.TrimSpace(s)
		w = strings.TrimSpace(w)
		if w == "" {
			if matchval(s, val) {
				return true
			}
			continue
		}
		if !strings.HasPrefix(w, "q=") {
			continue
		}
		q, err := strconv.ParseFloat(w[2:], 64)
		if err != nil {
			continue
		}
		if q > 0 && matchval(s, val) {
			return true
		}
	}
	return false
}

func matchval(s, val string) bool {
	return s == "*" || strings.EqualFold(s, val)
}
