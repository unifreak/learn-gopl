package methods

import (
	"strings"
	"testing"
	"time"
)

func Test(t *testing.T) {
	Print(time.Hour)

	Print(new(strings.Replacer))
}
