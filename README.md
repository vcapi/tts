# TTS
This is a TTS(Text-to-Speech) Library.

![Test](https://github.com/vcapi/tts/workflows/Test/badge.svg?branch=master)

### Usage
Use google TTS api
```go
package main

import (
    "context"
    "os"

    "github.com/vcapi/tts"
)

func main() {
    fs, err := os.OpenFile("hello.mp3", os.O_CREATE, 0600)
    if err != nil {
        panic(err)
    }
    defer fs.Close()

    ctx := context.TODO()
    err = tts.Google(ctx, "hello", "en", fs)
    if err != nil {
        panic(err)
    }
}
```


### LICENSE
MIT License

