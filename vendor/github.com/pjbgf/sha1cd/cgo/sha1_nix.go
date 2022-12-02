//go:build !windows
// +build !windows

package cgo

// #include <lib/sha1.h>
// #include <stdlib.h>
// #include <stddef.h>
import "C"

import "unsafe"

func (d *digest) Write(p []byte) (nn int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	data := (*C.char)(unsafe.Pointer(&p[0]))
	C.SHA1DCUpdate(&d.ctx, data, (C.ulong)(len(p)))

	return len(p), nil
}
