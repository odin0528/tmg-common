package uid

import (
	"errors"
	"strconv"
)

func getCustomizeUniqueIDByID(uniqueID int64) (int64, error) {
	binaryUniqueID := get64BitBinaryUniqueID(uniqueID)

	if len(binaryUniqueID) != UNIQUE_ID_LENGTH {
		return uniqueID, errors.New("unique id length error")
	}

	unused := binaryUniqueID[0:1]
	timestamp := binaryUniqueID[1:42]
	nodeID := binaryUniqueID[42:52]
	sequenceID := binaryUniqueID[52:64]

	return getCustomizeUniqueIDByBineryStrings(unused, timestamp, sequenceID, nodeID)
}

func getCustomizeUniqueIDByBineryStrings(unused, timestamp, sequenceID, nodeID string) (int64, error) {
	temp := unused + timestamp + sequenceID + nodeID

	if len(temp) != UNIQUE_ID_LENGTH {
		return ERROR_UNIQUE_ID, errors.New("unique id length error")
	} 

	result, err := strconv.ParseInt(temp, BINARY, UNIQUE_ID_LENGTH)
	if err != nil {
		return ERROR_UNIQUE_ID, err
	}

	return result, nil
}

func get64BitBinaryUniqueID(uniqueID int64) string {
	binaryUniqueID := strconv.FormatInt(uniqueID, BINARY)
	for i := len(binaryUniqueID); i < UNIQUE_ID_LENGTH; i++ {
		binaryUniqueID = "0" + binaryUniqueID
	}

	return binaryUniqueID
}

func get41BitBinaryTimestamp(timestamp int64) string {
	binaryTimestamp := strconv.FormatInt(timestamp, BINARY)
	for i := len(binaryTimestamp); i < TIMESTAMP_LENGTH; i++ {
		binaryTimestamp = "0" + binaryTimestamp
	}

	return binaryTimestamp
}

func get12BitBinarySequenceID(sequenceID int64) string {
	binarySequence := strconv.FormatInt(sequenceID, BINARY)
	for i := len(binarySequence); i < SEUENCE_ID_LENGTH; i++ {
		binarySequence = "0" + binarySequence
	}

	return binarySequence
}

func getCustomizeBinaryUnused(customizeBinaryUniqueID string) string {
	return customizeBinaryUniqueID[0:1]
}

func getCustomizeBinaryTimestamp(customizeBinaryUniqueID string) string {
	return customizeBinaryUniqueID[1:42]
}

func getCustomizeBinarySequenceID(customizeBinaryUniqueID string) string {
	return customizeBinaryUniqueID[42:54]
}

func getCustomizeBinaryNodeID(customizeBinaryUniqueID string) string {
	return customizeBinaryUniqueID[54:64]
}
