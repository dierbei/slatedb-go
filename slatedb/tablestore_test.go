package slatedb

import (
	"testing"
)

func TestSSTWriter(t *testing.T) {
	//bucket := objstore.NewInMemBucket()
	//format := newSSTableFormat(32, 1, CompressionNone)
	//tableStore := newTableStore(bucket, format, "")
	//sstID := newSSTableIDCompacted(ulid.Make())
	//
	//writer := tableStore.tableWriter(sstID)
	//writer.add([]byte("aaaaaaaaaaaaaaaa"), mo.Some([]byte("1111111111111111")))
	//writer.add([]byte("bbbbbbbbbbbbbbbb"), mo.Some([]byte("2222222222222222")))
	//writer.add([]byte("cccccccccccccccc"), mo.None[[]byte]())
	//writer.add([]byte("dddddddddddddddd"), mo.Some([]byte("4444444444444444")))
	//sst, err := writer.close()
	//assert.NoError(t, err)
	//
	//iter := newSSTIterator(sst, tableStore, 1, 1)
	//assertIterNext(t, iter, []byte("aaaaaaaaaaaaaaaa"), []byte("1111111111111111"))
	//assertIterNextEntry(t, iter, []byte("bbbbbbbbbbbbbbbb"), []byte("2222222222222222"))
	//assertIterNextEntry(t, iter, []byte("cccccccccccccccc"), nil)
	//assertIterNextEntry(t, iter, []byte("dddddddddddddddd"), []byte("4444444444444444"))
}
