// rewrite reverse (see rev/) function to use an array pointer instead of a slice.
package rev

const N = 10

func reverse(a *[N]int) {
    for i, j := 0, N-1; i < j; i, j = i+1, j-1 {
        a[i], a[j] = a[j], a[i]
    }
}