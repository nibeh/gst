package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
#include <gst/app/gstappsink.h>

#cgo LDFLAGS: -lgstapp-1.0
*/
import "C"

import (
	"unsafe"
)

// documentation can be found here: https://gstreamer.freedesktop.org/data/doc/gstreamer/head/gst-plugins-base-libs/html/gst-plugins-base-libs-appsink.html

type AppSink struct {
	BaseSink
}

func (a *AppSink) g() *C.GstAppSink {
	return (*C.GstAppSink)(a.GetPtr())
}

func (a *AppSink) AsAppSink() *AppSink {
	return a
}

func NewAppSink(name string) *AppSink {
	return (*AppSink)(unsafe.Pointer(ElementFactoryMake("appsink", name)))
}

// SetCaps implements void gst_app_sink_set_caps()
func (a *AppSink) SetCaps(caps *Caps) {
	C.gst_app_sink_set_caps(a.g(), (*C.GstCaps)(caps))
}

// GetCaps implements GstCaps * gst_app_sink_get_caps ()
func (a *AppSink) GetCaps() *Caps {
	return (*Caps)(C.gst_app_sink_get_caps(a.g()))
}

// gboolean gst_app_sink_is_eos ()

// SetEmitSignals implements void gst_app_sink_set_emit_signals ()
func (a *AppSink) SetEmitSignals(emit bool) {
	C.gst_app_sink_set_emit_signals(a.g(), C.int(Btoi(emit)))
}

// gboolean gst_app_sink_get_emit_signals ()

// SetMaxBuffers implements void gst_app_sink_set_max_buffers ()
func (a *AppSink) SetMaxBuffers(buffers int) {
	C.gst_app_sink_set_max_buffers(a.g(), C.uint(buffers))
}

// GetMaxBuffers implements guint gst_app_sink_get_max_buffers ()
func (a *AppSink) GetMaxBuffers() uint {
	return (uint)(C.gst_app_sink_get_max_buffers(a.g()))
}

// SetDrop implements void gst_app_sink_set_drop ()
func (a *AppSink) SetDrop(drop bool) {
	C.gst_app_sink_set_drop(a.g(), C.int(Btoi(drop)))
}

// gboolean gst_app_sink_get_drop ()
func (a *AppSink) GetDrop() bool {
	return Itob((int)(C.gst_app_sink_get_drop(a.g())))
}

// GstSample * gst_app_sink_pull_preroll ()

// PullSample implements GstSample * gst_app_sink_pull_sample ()
func (a *AppSink) PullSample() *Sample {
	s := new(Sample)
	s.GstSample = (*GstSampleStruct)(C.gst_app_sink_pull_sample(a.g()))
	return s
}

// GstSample * gst_app_sink_try_pull_preroll ()

// GstSample * gst_app_sink_try_pull_sample ()
func (a *AppSink) TryPullSample(timeout uint64) *Sample {
	s := new(Sample)
	s.GstSample = (*GstSampleStruct)(C.gst_app_sink_try_pull_sample(a.g(), C.GstClockTime(timeout)))
	return s
}

// gboolean gst_app_sink_get_buffer_list_support ()

// void gst_app_sink_set_buffer_list_support ()

// gboolean gst_app_sink_get_wait_on_eos ()

// void gst_app_sink_set_wait_on_eos ()

// void gst_app_sink_set_callbacks ()
