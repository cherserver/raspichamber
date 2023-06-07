package devices

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Thermometer_ProcessStatusData(t *testing.T) {
	dataStr := `
	{
		"thermometer": { "temp": 55.55, "hum": 10.32 }
	}
`
	dataStorage := make(map[string]interface{})
	err := json.Unmarshal([]byte(dataStr), &dataStorage)
	require.NoError(t, err)

	thData, fnd := dataStorage["thermometer"]
	require.True(t, fnd)

	statusStorage, castOk := thData.(map[string]interface{})
	require.True(t, castOk)

	var th Thermometer
	th.ProcessStatusData(statusStorage)

	assert.Equal(t, float32(55.55), th.temperature)
	assert.Equal(t, float32(10.32), th.humidity)
}
