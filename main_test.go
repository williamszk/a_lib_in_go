// This is just a test to run this mock lib.
package bananas

import (
	"testing"

	"github.com/williamszk/a_lib_in_go/my_package_01"
)

func TestSayHi(t *testing.T) {
	my_package_01.SayHi()
}

func TestSayBye(t *testing.T) {
	my_package_01.SayBye()
}
