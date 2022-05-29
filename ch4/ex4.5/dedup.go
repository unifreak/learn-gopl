// Write an in-place function to eliminate adjacent duplicates in a []string slice
package dedup

func dedup(s []string) []string {
    if len(s) == 0 {
        return s
    }

    n := 0
    for _, v := range s {
        if v == s[n] {
            continue
        }
        n++
        s[n] = v
    }
    return s[:n+1]
}