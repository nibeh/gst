package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

type GstSampleStruct C.GstSample

type Sample struct {
	GstSample *GstSampleStruct
}

// GetBuffer implements GstBuffer * gst_sample_get_buffer ()
func (s *Sample) GetBuffer() *Buffer {
	buffer := new(Buffer)
	buffer.GstBuffer = (*GstBufferStruct)(C.gst_sample_get_buffer((*C.GstSample)(s.GstSample)))
	return buffer
}

// GstBufferList * gst_sample_get_buffer_list ()

// GetCaps implements GstCaps * gst_sample_get_caps ()
func (s *Sample) GetCaps() *Caps {
	return (*Caps)(C.gst_sample_get_caps((*C.GstSample)(s.GstSample)))
}

// const GstStructure * gst_sample_get_info ()

// GstSegment * gst_sample_get_segment ()

// void gst_sample_set_buffer ()

// void gst_sample_set_buffer_list ()

// void gst_sample_set_caps ()

// void gst_sample_set_segment ()

// gboolean gst_sample_set_info ()

// NewSample implements GstSample * gst_sample_new ()
// func NewSample() *Sample {
// 	sample := new(Sample)
// 	sample.GstSample = (*GstSampleStruct)(C.gst_sample_new())
// 	return sample
// }

// GstSample * gst_sample_ref ()

// Unref implements void gst_sample_unref ()
func (s *Sample) Unref() {
	C.gst_sample_unref((*C.GstSample)(s.GstSample))
}

// GstSample * gst_sample_copy ()

// #define gst_sample_is_writable()

// #define gst_sample_make_writable()
