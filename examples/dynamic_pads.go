// This simple test application create live WebM content from test source,
// decode it and display.
package main

import (
	"github.com/ziutek/glib"
	"github.com/ziutek/gst"
	"fmt"
)

func main() {
	src := gst.ElementFactoryMake("videotestsrc", "Test source")
	src.SetProperty("do-timestamp", true)
	src.SetProperty("pattern", 18) // ball

	enc := gst.ElementFactoryMake("vp8enc", "VP8 encoder")

	mux := gst.ElementFactoryMake("webmmux", "WebM muxer")
	mux.SetProperty("streamable", true)

	demux := gst.ElementFactoryMake("matroskademux", "Matroska demuxer")

	dec := gst.ElementFactoryMake("vp8dec", "VP8 dcoder")

	sink := gst.ElementFactoryMake("autovideosink", "Video display element")

	pl := gst.NewPipeline("MyPipeline")

	pl.Add(src, enc, mux, demux, dec, sink)

	src.Link(enc, mux, demux)
	demux.ConnectNoi("pad-added", cbPadAdded, dec.GetStaticPad("sink"))
	dec.Link(sink)
	pl.SetState(gst.STATE_PLAYING)

	glib.NewMainLoop(nil).Run()
}

// Callback function for "pad-added" event
func cbPadAdded(dec_sink_pad, demux_new_pad *gst.Pad) {
	fmt.Println("New pad:", demux_new_pad.GetName())
	if demux_new_pad.CanLink(dec_sink_pad) {
		if demux_new_pad.Link(dec_sink_pad) != gst.PAD_LINK_OK {
			fmt.Println("Link error")
		}
	} else {
		fmt.Println("Can't link it with:", dec_sink_pad.GetName())
	}
}
