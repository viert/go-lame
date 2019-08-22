package lame

/*
#cgo LDFLAGS: -lmp3lame
#include <stdlib.h>
#include <lame/lame.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

const (
	tagBufferSizeInitial = 32768
)

// InitID3Tag initializes id3 metadata
func (e *Encoder) InitID3Tag() {
	C.id3tag_init(e.lgf)
}

// ID3TagAddV2 forces addition of version 2 tag
func (e *Encoder) ID3TagAddV2() {
	C.id3tag_add_v2(e.lgf)
}

// ID3TagV1Only sets addition of only a version 1 tag
func (e *Encoder) ID3TagV1Only() {
	C.id3tag_v1_only(e.lgf)
}

// ID3TagV2Only sets addition of only a version 2 tag
func (e *Encoder) ID3TagV2Only() {
	C.id3tag_v2_only(e.lgf)
}

// ID3TagSpaceV1 sets version 1 tag padding with spaces instead of nulls
func (e *Encoder) ID3TagSpaceV1() {
	C.id3tag_space_v1(e.lgf)
}

// ID3TagPadV2 pads version 2 tag with extra 128 bytes
func (e *Encoder) ID3TagPadV2() {
	C.id3tag_pad_v2(e.lgf)
}

// ID3TagSetPad pads version 2 tag with extra n bytes
func (e *Encoder) ID3TagSetPad(n int) {
	C.id3tag_set_pad(e.lgf, C.size_t(n))
}

// ID3TagSetTitle sets id3 title
func (e *Encoder) ID3TagSetTitle(value string) {
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))
	C.id3tag_set_title(e.lgf, cstr)
}

// ID3TagSetArtist sets id3 artist
func (e *Encoder) ID3TagSetArtist(value string) {
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))
	C.id3tag_set_artist(e.lgf, cstr)
}

// ID3TagSetAlbum sets id3 album
func (e *Encoder) ID3TagSetAlbum(value string) {
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))
	C.id3tag_set_album(e.lgf, cstr)
}

// ID3TagSetYear sets id3 year
func (e *Encoder) ID3TagSetYear(value string) {
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))
	C.id3tag_set_year(e.lgf, cstr)
}

// ID3TagSetComment sets id3 comment
func (e *Encoder) ID3TagSetComment(value string) {
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))
	C.id3tag_set_comment(e.lgf, cstr)
}

// ID3TagSetTrack sets id3 track
//  note that lame accepts a string and returns an error if
//  track value is out of range
func (e *Encoder) ID3TagSetTrack(value string) error {
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))
	errcode := C.id3tag_set_track(e.lgf, cstr)
	if errcode == -1 {
		return fmt.Errorf("id3 track value out of range")
	}
	return nil
}

// ID3TagSetGenre sets id3 genre
func (e *Encoder) ID3TagSetGenre(value string) error {
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))
	errcode := C.id3tag_set_genre(e.lgf, cstr)
	switch errcode {
	case -1:
		return fmt.Errorf("id3 genre number out of range")
	case -2:
		return fmt.Errorf("id3 v1 tag set to 'other'")
	default:
		return nil
	}
}

// ID3V1Tag returns version 1 id3 tag
func (e *Encoder) ID3V1Tag() []byte {
	buffer := make([]byte, tagBufferSizeInitial)
	bptr := (*C.uchar)(&buffer[0])
	reqsize := int(C.lame_get_id3v1_tag(e.lgf, bptr, tagBufferSizeInitial))
	if reqsize < tagBufferSizeInitial {
		return buffer[:reqsize]
	}
	buffer = make([]byte, reqsize)
	bptr = (*C.uchar)(&buffer[0])
	C.lame_get_id3v1_tag(e.lgf, bptr, C.size_t(reqsize))
	return buffer
}

// ID3V2Tag returns version 2 id3 tag
func (e *Encoder) ID3V2Tag() []byte {
	buffer := make([]byte, tagBufferSizeInitial)
	bptr := (*C.uchar)(&buffer[0])
	reqsize := int(C.lame_get_id3v2_tag(e.lgf, bptr, tagBufferSizeInitial))
	if reqsize < tagBufferSizeInitial {
		return buffer[:reqsize]
	}
	buffer = make([]byte, reqsize)
	bptr = (*C.uchar)(&buffer[0])
	C.lame_get_id3v2_tag(e.lgf, bptr, C.size_t(reqsize))
	return buffer
}

// SetWriteID3TagAutomatic sets automatic write of id3 tag
//   Normaly lame_init_param writes ID3v2 tags into the audio stream.
//   Here in Encoder lame_init_param is launched on first write to encoder instance.
//   Call SetWriteID3TagAutomatic(false) before writing to encoder
//   to turn off this behaviour and get ID3v2 tag with above function
//   write it yourself into your file.
func (e *Encoder) SetWriteID3TagAutomatic(auto bool) {
	var value int
	if auto {
		value = 1
	}
	C.lame_set_write_id3tag_automatic(e.lgf, C.int(value))
}

// WriteID3TagAutomatic returns current automatic tag write flag
func (e *Encoder) WriteID3TagAutomatic() bool {
	res := C.lame_get_write_id3tag_automatic(e.lgf)
	return res == 1
}

// LameTagFrame returns the final LAME-tag frame
//  * NOTE:
//  * if VBR  tags are turned off by the user, or turned off by LAME,
//  * this call does nothing and returns 0.
//  * NOTE:
//  * LAME inserted an empty frame in the beginning of mp3 audio data,
//  * which you have to replace by the final LAME-tag frame after encoding.
//  * In case there is no ID3v2 tag, usually this frame will be the very first
//  * data in your mp3 file. If you put some other leading data into your
//  * file, you'll have to do some bookkeeping about where to write this buffer.
func (e *Encoder) LameTagFrame() []byte {
	buffer := make([]byte, tagBufferSizeInitial)
	bptr := (*C.uchar)(&buffer[0])
	bsize := C.lame_get_lametag_frame(e.lgf, bptr, C.size_t(tagBufferSizeInitial))
	return buffer[:bsize]
}
