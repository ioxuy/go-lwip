//go:build linux || darwin
// +build linux darwin

package golwip

/*
#cgo CFLAGS: -I./c/include
#include "lwip/init.h"
*/
import "C"

func lwipInit() {
	C.lwip_init() // Initialze modules.
}
