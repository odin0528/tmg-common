package uid

import (
	"sync"

	"github.com/bwmarrin/snowflake"
)

const (
	UNIQUE_ID_LENGTH  = 64
	TIMESTAMP_LENGTH  = 41
	SEUENCE_ID_LENGTH = 12

	BINARY = 2

	SEQUENCE_ID_THRESHOLD = 4096

	TIME_BACK_MILLISECOND_THRESHOLD = 5000

	ERROR_UNIQUE_ID = -1
)

var uniqueIDNode *snowflake.Node
var idMutex sync.RWMutex
var currentID int64
