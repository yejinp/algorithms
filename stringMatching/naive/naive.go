package naive

import (
	"strings"
)

func NaiveStringMatcher(p, t string) []int {
	ans := []int{}
	n, m := len(t), len(p)
	for s := 0; s < n-m; s++ {
		if 0 == strings.Compare(p, string(t[s:s+m])) {
			ans = append(ans, s)
		}
	}
	return ans
}
