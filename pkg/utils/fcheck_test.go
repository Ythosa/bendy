package utils_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ythosa/bendy/pkg/utils"
)

func TestCheckIsFilePathsValid(t *testing.T) {
	t.Parallel()

	const (
		validFilePath   = "./valid_file_path"
		invalidFilePath = "./invalid_file_path"
	)

	_, _ = os.Create(validFilePath)

	testCases := []struct {
		filePaths     []string
		expectedError bool
	}{
		{
			filePaths:     []string{validFilePath},
			expectedError: false,
		},
		{
			filePaths:     []string{invalidFilePath},
			expectedError: true,
		},
		{
			filePaths:     []string{validFilePath, invalidFilePath},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		if tc.expectedError {
			assert.NotNil(t, utils.CheckIsFilePathsValid(tc.filePaths))
		} else {
			assert.Nil(t, utils.CheckIsFilePathsValid(tc.filePaths))
		}
	}

	_ = os.Remove(validFilePath)
}
