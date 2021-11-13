package main

import (
	"testing"

	. "github.com/giladreich/go-errori18n/errori18n"
	"github.com/stretchr/testify/require"
)

func TestDivide(t *testing.T) {
	testcases := []struct {
		name         string
		leftOperand  int
		rightOperand int
		expectedRet  int
		expectedErr  *Error
	}{
		{
			name:         "Good: proper division",
			leftOperand:  8,
			rightOperand: 4,
			expectedRet:  2,
			expectedErr:  nil,
		},
		{
			name:         "Bad: divide by zero",
			leftOperand:  8,
			rightOperand: 0,
			expectedRet:  0,
			expectedErr:  Errorf(ERR_DIVIDE_BY_ZERO, 8),
		},
		{
			name:         "Bad: divide by zero code",
			leftOperand:  8,
			rightOperand: 0,
			expectedRet:  0,
			expectedErr:  Errorf(ERR_UNKNOWN), // fail on purpose to see the nice fail message
		},
		{
			name:         "Bad: divide by zero string",
			leftOperand:  8,
			rightOperand: 0,
			expectedRet:  0,
			expectedErr:  Errorf(ERR_DIVIDE_BY_ZERO, 6), // fail on purpose to see the nice fail message
		},
	}

	// Need to change language during runtime? no problem.
	SetLanguage(LANG_DE)

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actualRet, actualErr := Divide(tc.leftOperand, tc.rightOperand)
			if actualErr != nil {
				require.NotNil(t, tc.expectedErr)

				// Either test against error codes
				require.Equal(t, tc.expectedErr.ErrorCode, actualErr.ErrorCode)

				// or strings for nice output in failed tests
				require.Equal(t, tc.expectedErr.String(), actualErr.String())

				require.Equal(t, tc.expectedRet, 0)
			} else {
				require.Nil(t, tc.expectedErr)
				require.Equal(t, tc.expectedRet, actualRet)
			}
		})
	}
}
