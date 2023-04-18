package agin

import "github.com/google/wire"

var Provider = wire.NewSet(NewServer)
