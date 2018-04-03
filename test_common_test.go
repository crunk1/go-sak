package sak

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/kylelemons/godebug/diff"
)

var spewCfg = spew.ConfigState{
	DisablePointerAddresses: true,
}

func compare(x, y interface{}) string {
	return diff.Diff(spewCfg.Sdump(x), spewCfg.Sdump(y))
}
