package utils

import "strings"

func NormalizeList(val any) []string {
	switch v := val.(type) {
	case string:
		return strings.Split(v, ",")
	case []string:
		return v
	case []any:
		out := make([]string, 0, len(v))
		for _, x := range v {
			if s, ok := x.(string); ok {
				out = append(out, s)
			}
		}
		return out
	default:
		return nil
	}
}
