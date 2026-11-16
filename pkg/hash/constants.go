package hash

import (
	"encoding/binary"

	"dinodb/pkg/pager"
)

/////////////////////////////////////////////////////////////////////////////
////////////////////////// Low-level Constants //////////////////////////////
/////////////////////////////////////////////////////////////////////////////

const (
	ROOT_PN            int64 = 0
	PAGESIZE           int64 = pager.Pagesize
	DEPTH_OFFSET       int64 = 0
	DEPTH_SIZE         int64 = binary.MaxVarintLen64
	NUM_KEYS_OFFSET    int64 = DEPTH_OFFSET + DEPTH_SIZE
	NUM_KEYS_SIZE      int64 = binary.MaxVarintLen64
	BUCKET_HEADER_SIZE int64 = DEPTH_SIZE + NUM_KEYS_SIZE
	ENTRYSIZE          int64 = binary.MaxVarintLen64 * 2                   // int64 key, int64 value
	MAX_BUCKET_SIZE    int64 = (PAGESIZE - BUCKET_HEADER_SIZE) / ENTRYSIZE // max number of entries that can live in a bucket
)
