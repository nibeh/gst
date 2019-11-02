package gst

/*
#include <stdlib.h>
#include <gst/gstconfig.h>
#include <gst/gstbufferpool.h>
*/
import "C"

const GST_PADDING = C.GST_PADDING

type gint C.gint
type gpointer C.gpointer
type GstBufferPoolPrivate C.GstBufferPoolPrivate

type BufferPool struct {
	Object GstObj

	Flushing gint

	GstBufferPoolPrivate *GstBufferPoolPrivate

	_gst_reserved [GST_PADDING]gpointer
}
