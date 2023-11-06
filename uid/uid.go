package uid

import (
	"log"
	"mgmt/common/configs"
	"os"
	"strconv"

	"github.com/bwmarrin/snowflake"
)

func Init() {
	systemID := configs.GetInt(configs.SECTION_SYSTEM, configs.SYSTEM_ID, -1)
	if systemID == -1 {
		log.Printf("system id is not be assigned")
		os.Exit(1)
		return
	}
	uniqueIDNode, _ = snowflake.NewNode(int64(systemID))

	var err error
	currentID, err = getCustomizeUniqueIDByID(uniqueIDNode.Generate().Int64())
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}

func GenerateUniqueID() int {
	idMutex.Lock()
	defer idMutex.Unlock()

	binaryCurrentID := get64BitBinaryUniqueID(currentID)

	binaryUnused := getCustomizeBinaryUnused(binaryCurrentID)
	binaryTimestamp := getCustomizeBinaryTimestamp(binaryCurrentID)
	binarySequenceID := getCustomizeBinarySequenceID(binaryCurrentID)
	binaryNodeID := getCustomizeBinaryNodeID(binaryCurrentID)

	timestamp, _ := strconv.ParseInt(binaryTimestamp, BINARY, UNIQUE_ID_LENGTH)
	sequenceID, _ := strconv.ParseInt(binarySequenceID, BINARY, UNIQUE_ID_LENGTH)

	sequenceID++
	if sequenceID >= SEQUENCE_ID_THRESHOLD {
		sequenceID = 0
		timestamp++
	}

	binaryTimestamp = get41BitBinaryTimestamp(timestamp)
	binarySequenceID = get12BitBinarySequenceID(sequenceID)

	currentID, _ = getCustomizeUniqueIDByBineryStrings(binaryUnused, binaryTimestamp, binarySequenceID, binaryNodeID)

	return int(currentID)
}
