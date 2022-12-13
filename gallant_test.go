package gallant

import (
	"image"
	"image/draw"
	"image/png"
	"os"
	"testing"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func runeRange(t *testing.T, from, count int) string {
	t.Helper()
	b := make([]rune, count)
	for i := 0; i < count; i++ {
		b[i] = rune(from + i)
	}
	return string(b)
}

func writeImage(t *testing.T, name string, img image.Image) {
	t.Helper()
	out, err := os.Create(name)
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()
	if err := png.Encode(out, img); err != nil {
		t.Fatal(err)
	}
}

func TestPreview(t *testing.T) {
	dst := image.NewRGBA(image.Rect(0, 0, 560, 300))
	draw.Draw(dst, dst.Bounds(), image.White, image.Point{}, draw.Src)
	d := &font.Drawer{
		Dst:  dst,
		Src:  image.Black,
		Face: Gallant,
		Dot:  fixed.P(20, 30),
	}
	d.DrawString("The quick brown fox jumps over the lazy dog.")
	for i := 0; i < 8; i++ {
		d.Dot = fixed.P(20, 70+i*30)
		d.DrawString(runeRange(t, i*0x20, 0x20))
	}
	writeImage(t, "preview.png", dst)
}
