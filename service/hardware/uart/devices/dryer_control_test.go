package devices

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cherserver/raspichamber/service/software"
)

func Test_DryerControl_ProcessStatusData(t *testing.T) {
	dataStr := `
	{
		"dryer": { "mode": "50" }
	}
`
	dataStorage := make(map[string]interface{})
	err := json.Unmarshal([]byte(dataStr), &dataStorage)
	require.NoError(t, err)

	dryerData, fnd := dataStorage["dryer"]
	require.True(t, fnd)

	statusStorage, castOk := dryerData.(map[string]interface{})
	require.True(t, castOk)

	var dryer DryerControl
	dryer.ProcessStatusData(statusStorage)

	assert.Equal(t, software.DryerStateOn50Degrees, dryer.state)
}
