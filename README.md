# go-lame

A new generation lamemp3 Go bindings to replace legacy github.com/viert/lame

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
