package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/core"
	"io"
)

const numOfBytesUuid = 24

type UUIDCoder struct{}

// makeIdentifier creates uuids for jobs, moved as-is from gocraft/work lib.
func (UUIDCoder) MakeIdentifier() string {
	b := make([]byte, 12)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", b)
}

// encodeTypeWithGuid adds an extra byte to a guid to store the type of a task and encodes the
// result as base64.
func (UUIDCoder) EncodeTypeWithGuid(guid string, typ core.TaskType) string {
	resultBytes := make([]byte, numOfBytesUuid+1)
	copy(resultBytes, guid)

	taskCode := core.CodeByTaskType[typ]
	resultBytes[numOfBytesUuid] = taskCode

	b64 := base64.StdEncoding.EncodeToString(resultBytes)
	return b64
}

// decodeTypeFromId retrieves original uuid and the type of a task from encoded string.
func (UUIDCoder) DecodeTypeFromId(encodedId string) (string, core.TaskType, error) {
	bts, err := base64.StdEncoding.DecodeString(encodedId)
	if err != nil {
		return "", "", err
	}
	if len(bts) < 25 {
		//m.log
		return "", "", fmt.Errorf("")
	}

	guidBytes := make([]byte, numOfBytesUuid)
	copy(guidBytes, bts)
	guid := string(guidBytes)

	taskCode := bts[numOfBytesUuid]
	taskType := core.TaskTypeByCode[taskCode]

	return guid, taskType, nil
}
