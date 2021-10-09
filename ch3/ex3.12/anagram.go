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