// +build darwin linux

package main

import (
	"github.com/pankona/gomo-simra/examples/audio/scene"
	"github.com/pankona/gomo-simra/simra"
)

func main() {
	simra.LogDebug("[IN]")
	sim := simra.NewSimra()
	sim.Start(&scene.Title{})
	simra.LogDebug("[OUT]")
}
