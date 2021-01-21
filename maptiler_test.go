package sm_test

import (
	"github.com/golang/geo/s2"
	"github.com/journeyai/go-staticmaps"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestMapTiler(t *testing.T) {

	latitude := 35.6693042057878
	longitude := 139.7663559789293

	//provider := sm.NewTileProviderMapTiler()
	//provider := sm.NewTileProviderWikimedia()
	provider := sm.NewTileProviderJourney()

	ctx := sm.NewContext()
	ctx.SetTileProvider(provider)
	ctx.SetSize(960, 540)
	ctx.SetCountry("JP")

	orange := color.RGBA{0xff, 0x5a, 0x1f, 0xff}
	marker := sm.NewMarker(s2.LatLngFromDegrees(latitude, longitude), orange, 20.0)
	ctx.AddMarker(marker)

	img, err := ctx.Render()
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Create("example.png")
	if err != nil {
		t.Fatal(err)
	}
	png.Encode(f, img)
}

// don't use japan higher zoom level tiles
func TestWorldZoomLevel(t *testing.T) {

	latitude := 35.6693042057878
	longitude := 139.7663559789293

	//provider := sm.NewTileProviderMapTiler()
	//provider := sm.NewTileProviderWikimedia()
	provider := sm.NewTileProviderJourney()

	ctx := sm.NewContext()
	ctx.SetTileProvider(provider)
	ctx.SetSize(960, 540)
	ctx.SetCountry("")

	orange := color.RGBA{0xff, 0x5a, 0x1f, 0xff}
	marker := sm.NewMarker(s2.LatLngFromDegrees(latitude, longitude), orange, 20.0)
	ctx.AddMarker(marker)

	img, err := ctx.Render()
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Create("less-zoom-example.png")
	if err != nil {
		t.Fatal(err)
	}
	png.Encode(f, img)
}
