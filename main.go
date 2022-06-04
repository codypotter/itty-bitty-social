package main

import (
	"github.com/codypotter/itty-bitty-social/applayer"
	"github.com/codypotter/itty-bitty-social/httplayer"
	"github.com/codypotter/itty-bitty-social/storelayer"
)

func main() {
	// create store layer
	storeLayer := storelayer.New()

	// create app layer
	appLayer := applayer.New(storeLayer)

	// create http layer
	api := httplayer.New(appLayer)

	api.Engage()
}
