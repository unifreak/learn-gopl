// Write a version of rotate that operates in a single pass.
package rotate

// Rotate rotate a slice left or right by pos.
// if pos is positive, rotate right; if pos is negative, rotate left.
func rotate(s []int, pos int) []int {
    pos = pos % len(s)
    if pos > 0 {
        return append(s[len(s)-pos:], s[:len(s)-pos]...)
    } else {
        pos = -pos
        return append(s[pos:], s[:pos]...)
    }
}