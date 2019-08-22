package lame

/*
#cgo LDFLAGS: -lmp3lame
#include <lame/lame.h>
*/
import "C"

import (
	"bufio"
	"io"
	"runtime"
	"unsafe"
)

// Encoder represents a Writer interface to lame encoder
type Encoder struct {
	lgf          lameglobal
	output       *bufio.Writer
	closed       bool
	initialized  bool
	inremainder  []byte
	outremainder []byte
}

// NewEncoder creates a new encoder
func NewEncoder(w io.Writer) *Encoder {
	e := &Encoder{
		lgf:          C.lame_init(),
		output:       bufio.NewWriter(w),
		initialized:  false,
		closed:       false,
		inremainder:  nil,
		outremainder: nil,
	}
	runtime.SetFinalizer(e, finalize)
	return e
}

func finalize(e *Encoder) {
	e.Close()
}

// SetVBR sets vbr mode
func (e *Encoder) SetVBR(mode VBRMode) error {
	res := int(C.lame_set_VBR(e.lgf, C.vbr_mode(mode)))
	return convError(res)
}

// SetVBRMeanBitrateKbps sets VBR mean bitrate
//  Ignored unless VBRABR mode is used
func (e *Encoder) SetVBRMeanBitrateKbps(kbps int) error {
	res := int(C.lame_set_VBR_mean_bitrate_kbps(e.lgf, C.int(kbps)))
	return convError(res)
}

// VBRMeanBitrateKbps returns VBR mean bitrate
func (e *Encoder) VBRMeanBitrateKbps() int {
	return int(C.lame_get_VBR_mean_bitrate_kbps(e.lgf))
}

// SetVBRMinBitrateKbps sets min bitrate
// I gnored unless VBRABR mode is used
func (e *Encoder) SetVBRMinBitrateKbps(kbps int) error {
	res := int(C.lame_set_VBR_min_bitrate_kbps(e.lgf, C.int(kbps)))
	return convError(res)
}

// VBRMinBitrateKbps returns VBR mean bitrate
func (e *Encoder) VBRMinBitrateKbps() int {
	return int(C.lame_get_VBR_min_bitrate_kbps(e.lgf))
}

// SetVBRMaxBitrateKbps sets max bitrate
//  Ignored unless VBRABR mode is used
func (e *Encoder) SetVBRMaxBitrateKbps(kbps int) error {
	res := int(C.lame_set_VBR_max_bitrate_kbps(e.lgf, C.int(kbps)))
	return convError(res)
}

// VBRMaxBitrateKbps returns VBR mean bitrate
func (e *Encoder) VBRMaxBitrateKbps() int {
	return int(C.lame_get_VBR_max_bitrate_kbps(e.lgf))
}

// SetVBRHardMin when enforce==true, strictly enforces min bitrate
//  Normally it will be violated for analog silence
func (e *Encoder) SetVBRHardMin(enforce bool) error {
	var value int
	if enforce {
		value = 1
	}
	res := int(C.lame_set_VBR_hard_min(e.lgf, C.int(value)))
	return convError(res)
}

// VBRHardMin returns enforced min bitrate value
func (e *Encoder) VBRHardMin() bool {
	return int(C.lame_get_VBR_hard_min(e.lgf)) == 1
}

// SetVBRQuality sets VBR quality level.  0=highest  9=lowest, Range [0,...,10[
func (e *Encoder) SetVBRQuality(quality float64) error {
	res := int(C.lame_set_VBR_quality(e.lgf, C.float(quality)))
	return convError(res)
}

// SetLowPassFrequency applies lowpass filtering to frequency in Hz
//  0 - lame chooses
//  -1 - disable lowpass
//  default is 0
func (e *Encoder) SetLowPassFrequency(frequency int) error {
	res := int(C.lame_set_lowpassfreq(e.lgf, C.int(frequency)))
	return convError(res)
}

// LowPassFrequency returns current lowpass frequency value
func (e *Encoder) LowPassFrequency() int {
	return int(C.lame_get_lowpassfreq(e.lgf))
}

// SetLowPassWidth sets the width of transition band in Hz
//  default = one polyphase filter band
func (e *Encoder) SetLowPassWidth(frequency int) error {
	res := int(C.lame_set_lowpasswidth(e.lgf, C.int(frequency)))
	return convError(res)
}

// LowPassWidth returns current lowpass width value
func (e *Encoder) LowPassWidth() int {
	return int(C.lame_get_lowpasswidth(e.lgf))
}

// SetHighPassFrequency applies lowpass filtering to frequency in Hz
//  0 - lame chooses
//  -1 - disable lowpass
//  default is 0
func (e *Encoder) SetHighPassFrequency(frequency int) error {
	res := int(C.lame_set_highpassfreq(e.lgf, C.int(frequency)))
	return convError(res)
}

// HighPassFrequency returns current highpass frequency value
func (e *Encoder) HighPassFrequency() int {
	return int(C.lame_get_highpassfreq(e.lgf))
}

// SetHighPassWidth sets the width of transition band in Hz
//  default = one polyphase filter band
func (e *Encoder) SetHighPassWidth(frequency int) error {
	res := int(C.lame_set_highpasswidth(e.lgf, C.int(frequency)))
	return convError(res)
}

// HighPassWidth returns current highpass width value
func (e *Encoder) HighPassWidth() int {
	return int(C.lame_get_highpasswidth(e.lgf))
}

// SetNumChannels sets number of channels in input stream
//  default is 2
func (e *Encoder) SetNumChannels(num int) error {
	res := int(C.lame_set_num_channels(e.lgf, C.int(num)))
	return convError(res)
}

// NumChannels returns current input numchannels value
func (e *Encoder) NumChannels() int {
	return int(C.lame_get_num_channels(e.lgf))
}

// SetNumSamples sets number of samples.
//  default = 2^32-1
func (e *Encoder) SetNumSamples(numSamples uint32) error {
	res := int(C.lame_set_num_samples(e.lgf, C.ulong(numSamples)))
	return convError(res)
}

// NumSamples gets number of samples
func (e *Encoder) NumSamples() uint32 {
	return uint32(C.lame_get_num_samples(e.lgf))
}

// SetInSamplerate sets input sample rate in Hz
//  default is 44100
func (e *Encoder) SetInSamplerate(sampleRate int) error {
	res := int(C.lame_set_in_samplerate(e.lgf, C.int(sampleRate)))
	return convError(res)
}

// InSamplerate returns current input sample rate configured
func (e *Encoder) InSamplerate() int {
	return int(C.lame_get_in_samplerate(e.lgf))
}

// SetBrate sets one of brate compression ratio.
//  default is compression ratio of 11
func (e *Encoder) SetBrate(brate int) error {
	res := int(C.lame_set_brate(e.lgf, C.int(brate)))
	return convError(res)
}

// Brate returns current brate compression ratio
func (e *Encoder) Brate() int {
	return int(C.lame_get_brate(e.lgf))
}

// SetMode sets output audio mode
//  mode = 0,1,2,3 = stereo, jstereo, dual channel (not supported), mono
//  default: lame picks based on compression ration and input channels
func (e *Encoder) SetMode(mode MpegMode) error {
	res := int(C.lame_set_mode(e.lgf, C.MPEG_mode(mode)))
	return convError(res)
}

// SetQuality chooses internal algorithm selection.
//  True quality is determined by the bitrate
//  but this variable will effect quality by selecting expensive or cheap algorithms.
//  quality=0..9.  0=best (very slow).  9=worst.
//  recommended:  2     near-best quality, not too slow
//                5     good quality, fast
//                7     ok quality, really fast
func (e *Encoder) SetQuality(quality int) error {
	res := int(C.lame_set_quality(e.lgf, C.int(quality)))
	return convError(res)
}

// Quality returns current quality value
func (e *Encoder) Quality() int {
	return int(C.lame_get_quality(e.lgf))
}

func (e *Encoder) initParams() error {
	if e.initialized {
		return nil
	}
	err := convError(int(C.lame_init_params(e.lgf)))
	if err == nil {
		e.initialized = true
	}
	return err
}

// Write implements a default Writer interface
func (e *Encoder) Write(p []byte) (int, error) {
	var n int
	var err error

	// Write should always return the input data size if there was no error
	inputDataSize := len(p)

	if !e.initialized {
		e.initParams()
	}

	if e.inremainder != nil {
		p = append(e.inremainder, p...)
	}

	if len(p) == 0 {
		return 0, nil
	}

	blockAlignment := bitDepth / 8 * e.NumChannels() // 2 bytes per channel
	bytesRemain := len(p) % blockAlignment
	if bytesRemain > 0 {
		e.inremainder = p[len(p)-bytesRemain : len(p)]
		p = p[0 : len(p)-bytesRemain]
	} else {
		e.inremainder = nil
	}

	//
	// From lame.h:
	// The required mp3buf_size can be computed from num_samples,
	// samplerate and encoding rate, but here is a worst case estimate:
	//
	// mp3buf_size in bytes = 1.25*num_samples + 7200
	//
	numSamples := len(p) / blockAlignment
	estimatedSize := int(1.25*float64(numSamples)) + 7200
	o := make([]byte, estimatedSize)

	cp := (*C.short)(unsafe.Pointer(&p[0]))
	co := (*C.uchar)(unsafe.Pointer(&o[0]))

	if e.NumChannels() == 1 {
		n = int(C.lame_encode_buffer(
			e.lgf,
			cp,
			nil,
			C.int(numSamples),
			co,
			C.int(estimatedSize),
		))
	} else {
		n = int(C.lame_encode_buffer_interleaved(
			e.lgf,
			cp,
			C.int(numSamples),
			co,
			C.int(estimatedSize),
		))
	}

	if n < 0 {
		err = convError(n)
		return 0, err
	}

	o = o[:n]
	if e.outremainder != nil {
		o = append(e.outremainder, o...)
	}

	m, err := e.output.Write(o)
	if m < len(o) {
		e.outremainder = o[m : len(o)-m]
	} else {
		e.outremainder = nil
	}
	return inputDataSize, err
}

// Flush flushes the encoder buffer
func (e *Encoder) Flush() (n int, err error) {
	estimatedSize := 7200
	o := make([]byte, estimatedSize)
	co := (*C.uchar)(unsafe.Pointer(&o[0]))
	bytesOut := C.int(C.lame_encode_flush(
		e.lgf,
		co,
		C.int(estimatedSize),
	))
	if bytesOut != 0 {
		n, err = e.output.Write(o[:bytesOut])
	} else {
		n = 0
	}
	e.output.Flush()
	return
}

// Close closes the encoder if it's not closed yet
// Note that encoder is being closed automatically on GC
func (e *Encoder) Close() {
	if e.closed {
		return
	}
	e.Flush()
	C.lame_close(e.lgf)
	e.closed = true
}
