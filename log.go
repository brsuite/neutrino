package neutrino

import (
	"github.com/brsuite/brond/addrmgr"
	"github.com/brsuite/brond/blockchain"
	"github.com/brsuite/brond/connmgr"
	"github.com/brsuite/brond/peer"
	"github.com/brsuite/brond/txscript"
	"github.com/brsuite/bronlog"
	"github.com/brsuite/neutrino/blockntfns"
	"github.com/brsuite/neutrino/pushtx"
	"github.com/brsuite/neutrino/query"
)

// log is a logger that is initialized with no output filters.  This
// means the package will not perform any logging by default until the caller
// requests it.
var log bronlog.Logger

// The default amount of logging is none.
func init() {
	DisableLog()
}

// DisableLog disables all library log output.  Logging output is disabled
// by default until either UseLogger or SetLogWriter are called.
func DisableLog() {
	log = bronlog.Disabled
}

// UseLogger uses a specified Logger to output package logging info.
// This should be used in preference to SetLogWriter if the caller is also
// using bronlog.
func UseLogger(logger bronlog.Logger) {
	log = logger
	blockchain.UseLogger(logger)
	txscript.UseLogger(logger)
	peer.UseLogger(logger)
	addrmgr.UseLogger(logger)
	blockntfns.UseLogger(logger)
	pushtx.UseLogger(logger)
	connmgr.UseLogger(logger)
	query.UseLogger(logger)
}
