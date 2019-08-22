# go-lame

A new generation lamemp3 Go bindings to replace legacy github.com/viert/lame

## What's new

  * more lame library code bound
  * better id3 tag support
  * used in a real project of mine so is going to be developed and maintained more rapidly and carefully (and hopefully)

## Example

```go
package main

import (
	"bufio"
	"os"

    "github.com/viert/go-lame"
)

func main() {
	of, err := os.Create("output.mp3")
	if err != nil {
		panic(err)
	}
	defer of.Close()
	enc := lame.NewEncoder(of)
	defer enc.Close()

	inf, err := os.Open("input.wav")
	if err != nil {
		panic(err)
	}
	defer inf.Close()

	r := bufio.NewReader(inf)
	r.WriteTo(enc)
}
```
