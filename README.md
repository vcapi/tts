# TTS
This is a TTS(Text-to-Speech) Library.


### Usage
Use google TTS api
```go
package main

import (
    "os"

    "github.com/vcapi/tts"
)

func main() {
    fs, err := os.OpenFile("hello.mp3", os.O_CREATE, 0600)
    if err != nil {
        panic(err)
    }
    defer fs.Close()
    err = tts.Google("hello", "en", fs)
    if err != nil {
        panic(err)
    }
}
```


### LICENSE
MIT License

