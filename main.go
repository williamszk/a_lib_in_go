package bananas

import (
	"github.com/williamszk/a_lib_in_go/my_package_01"
)

func SayHi() {
	my_package_01.SayHi()
}

func SayBye() {
	my_package_01.SayBye()
}

func AnotherFuncForPatch2(){
	my_package_01.SayHi()
	my_package_01.SayBye()
}
