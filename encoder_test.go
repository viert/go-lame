package lame

import (
	"io/ioutil"
	"reflect"
	"runtime"
	"testing"
)

type counter struct {
	cnt int
}

func (c *counter) Write(p []byte) (n int, err error) {
	n = len(p)
	err = nil
	c.cnt += n
	return
}

func intSetGet(setter func(int) error, getter func() int, expected int, t *testing.T) {
	setFnName := runtime.FuncForPC(reflect.ValueOf(setter).Pointer()).Name()
	getFnName := runtime.FuncForPC(reflect.ValueOf(getter).Pointer()).Name()

	err := setter(expected)
	if err != nil {
		t.Errorf("function %s returned error %s", setFnName, err)
	}

	value := getter()
	if value != expected {
		t.Errorf("function %s returned %d instead of %d", getFnName, value, expected)
	}
}

func TestSetGetConfigurationValues(t *testing.T) {
	var err error
	// a general test which checks the connection between go code
	// and underlying lame_global_flags structure
	w := ioutil.Discard
	enc := NewEncoder(w)
	intSetGet(enc.SetNumChannels, enc.NumChannels, 1, t)
	intSetGet(enc.SetInSamplerate, enc.InSamplerate, 22050, t)
	intSetGet(enc.SetVBRMeanBitrateKbps, enc.VBRMeanBitrateKbps, 18293, t)
	intSetGet(enc.SetVBRMinBitrateKbps, enc.VBRMinBitrateKbps, 11024, t)
	intSetGet(enc.SetVBRMaxBitrateKbps, enc.VBRMaxBitrateKbps, 42100, t)
	intSetGet(enc.SetLowPassFrequency, enc.LowPassFrequency, 44200, t)
	intSetGet(enc.SetLowPassWidth, enc.LowPassWidth, 10, t)
	intSetGet(enc.SetHighPassFrequency, enc.HighPassFrequency, 44200, t)
	intSetGet(enc.SetHighPassWidth, enc.HighPassWidth, 10, t)
	intSetGet(enc.SetBrate, enc.Brate, 4, t)
	intSetGet(enc.SetQuality, enc.Quality, 8, t)

	err = enc.SetVBRHardMin(true)
	if err != nil {
		t.Error(err)
	}
	enforce := enc.VBRHardMin()
	if !enforce {
		t.Error("VBRHardMin returned false, expected true")
	}

	err = enc.SetNumSamples(197382)
	if err != nil {
		t.Error(err)
	}
	samples := enc.NumSamples()
	if samples != 197382 {
		t.Errorf("NumSamples returned %d, expected 197382", samples)
	}
}

func TestEncoder(t *testing.T) {
	c := new(counter)
	enc := NewEncoder(c)
	enc.SetNumChannels(1)
	enc.SetQuality(9)
	enc.SetBrate(1)

	input := make([]byte, 8192)
	for i := 0; i < 8192; i++ {
		input[i] = byte(i)
	}

	enc.Write(input)
	enc.Close()

	if c.cnt == 0 {
		t.Error("encoder byte counter is zero")
	}
	if c.cnt > 300 {
		t.Error("encoder byte counter is greater than 1500 which is too high for the best compression and worst quality")
	}
}
