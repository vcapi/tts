package tts

import (
	"os"
	"testing"
)

func TestGoogleTTS(t *testing.T) {
	tests := []struct {
		Text   string
		Lang   string
		Output string
	}{
		{"Hello, Do you speak a foreign language? One language is never enough.", "en", "tts_google_en.mp3"},
		{"你好，你会不会说其他外语？只会一种语言是不够的。", "zh-Hans", "tts_google_zh-Hans.mp3"},
	}

	for _, tt := range tests {
		fs, err := os.OpenFile(tt.Output, os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			t.Errorf("Open tts output file error: %v", err)
			return
		}
		err = Google(tt.Text, tt.Lang, fs)
		if err != nil {
			t.Errorf("Google tts error: %v", err)
			return
		}
	}

}
