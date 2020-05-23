package gst

/*
#include <stdlib.h>
#include <gst/gst.h>

void caps_filter_set_caps(GstElement *elem, GstCaps *caps) {
	// gchar *caps_str = g_strdup_printf ("video/x-raw,width=%d,height=%d", 1280, 720);
	// GstCaps *caps2 = gst_caps_from_string (caps_str);

	gchar *capsstr;
	capsstr = gst_caps_to_string (caps);
	g_print ("caps: %s\n", capsstr);
	g_free (capsstr);

	g_object_set(G_OBJECT(elem), "caps", caps, NULL);
}

void caps_filter_set_resolution(GstElement *elem, int width, int height) {
	gchar *caps_str = g_strdup_printf ("video/x-raw,width=(int)%d,height=(int)%d,format=(string)%s", width, height, "YUY2");
	GstCaps *caps = gst_caps_from_string (caps_str);

	gchar *capsstr;
	capsstr = gst_caps_to_string (caps);
	g_print ("caps: %s\n", capsstr);
	g_free (capsstr);

	g_object_set(G_OBJECT(elem), "caps", caps, NULL);

	g_free (caps_str);
  	gst_caps_unref (caps);
}

#cgo LDFLAGS: -lgstbase-1.0
*/
import "C"
import "unsafe"

type CapsFilter struct {
	BaseTransform
}

// func (b *CapsFilter) g() *C.GstBaseTransform {
// 	return (*C.GstBaseTransform)(b.GetPtr())
// }

// func (b *BaseTransform) AsBaseTransform() *BaseTransform {
// 	return b
// }

func (b *CapsFilter) SetCaps(caps *Caps) {
	C.caps_filter_set_caps(b.AsElement().g(), caps.g())
}

func (b *CapsFilter) SetResolution(width, height int) {
	C.caps_filter_set_resolution(b.AsElement().g(), C.int(width), C.int(height))
}

func NewCapsFilter(name string) *CapsFilter {
	return (*CapsFilter)(unsafe.Pointer(ElementFactoryMake("capsfilter", name)))
}
