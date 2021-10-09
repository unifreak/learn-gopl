package word

import "testing"

func TestPalindrome(t *testing.T) {
    if !isPalindrome("detartrated") {
        t.Error(`IsPalindrome("detartrated") = false`)
    }
    if !isPalindrome("kayak") {
        t.Error(`IsPalindrom("kayak") = false`)
    }
}

func TestNonPanlindrom(t *testing.T) {
    if isPalindrome("panlindrom") {
        t.Error(`IsPanlindrom("palindrom") = true`)
    }
}

func TestFrenchPanlindrome(t *testing.T) {
    if !isPalindrome("été") {
        t.Error(`IsPalindrome("été") = false`)
    }
}

func TestCanalPalindrome(t *testing.T) {
    input := "A man, a plan, a canal, Panama"
    if !isPalindrome(input) {
        t.Errorf(`isPanlindrome(%q) = false`, input)
    }
}