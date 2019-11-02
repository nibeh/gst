package gst

/*
#include <stdlib.h>
#include <gst/base/gstbasesrc.h>
*/
import "C"

type BaseSrc struct {
	Element
}

func (b *BaseSrc) g() *C.GstBaseSrc {
	return (*C.GstBaseSrc)(b.GetPtr())
}

func (b *BaseSrc) AsBaseSrc() *BaseSrc {
	return b
}
