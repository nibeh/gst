package gst

/*
#include <stdlib.h>
#include <gst/gstobject.h>
#include <gst/gstminiobject.h>
#include <gst/gstmemory.h>
#include <gst/gstallocator.h>
*/
import "C"

type Memory C.GstMemory

func Allocate(size uint32) *Memory {
	return (*Memory)(C.gst_allocator_alloc(nil, C.gsize(size), nil))
}

func Free(ptr *Memory) {
	C.gst_allocator_free(nil, (*C.GstMemory)(ptr))
}
