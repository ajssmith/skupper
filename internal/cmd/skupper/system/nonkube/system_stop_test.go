package nonkube

import (
	"fmt"
	"testing"

	"github.com/skupperproject/skupper/internal/cmd/skupper/common"
	"github.com/skupperproject/skupper/internal/cmd/skupper/common/testutils"
	"gotest.tools/v3/assert"
)

func TestCmdSystemTearDown_ValidateInput(t *testing.T) {
	type test struct {
		name          string
		args          []string
		expectedError string
	}

	testTable := []test{
		{
			name:          "arg-not-accepted",
			args:          []string{"namespace"},
			expectedError: "this command does not accept arguments",
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {

			command := &CmdSystemStop{}
			command.CobraCmd = common.ConfigureCobraCommand(common.PlatformLinux, common.SkupperCmdDescription{}, command, nil)

			testutils.CheckValidateInput(t, command, test.expectedError, test.args)
		})
	}
}

func TestCmdSystemTearDown_InputToOptions(t *testing.T) {

	type test struct {
		name              string
		args              []string
		namespace         string
		expectedNamespace string
	}

	testTable := []test{
		{
			name:              "options-by-default",
			expectedNamespace: "default",
		},
		{
			name:              "namespace-provided",
			args:              []string{"east"},
			namespace:         "east",
			expectedNamespace: "east",
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {

			cmd := newCmdSystemTeardownWithMocks(false)
			cmd.Namespace = test.namespace
			cmd.InputToOptions()

			assert.Check(t, cmd.Namespace == test.expectedNamespace)

		})
	}
}

func TestCmdSystemTeardown_Run(t *testing.T) {
	type test struct {
		name          string
		teardownFails bool
		errorMessage  string
	}

	testTable := []test{
		{
			name:          "runs ok",
			teardownFails: false,
			errorMessage:  "",
		},
		{
			name:          "teardown fails",
			teardownFails: true,
			errorMessage:  "System teardown has failed: fail",
		},
	}

	for _, test := range testTable {
		command := newCmdSystemTeardownWithMocks(test.teardownFails)

		t.Run(test.name, func(t *testing.T) {

			err := command.Run()
			if err != nil {
				assert.Check(t, test.errorMessage == err.Error())
			} else {
				assert.Check(t, err == nil)
			}
		})
	}
}

// --- helper methods

func newCmdSystemTeardownWithMocks(systemTeardDownFails bool) *CmdSystemStop {

	cmdMock := &CmdSystemStop{
		TearDown: mockCmdSystemTeardown,
	}
	if systemTeardDownFails {
		cmdMock.TearDown = mockCmdSystemTeardownFails
	}

	return cmdMock
}

func mockCmdSystemTeardown(namespace string) error { return nil }
func mockCmdSystemTeardownFails(namespace string) error {
	return fmt.Errorf("fail")
}
