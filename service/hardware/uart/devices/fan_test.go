package devices

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Fan_ProcessStatusData(t *testing.T) {
	dataStr := `
	{
		"fan": { "p": 50, "rpm": 1000 }
	}
`
	dataStorage := make(map[string]interface{})
	err := json.Unmarshal([]byte(dataStr), &dataStorage)
	require.NoError(t, err)

	fanData, fnd := dataStorage["fan"]
	require.True(t, fnd)

	statusStorage, castOk := fanData.(map[string]interface{})
	require.True(t, castOk)

	var fan Fan
	fan.ProcessStatusData(statusStorage)

	assert.Equal(t, uint8(50), fan.speedPercent)
	assert.Equal(t, uint32(1000), fan.rpm)
}
