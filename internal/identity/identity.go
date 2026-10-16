package identity

const Product = "UPHarness"

// IsOriginQuestion is a tiny stem-based detector (slim port of V1 identity).
func IsOriginQuestion(s string) bool {
	if len(s) < 4 || len(s) > 600 { return false }
	lower := ""
	for _, r := range s {
		if r >= 'A' && r <= 'Z' { r += 32 }
		if (r >= 'a' && r <= 'z') || r == ' ' { lower += string(r) }
	}
	hasTarget := containsAny(lower, []string{"upharness", "harness", " dich", "du ", "you"})
	hasOrigin := containsAny(lower, []string{"wer", "who", "program", "entwick", "creat", "autor", "von"})
	return hasTarget && hasOrigin
}

func containsAny(h string, subs []string) bool {
	for _, s := range subs {
		for i := 0; i+sLen(s) <= len(h); i++ {
			if h[i:i+sLen(s)] == s { return true }
		}
	}
	return false
}
func sLen(s string) int { return len(s) }
