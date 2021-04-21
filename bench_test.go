package gendoc_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/pseudomuto/protoc-gen-doc"
	"github.com/pseudomuto/protokit/utils"
)

func BenchmarkParseCodeRequest(b *testing.B) {
	set, _ := utils.LoadDescriptorSet("fixtures", "fileset.pb")
	req := utils.CreateGenRequest(set, "Booking.proto", "Vehicle.proto")
	plugin := new(Plugin)

	for i := 0; i < b.N; i++ {
		_, err := plugin.Generate(req)
		require.Nil(b, err)
	}
}
