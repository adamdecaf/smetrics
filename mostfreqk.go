package smetrics

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const DefaultK = 2
const MostFreqDefaultSimilarityLimit = 10

func MostFreqKSDF(s1, s2 string, max_dist int) int {
	x1 := most_freq_k_hash(s1)
	x2 := most_freq_k_hash(s2)

	return most_freq_k_sim(x1, x2, max_dist)
}

func most_freq_k_hash(s string) string {
	counts := make(map[string]int)
	for i := range s {
		c := counts[string(s[i])]
		counts[string(s[i])] = c + 1
	}

	sorted := ""

	max_char := ""
	max_count := -1

	for ii := 0; ii < DefaultK; ii++ {
		for k,v := range counts {
			// we have a totally new max so reset
			if v > max_count {
				max_char = k
				max_count = v
			}

			// we've matched, so go maps aren't ordered
			// need to figure out which is first in the input string:
			//  current max (max_char, max_count) or new max (k,v)
			if v == max_count {
				// When k comes first swap max_* for it
				if strings.Index(s, k) < strings.Index(s, max_char) {
					max_char = k
					max_count = v
				}
			}
		}

		delete(counts, max_char)

		// if we haven't set a count we can't add to the metric counts
		// this usually happens if strings are not long enough, e.g. 'a'
		if max_count != -1 {
			sorted = sorted + fmt.Sprintf("%s%d", max_char, max_count)
		}

		max_char = ""
		max_count = -1
	}

	return sorted
}

func most_freq_k_sim(x1, x2 string, limit int) int {
	similarity := 0

	// copied from http://arxiv.org/pdf/1401.6596.pdf

	if len(x1) >= 2 && len(x2) >= 2 {
		if x1[0] == x2[0] {
			similarity = similarity + max_s2i(x1[1], x2[1])
		}
	}

	if len(x1) >= 4 && len(x2) >= 4 {
		if x1[0] == x2[2] {
			similarity = similarity + max_s2i(x1[1], x2[3])
		}

		if x1[2] == x2[0] {
			similarity = similarity + max_s2i(x1[3], x2[1])
		}

		if x1[2] == x2[2] {
			similarity = similarity + max_s2i(x1[3], x2[3])
		}
	}

	return limit - similarity
}

// The paper has us do x1[k] + x2[l], but that doesn't work with their results..
func max_s2i(b1, b2 byte) int {
	a, _ := strconv.ParseFloat(string(b1), 64)
	b, _ := strconv.ParseFloat(string(b2), 64)
	return int(math.Max(a, b))
}
