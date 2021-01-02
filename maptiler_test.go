package sm_test


import (
	"os"
	"image/color"
	"image/png"
	"testing"
	"github.com/golang/geo/s2"
	"github.com/journeyai/go-staticmaps"
)


func TestMapTiler(t *testing.T) {

	latitude := 35.6693042057878
	longitude := 139.7663559789293

	provider := sm.NewTileProviderMapTiler()
	//provider := sm.NewTileProviderWikimedia()
	cache := sm.NewTileCache(os.Getenv("MAP_TILE_CACHE"), 0777)

	ctx := sm.NewContext()
	ctx.SetTileProvider(provider)
	ctx.SetSize(960, 540)
	ctx.SetCache(cache)

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
