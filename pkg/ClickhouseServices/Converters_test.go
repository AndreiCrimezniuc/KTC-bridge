package ClickhouseServices

import (
	DataStructures "Portal/Core/internal/structures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func getMockProbesWithCountry() map[string]string {
	probesWithCountries := make(map[string]string, 2)

	probesWithCountries["pp1"] = "omd"
	probesWithCountries["Gi_TOD2"] = "oro"

	return probesWithCountries

}

func TestGetCountryByProb(t *testing.T) {
	probesWithCountries := getMockProbesWithCountry()

	_, err := getCountryByProb("pp1", probesWithCountries)
	require.NoError(t, err)
	_, err = getCountryByProb("Gi_TOD2", probesWithCountries)
	require.NoError(t, err)
	_, err = getCountryByProb("pp3", probesWithCountries)
	require.Error(t, err)
}

func getValidConnectionUpData() string {
	return "99999\t99999\thello\thello\t255\thello\thello\t12312313\thello\thello\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t255\thello\thello\thello\t999\t999\t999\t"
}

func getValidConnectionDownData() string {
	return "1111111\t2\tfirst-second.bobble.yu\t99999999\t6\tcom.roworw.hello/8.0\t99.999.99.99\t9\t9\t999.999.999.999\t9\t99\t9999\t9999\t999\t9\t9\t9999\t99\t9\t9\t9\t9\t9\t9\t9\t9\t9\t9\t999\t9\t9\t9\t9\t9\t9\t9\t9\t9\t9\t9\t9\t9\t9\t9\t999\t9\t9\t9\t9\t9\t9\t9\t999\t99999:9999:9:99:99:F:E:N3:V:S:S:S\t9999999\t999b999f999b\t009999993\t0009999900\t999999\t999999\t9999\t9"
}

func getNonValidConnectionUpData() string {
	return "2\tfirst-second.bobble.yu\t99999999\t6\tcom.roworw.hello/8.0\t99.999.99.99\t9\t9\t999.999.999.999\t9\t99\t9999\t9999\t999\t9\t9\t9999\t99\t9\t9\t9\t9\t9\t9\t9\t9\t9\t9\t999\t9\t9\t9\t9\t9\t9\t9\t9\t9\t9\t9\t9\t9\t9\t9\t999\t9\t9\t9\t9\t9\t9\t9\t999\t99999:9999:9:99:99:F:E:N3:V:S:S:S\t9999999\t999b999f999b\t009999993\t0009999900\t999999\t999999\t9999\t9"
}

func getNonValidConnectionDownData() string {
	return "tttttt\t99999\thello\thello\t257\thello\thello\t12312313\thello\thello\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t999\t255\thello\thello\thello\t999\t999\t999\t"
}

func TestConvertDataToCHModel(t *testing.T) {
	probesWithCountries := getMockProbesWithCountry()
	consTypeDown := "0"
	consTypeUp := "1"

	_, errUp := ConvertDataToCHModel("newcust-"+consTypeUp, getValidConnectionUpData(), "pp1#123123", probesWithCountries, consTypeUp)
	require.NoError(t, errUp)

	_, erUp := ConvertDataToCHModel("newcust-"+consTypeUp, getNonValidConnectionUpData(), "pp1#123123", probesWithCountries, consTypeUp)
	require.Error(t, erUp)

	_, errDown := ConvertDataToCHModel("newcust-"+consTypeDown, getValidConnectionDownData(), "pp1#123123", probesWithCountries, consTypeDown)
	require.NoError(t, errDown)

	_, er := ConvertDataToCHModel("newcust-"+consTypeDown, getNonValidConnectionDownData(), "pp1#123123", probesWithCountries, consTypeDown)
	require.Error(t, er)
}

func TestAddAdditionalFields(t *testing.T) {
	var generalFields DataStructures.GeneralTrafficFields
	key := "pp1#123123123"
	probesWithCountry := getMockProbesWithCountry()
	consType := "0"

	AddAdditionalFields(&generalFields, key, probesWithCountry, consType)

	assert.Equal(t, generalFields.Country, "omd")
	assert.Equal(t, generalFields.Probe, "pp1")
	assert.Equal(t, generalFields.ProbeTimestamp, time.Unix(123123123, 0))
	assert.Equal(t, generalFields.TypeConnection, uint8(0))
}
