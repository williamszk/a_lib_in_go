// This is just a test to run this mock lib.
package main

import (
	"github.com/williamszk/a_lib_in_go/my_package_01"
	"testing"
)

func TestSayHi(t *testing.T) {
	my_package_01.SayHi()
}

func TestSayBye(t *testing.T) {
	my_package_01.SayBye()
}
