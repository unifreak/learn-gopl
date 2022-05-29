// Write a function that reports whether two strings are anagrams of each other,
// that is, they contain the same letters in a different order.
package anagram

func IsAnagram(a, b string) bool {
    var counter = make(map[rune]int)
    for _, r := range a {
        counter[r]++
    }
    for _, r := range b {
        counter[r]--
    }
    for _, count := range counter {
        if count != 0 {
            return false
        }
    }
    return true
}