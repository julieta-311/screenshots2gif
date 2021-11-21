package animate

import (
	"context"
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
	"path/filepath"

	_ "image/png" // needed to handle png files
)

// Animate makes a gif out of a series of png images available in imgDir. If loop
// is true, the animation will loop infinitely, otherwise it'll do a single loop.
func Animate(ctx context.Context, imgDir string, loop bool, delay int) error {
	files, _ := filepath.Glob(fmt.Sprintf("%s/*.png", imgDir))

	outGif := &gif.GIF{}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return err
		}

		img, _, err := image.Decode(f)
		if err != nil {
			return err
		}

		paletted := image.NewPaletted(img.Bounds(), palette.Plan9)
		draw.FloydSteinberg.Draw(paletted, img.Bounds(), img, image.Point{})

		f.Close()
		outGif.Image = append(outGif.Image, paletted)
		outGif.Delay = append(outGif.Delay, delay)
	}

	outGif.LoopCount = 1
	if loop {
		outGif.LoopCount = 0
	}

	outPath := fmt.Sprintf("%s/out.gif", imgDir)
	f, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE, 0o600)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := gif.EncodeAll(f, outGif); err != nil {
		return err
	}
	return nil
}
