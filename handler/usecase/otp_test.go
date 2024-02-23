package usecase

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func Test_GenOTP(t *testing.T) {
	rList := []string{}
	for i := 1; i < 1000; i++ {
		t.Run(
			fmt.Sprintf("ok/%d", i), func(t *testing.T) {
				otp := GenOTP()
				assert.NotEmpty(t, otp)
				assert.False(t, slices.Contains(rList, otp))
				rList = append(rList, otp)
			},
		)
	}
}
