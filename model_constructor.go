package badgerwrap

import (
	badger "github.com/dgraph-io/badger/v2"
)

// NewBStore returns a type as per defined store interface. This way only the contract is exposed.
func NewBStore(pDBFilePath string, pSyncRights bool, pTheExtLogger CustomLogger) (Store, error) {
	pTheExtLogger.Info("-----------")

	var options badger.Options

	if len(pDBFilePath) == 0 {
		options = badger.DefaultOptions("").WithInMemory(true)
	} else {
		options = badger.DefaultOptions(pDBFilePath)
		options.WithSyncWrites(pSyncRights)
	}
	result, errOpen := badger.Open(options)

	return bstore{
		theLogger: pTheExtLogger,
		b:         result,
	}, errOpen
}

// Close closes the store.
func (s bstore) Close() error {
	return s.b.Close()
}