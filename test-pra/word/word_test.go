package word

import (
	"math/rand"
	"testing"
	"time"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

// func BenchmarkIsPalindrome(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		IsPalindrome("A man, a plan, a canal: Panama")
// 	}
// }

func isPalindromeBase(b *testing.B, size int) {
	for i := 0; i < size; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func BenchmarkIsPalindrome10(b *testing.B) {
	isPalindromeBase(b, 10)
}

func BenchmarkIsPalindrome100(b *testing.B) {
	isPalindromeBase(b, 100)
}

func BenchmarkIsPalindrome10000(b *testing.B) {
	isPalindromeBase(b, 10000)
}

func BenchmarkIsPalindrome1000000(b *testing.B) {
	isPalindromeBase(b, 1000000)
}
