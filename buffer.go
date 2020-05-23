package gst

/*
#include <stdlib.h>
#include <gst/gstbuffer.h>

void _golang_gst_set_dts( GstBuffer *buffer, guint64 value ) {
  GST_BUFFER_DTS(buffer) = value;
}

void _golang_gst_set_pts( GstBuffer *buffer, guint64 value ) {
  GST_BUFFER_PTS(buffer) = value;
}

guint64 _golang_gst_get_duration(GstBuffer *buffer ) {
  return GST_BUFFER_DURATION(buffer);
}

void _golang_gst_set_duration( GstBuffer *buffer, guint64 value ) {
  GST_BUFFER_DURATION(buffer) = value;
}

void _golang_gst_set_offset( GstBuffer *buffer, guint64 value ) {
  GST_BUFFER_OFFSET(buffer) = value;
}

void _golang_gst_set_offset_end( GstBuffer *buffer, guint64 value ) {
  GST_BUFFER_OFFSET_END(buffer) = value;
}

*/
import "C"

import (
	"unsafe"
)

type GstMapInfoStruct C.GstMapInfo

type MapInfo C.GstMapInfo

func (m *MapInfo) GetData() []byte {
	return C.GoBytes(unsafe.Pointer(m.data), C.int(m.size))
}

type MapFlags C.GstMapFlags

const (
	MAP_READ      = MapFlags(C.GST_MAP_READ)
	MAP_WRITE     = MapFlags(C.GST_MAP_WRITE)
	MAP_FLAG_LAST = MapFlags(C.GST_MAP_FLAG_LAST)
)

func (f *MapFlags) g() *C.GstMapFlags {
	return (*C.GstMapFlags)(f)
}

func (f MapFlags) String() string {
	switch f {
	case MAP_READ:
		return "MAP_READ"
	case MAP_WRITE:
		return "MAP_WRITE"
	case MAP_FLAG_LAST:
		return "MAP_FLAG_LAST"
	}
	panic("Unknown map flag")
}

type GstBufferStruct C.GstBuffer

type Buffer struct {
	GstBuffer *GstBufferStruct
}

// SetPTS implements #define	GST_BUFFER_PTS()
func (b *Buffer) SetPTS(value uint64) {
	C._golang_gst_set_pts((*C.GstBuffer)(b.GstBuffer), C.guint64(value))
}

// SetDTS implements #define	GST_BUFFER_DTS()
func (b *Buffer) SetDTS(value uint64) {
	C._golang_gst_set_dts((*C.GstBuffer)(b.GstBuffer), C.guint64(value))
}

// #define	GST_BUFFER_DTS_OR_PTS()

func (b *Buffer) GetDuration() uint64 {
	return (uint64)(C._golang_gst_get_duration((*C.GstBuffer)(b.GstBuffer)))
}

// SetDuration implements #define	GST_BUFFER_DURATION()
func (b *Buffer) SetDuration(value uint64) {
	C._golang_gst_set_duration((*C.GstBuffer)(b.GstBuffer), C.guint64(value))
}

// SetOffset implements #define	GST_BUFFER_OFFSET()
func (b *Buffer) SetOffset(value uint64) {
	C._golang_gst_set_offset((*C.GstBuffer)(b.GstBuffer), C.guint64(value))
}

// SetOffsetEnd implements #define	GST_BUFFER_OFFSET_END()
func (b *Buffer) SetOffsetEnd(value uint64) {
	C._golang_gst_set_offset_end((*C.GstBuffer)(b.GstBuffer), C.guint64(value))
}

// #define	GST_BUFFER_DURATION_IS_VALID()

// #define	GST_BUFFER_PTS_IS_VALID()

// #define	GST_BUFFER_DTS_IS_VALID()

// #define	GST_BUFFER_OFFSET_IS_VALID()

// #define	GST_BUFFER_OFFSET_END_IS_VALID()

// #define	GST_BUFFER_IS_DISCONT()

// NewBuffer implements GstBuffer *	gst_buffer_new ()
func NewBuffer() *Buffer {
	buffer := new(Buffer)
	buffer.GstBuffer = (*GstBufferStruct)(C.gst_buffer_new())
	return buffer
}

// NewBufferAllocate implements GstBuffer *	gst_buffer_new_allocate ()
func NewBufferAllocate(size uint) *Buffer {
	buffer := new(Buffer)
	buffer.GstBuffer = (*GstBufferStruct)(C.gst_buffer_new_allocate(nil, C.gsize(size), nil))
	return buffer
}

// GstBuffer *	gst_buffer_new_wrapped ()

// GstBuffer *	gst_buffer_new_wrapped_bytes ()

// GstBuffer *	gst_buffer_new_wrapped_full ()

// GstBuffer *	gst_buffer_ref ()

// Unref implements void	gst_buffer_unref ()
func (b *Buffer) Unref() {
	C.gst_buffer_unref((*C.GstBuffer)(b.GstBuffer))
}

// gsize	gst_buffer_get_sizes ()

// GetSize implements gsize	gst_buffer_get_size ()
func (b *Buffer) GetSize() uint {
	return (uint)(C.gst_buffer_get_size((*C.GstBuffer)(b.GstBuffer)))
}

// gsize	gst_buffer_get_sizes_range ()

// gboolean	gst_buffer_resize_range ()

// void	gst_buffer_resize ()

// void	gst_buffer_set_size ()

// guint	gst_buffer_get_max_memory ()

// GstMemory *	gst_buffer_peek_memory ()

// guint	gst_buffer_n_memory ()

// void	gst_buffer_insert_memory ()

// void	gst_buffer_replace_memory_range ()

// GstMemory *	gst_buffer_get_memory_range ()

// void	gst_buffer_remove_memory_range ()

// gboolean	gst_buffer_find_memory ()

// void	gst_buffer_prepend_memory ()

// AppendMemory implements void	gst_buffer_append_memory ()
func (b *Buffer) AppendMemory(memory *Memory) {
	C.gst_buffer_append_memory((*C.GstBuffer)(b.GstBuffer), (*C.GstMemory)(memory))
}

// void	gst_buffer_replace_memory ()

// void	gst_buffer_replace_all_memory ()

// GstMemory *	gst_buffer_get_memory ()

// GstMemory *	gst_buffer_get_all_memory ()

// void	gst_buffer_remove_memory ()

// void	gst_buffer_remove_all_memory ()

// gboolean	gst_buffer_is_all_memory_writable ()

// gboolean	gst_buffer_is_memory_range_writable ()

// Map implements gboolean	gst_buffer_map ()
func (b *Buffer) Map(flags MapFlags) (bool, *MapInfo) {
	mi := new(MapInfo)
	success := Itob((int)(C.gst_buffer_map((*C.GstBuffer)(b.GstBuffer), (*C.GstMapInfo)(mi), (C.GstMapFlags)(flags))))
	return success, mi
}

// gboolean	gst_buffer_map_range ()

// void	gst_buffer_unmap ()
func (b *Buffer) Unmap(mi *MapInfo) {
	C.gst_buffer_unmap((*C.GstBuffer)(b.GstBuffer), (*C.GstMapInfo)(mi))
}

// gint	gst_buffer_memcmp ()

// gsize	gst_buffer_extract ()

// void	gst_buffer_extract_dup ()

// Fill implements gsize	gst_buffer_fill ()
func (b *Buffer) Fill(offset uint, src unsafe.Pointer, size uint) int {
	return (int)(C.gst_buffer_fill((*C.GstBuffer)(b.GstBuffer), C.gsize(offset), C.gconstpointer(src), C.gsize(size)))
}

// MemSet implements gsize	gst_buffer_memset ()
func (b *Buffer) MemSet(offset uint, val byte, size uint) int {
	return (int)(C.gst_buffer_memset((*C.GstBuffer)(b.GstBuffer), C.gsize(offset), C.guint8(val), C.gsize(size)))
}

// GstBuffer *	gst_buffer_copy ()

// gboolean	gst_buffer_copy_into ()

// GstBuffer *	gst_buffer_copy_region ()

// GstBuffer *	gst_buffer_copy_deep ()

// #define	gst_buffer_is_writable()

// #define	gst_buffer_make_writable()

// gboolean	gst_buffer_replace ()

// GstBuffer *	gst_buffer_append ()

// GstBuffer *	gst_buffer_append_region ()

// GstMeta *	gst_buffer_get_meta ()

// guint	gst_buffer_get_n_meta ()

// GstMeta *	gst_buffer_add_meta ()

// gboolean	gst_buffer_remove_meta ()

// GstMeta *	gst_buffer_iterate_meta ()

// GstMeta *	gst_buffer_iterate_meta_filtered ()

// gboolean	(*GstBufferForeachMetaFunc) ()

// gboolean	gst_buffer_foreach_meta ()

// GstParentBufferMeta *	gst_buffer_add_parent_buffer_meta ()

// #define	gst_buffer_get_parent_buffer_meta()

// GstReferenceTimestampMeta *	gst_buffer_add_reference_timestamp_meta ()

// GstReferenceTimestampMeta *	gst_buffer_get_reference_timestamp_meta ()

// GstBufferFlags	gst_buffer_get_flags ()

// gboolean	gst_buffer_set_flags ()

// gboolean	gst_buffer_unset_flags ()

// gboolean	gst_buffer_has_flags ()
