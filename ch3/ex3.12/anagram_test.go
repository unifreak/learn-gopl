package anagram

import "testing"

func TestIsAnagram(t *testing.T) {
    tests := map[[2]string]bool{
        {"", ""}: true,
        {"", "a"}: false,
        {"ab", "ba"}: true,
        {"abaaa", "aabaa"}: true,
        {"abaaa", "aaba"}: false,
        {"abaa ", "aa ba"}: true,
    }
    for test, expect := range tests {
        if got := IsAnagram(test[0], test[1]); got != expect {
            t.Errorf("IsAnagrame(%s, %s) == %t", test[0], test[1], got)
        }
    }
}