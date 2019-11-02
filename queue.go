package gst

/*
#include <stdlib.h>
// #include <gst/gst.h>
*/
import "C"

type QueueLeaky int // TODO C.GstQueueLeaky

const (
	QUEUE_NO_LEAK         = 0 // TODO QueueLeaky(C.GST_QUEUE_NO_LEAK)
	QUEUE_LEAK_UPSTREAM   = 1 // TODO QueueLeaky(C.GST_QUEUE_LEAK_UPSTREAM)
	QUEUE_LEAK_DOWNSTREAM = 2 // TODO QueueLeaky(C.GST_QUEUE_LEAK_DOWNSTREAM)
)

func (q QueueLeaky) String() string {
	switch q {
	case QUEUE_NO_LEAK:
		return "QUEUE_NO_LEAK"
	case QUEUE_LEAK_UPSTREAM:
		return "QUEUE_LEAK_UPSTREAM"
	case QUEUE_LEAK_DOWNSTREAM:
		return "QUEUE_LEAK_DOWNSTREAM"
	}
	panic("Unknown queue leaky")
}
