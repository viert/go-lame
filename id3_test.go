package lame

import (
	"io/ioutil"
	"testing"
)

func TestID3V1(t *testing.T) {
	w := ioutil.Discard
	enc := NewEncoder(w)

	enc.ID3TagSetTitle("Super Song")
	data := enc.ID3V1Tag()[:13]
	if string(data) != "TAGSuper Song" {
		t.Error("ID3 test failed. String 'TAGSuper Song' not found")
	}
}

func TestID3V2(t *testing.T) {
	w := ioutil.Discard
	enc := NewEncoder(w)

	enc.ID3TagAddV2()
	enc.ID3TagSetTitle("Super Song")
	data := enc.ID3V2Tag()
	if string(data[:3]) != "ID3" {
		t.Error("ID3 marker not found")
	}
	if string(data[10:14]) != "TIT2" {
		t.Error("TIT2 marker not found")
	}
	if string(data[len(data)-10:len(data)]) != "Super Song" {
		t.Error("Song name not found")
	}
}
