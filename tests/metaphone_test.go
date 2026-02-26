package tests

import (
	"fmt"
	"github.com/xrash/smetrics"
	"testing"
)

func TestMetaphone(t *testing.T) {
	cases := []metaphonecase{
		// Empty and single character.
		{"", ""},
		{"A", "A"},
		{"a", "A"},
		{"B", "B"},

		// Non-letter input.
		{"123", ""},
		{"a1b", "AB"},

		// Common words.
		{"Donald", "TNLT"},
		{"Zach", "SX"},
		{"Campbel", "KMPBL"},
		{"David", "TFT"},
		{"Metaphone", "MTFN"},
		{"garbage", "KRBJ"},
		{"tech", "TX"},
		{"orange", "ORNJ"},
		{"science", "SNS"},
		{"dumb", "TM"},
		{"dodge", "TJ"},
		{"jack", "JK"},
		{"queen", "KN"},
		{"widow", "WT"},
		{"xerox", "SRKS"},
		{"vowel", "FWL"},
		{"zen", "SN"},
		{"tiamat", "XMT"},
		{"ratio", "RX"},
		{"the", "0"},
		{"acacia", "AKX"},
		{"bosch", "BSK"},
		{"chop", "XP"},
		{"hatch", "HX"},
		{"Incomprehensibility", "INKMPRHNSBLT"},

		// Initial character transforms.
		{"knight", "NT"},  // KN → N
		{"gnome", "NM"},   // GN → N
		{"pneumatic", "NMTK"}, // PN → N
		{"aesthetic", "ES0TK"}, // AE → E
		{"wrong", "RNK"},  // WR → R

		// Silent letters and digraphs.
		{"lamb", "LM"},    // silent B after M at end
		{"phone", "FN"},   // PH → F
		{"ship", "XP"},    // SH → X
		{"thick", "0K"},   // TH → 0 (theta)
		{"match", "MX"},   // TCH → silent T, CH → X
		{"edge", "EJ"},    // DGE → J

		// CH always maps to X per original algorithm.
		{"church", "XRX"},
		{"chess", "XS"},

		// Full output without truncation.
		{"crying", "KRYNK"},

		// GH silent at end and before consonant.
		{"weigh", "W"},
		{"ghost", "KST"},

		// WH → W at start.
		{"wh", ""},
		{"whistle", "WSTL"},

		// GN silences G at any non-initial position (per original algorithm).
		{"signal", "SNL"},
	}

	for _, c := range cases {
		if r := smetrics.Metaphone(c.s); r != c.t {
			fmt.Println(r, "instead of", c.t, "for", c.s)
			t.Fail()
		}
	}
}
