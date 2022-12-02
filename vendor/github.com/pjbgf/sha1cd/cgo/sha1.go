package cgo

// #include <lib/sha1.h>
// #include <lib/sha1.c>
// #include <lib/ubc_check.h>
// #include <lib/ubc_check.c>
import "C"

import (
	"crypto"
	"hash"
	"unsafe"
)

const (
	Size      = 20
	BlockSize = 64
)

func init() {
	crypto.RegisterHash(crypto.SHA1, New)
}

func New() hash.Hash {
	d := new(digest)
	d.Reset()
	return d
}

type digest struct {
	ctx C.SHA1_CTX
}

func (d *digest) sum() ([]byte, bool) {
	b := make([]byte, Size)
	c := C.SHA1DCFinal((*C.uchar)(unsafe.Pointer(&b[0])), &d.ctx)
	if c != 0 {
		return b, true
	}

	return b, false
}

func (d *digest) Sum(in []byte) []byte {
	d0 := *d // use a copy of d to avoid race conditions.
	h, _ := d0.sum()
	return append(in, h...)
}

func (d *digest) CollisionResistantSum(in []byte) ([]byte, bool) {
	d0 := *d // use a copy of d to avoid race conditions.
	h, c := d0.sum()
	return append(in, h...), c
}

func (d *digest) Reset() {
	C.SHA1DCInit(&d.ctx)
}

func (d *digest) Size() int { return Size }

func (d *digest) BlockSize() int { return BlockSize }

func Sum(data []byte) ([]byte, bool) {
	d := New().(*digest)
	d.Write(data)

	return d.sum()
}
