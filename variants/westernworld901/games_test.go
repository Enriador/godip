package westernworld901

import (
	"testing"

	"github.com/zond/godip"

	tst "github.com/zond/godip/variants/testing"
)

func init() {
	godip.Debug = true
}

func TestGames(t *testing.T) {
	tst.TestGames(t, WesternWorld901Variant)
}
