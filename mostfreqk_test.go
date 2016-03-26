package smetrics

import (
	"testing"
)

func TestMostFreqKSMF(t *testing.T) {
	// night, nacht
	check_most_freq_k_hash(t, "night", "n1i1")
	check_most_freq_k_hash(t, "nacht", "n1a1")
	check_most_freq_k(t, "night", "nacht", 9)

	// my, a
	check_most_freq_k_hash(t, "my", "m1y1")
	check_most_freq_k_hash(t, "a", "a1")
	check_most_freq_k(t, "my", "a", 10)

	// research, research
	check_most_freq_k_hash(t, "research", "r2e2")
	check_most_freq_k_hash(t, "research", "r2e2")
	check_most_freq_k(t, "research", "research", 6)

	// aaaaabbbb, ababababa
	check_most_freq_k_hash(t, "aaaaabbbb", "a5b4")
	check_most_freq_k_hash(t, "ababababa", "a5b4")
	check_most_freq_k(t, "aaaaabbbb", "ababababa", 1)

	// significant, capabilities
	check_most_freq_k_hash(t, "significant", "i3n2")
	check_most_freq_k_hash(t, "capabilities", "i3a2")
	check_most_freq_k(t, "significant", "capabilities", 7)

	// example from paper
	s1 := "LCLYTHIGRNIYYGSYLYSETWNTGIMLLLITMATAFMGYVLPWGQMSFWGATVITNLFSAIPYIGTNLV"
	s2 := "EWIWGGFSVDKATLNRFFAFHFILPFTMVALAGVHLTFLHETGSNNPLGLTSDSDKIPFHPYYTIKDFLG"
	sdf := 91 // was 83 in the paper, why??
	check_most_freq_k_hash(t, s1, "L9T8")
	check_most_freq_k_hash(t, s2, "F9L8")
	for i := 0; i < DefaultSimilarityRetries; i++ {
		if r := MostFreqKSDF(s1, s2, 100); r != sdf {
			t.Fatalf("MostFreqKSDF(%s, %s) = %d, expected %d", s1, s2, r, sdf)
		}
	}
}

// setting these to make sure we don't have mis-ordering from golang maps
const DefaultHashRetries = 100
const DefaultSimilarityRetries = 100

func check_most_freq_k_hash(t *testing.T, s, hash_res string) {
	for i := 0; i < DefaultHashRetries; i++ {
		if r := most_freq_k_hash(s); r != hash_res {
			t.Fatalf("most_freq_k_hash(%s) = %s, expected %s", s, r, hash_res)
		}
	}
}

func check_most_freq_k(t *testing.T, s1, s2 string, sdf int) {
	for i := 0; i < DefaultSimilarityRetries; i++ {
		if r := MostFreqKSDF(s1, s2, MostFreqDefaultSimilarityLimit); r != sdf {
			t.Fatalf("MostFreqKSDF(%s, %s) = %d, expected %d", s1, s2, r, sdf)
		}
	}
}
