package devices

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_DryerHatch_ProcessStatusData(t *testing.T) {
	dataStr := `
	{
		"hatch": { "angle": 17 }
	}
`
	dataStorage := make(map[string]interface{})
	err := json.Unmarshal([]byte(dataStr), &dataStorage)
	require.NoError(t, err)

	hatchData, fnd := dataStorage["hatch"]
	require.True(t, fnd)

	statusStorage, castOk := hatchData.(map[string]interface{})
	require.True(t, castOk)

	var hatch DryerHatch
	hatch.ProcessStatusData(statusStorage)

	assert.Equal(t, uint8(17), hatch.angle)
}
