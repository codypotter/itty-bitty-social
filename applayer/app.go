package applayer

import store "github.com/codypotter/itty-bitty-social/storelayer"

type App interface{}

type app struct {
	storeLayer store.Layer
}

func New(sl store.Layer) *app {
	return &app{
		storeLayer: sl,
	}
}
