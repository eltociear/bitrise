package integration

import (
	"testing"

	"github.com/bitrise-io/go-utils/command"
	"github.com/stretchr/testify/require"
)

const configPath = "timeout_no_output.yml"

func Test_GivenHangDetectionOn_WhenThereIsOutput_ThenDoesNotAbort(t *testing.T) {
	cmd := command.New(binPath(), "run", "output_consistent", "--config", configPath)
	out, err := cmd.RunAndReturnTrimmedCombinedOutput()

	require.NoError(t, err, "Bitrise CLI failed, output: %s", out)
}

func Test_GivenHangDetectionOn_WhenOutputSlowsDown_ThenAborts(t *testing.T) {
	cmd := command.New(binPath(), "run", "output_slows_down", "--config", configPath)
	out, err := cmd.RunAndReturnTrimmedCombinedOutput()

	require.EqualError(t, err, "exit status 92", "Bitrise CLI did not abort hanged build, output: %s", out)
}
