package gst

/*
#include <stdlib.h>
#include <gst/base/gstbasetransform.h>

#cgo LDFLAGS: -lgstbase-1.0
*/
import "C"

type BaseTransform struct {
	Element
}

func (b *BaseTransform) g() *C.GstBaseTransform {
	return (*C.GstBaseTransform)(b.GetPtr())
}

func (b *BaseTransform) AsBaseTransform() *BaseTransform {
	return b
}

func (b *BaseTransform) ReconfigureSink() {
	C.gst_base_transform_reconfigure_sink(b.g())
}

func (b *BaseTransform) ReconfigureSrc() {
	C.gst_base_transform_reconfigure_src(b.g())
}

func (b *BaseTransform) UpdateSrcCaps(caps *Caps) bool {
	return C.gst_base_transform_update_src_caps(b.g(), caps.g()) != 0
}
