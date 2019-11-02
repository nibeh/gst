package gst

/*
#include <stdlib.h>
#include <gst/base/gstbasesink.h>

#cgo LDFLAGS: -lgstbase-1.0
*/
import "C"

type BaseSink struct {
	Element
}

func (b *BaseSink) g() *C.GstBaseSink {
	return (*C.GstBaseSink)(b.GetPtr())
}

func (b *BaseSink) AsBaseSink() *BaseSink {
	return b
}

// gboolean gst_base_sink_query_latency ()

// GstClockTime gst_base_sink_get_latency ()

// GstFlowReturn gst_base_sink_do_preroll ()

// GstFlowReturn gst_base_sink_wait ()

// GstFlowReturn gst_base_sink_wait_preroll ()

// GstClockReturn gst_base_sink_wait_clock ()

// SetSync implements void gst_base_sink_set_sync ()
func (b *BaseSink) SetSync(sync bool) {
	C.gst_base_sink_set_sync(b.g(), C.int(Btoi(sync)))
}

// GetSync implements gboolean gst_base_sink_get_sync ()
func (b *BaseSink) GetSync() bool {
	return Itob((int)(C.gst_base_sink_get_sync(b.g())))
}

// void gst_base_sink_set_max_lateness ()

// gint64 gst_base_sink_get_max_lateness ()

// void gst_base_sink_set_qos_enabled ()

// gboolean gst_base_sink_is_qos_enabled ()

// void gst_base_sink_set_async_enabled ()
func (b *BaseSink) SetAsync(enabled bool) {
	C.gst_base_sink_set_async_enabled(b.g(), C.int(Btoi(enabled)))
}

// gboolean gst_base_sink_is_async_enabled ()

// void gst_base_sink_set_ts_offset ()

// GstClockTimeDiff gst_base_sink_get_ts_offset ()

// void gst_base_sink_set_render_delay ()

// GstClockTime gst_base_sink_get_render_delay ()

// GstSample * gst_base_sink_get_last_sample ()

// void gst_base_sink_set_blocksize ()

// guint gst_base_sink_get_blocksize ()

// guint64 gst_base_sink_get_throttle_time ()

// void gst_base_sink_set_throttle_time ()

// void gst_base_sink_set_max_bitrate ()

// guint64 gst_base_sink_get_max_bitrate ()

// void gst_base_sink_set_last_sample_enabled ()

// gboolean gst_base_sink_is_last_sample_enabled ()

// gboolean gst_base_sink_get_drop_out_of_segment ()

// void gst_base_sink_set_drop_out_of_segment ()

// #define GST_BASE_SINK_PAD()

// #define GST_BASE_SINK_GET_PREROLL_COND()

// #define GST_BASE_SINK_GET_PREROLL_LOCK()

// #define GST_BASE_SINK_PREROLL_BROADCAST()

// #define GST_BASE_SINK_PREROLL_LOCK()

// #define GST_BASE_SINK_PREROLL_SIGNAL()

// #define GST_BASE_SINK_PREROLL_TRYLOCK()

// #define GST_BASE_SINK_PREROLL_UNLOCK()

// #define GST_BASE_SINK_PREROLL_WAIT()

// #define GST_BASE_SINK_PREROLL_WAIT_UNTIL()
