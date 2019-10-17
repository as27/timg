/*
Package timg allows you to generate an image. Image from a
given string.
There is no encoding included, so you can define the encoding
for the generated image by yourself.
You can use a ttf font, which should be loaded into the options
as []byte. If you not provide own fonts the gomono font
golang.org/x/image/font/gofont/gomono will be used.
*/
package timg

import (
	"reflect"
	"testing"
)

func TestWrap(t *testing.T) {
	type args struct {
		s             string
		maxCharacters int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"simple case",
			args{"this is a short text with some content", 10},
			[]string{"this is a", "short text", "with some", "content"},
		},
		{
			"long word",
			args{"abcd e abcdefghij kl mn op, qrs", 6},
			[]string{"abcd e", "abcdefghij", "kl mn", "op, qrs"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wrap(tt.args.s, tt.args.maxCharacters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wrap() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
