package usecase

import (
	"fmt"
	"testing"
	"time"

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

func Test_IsOTPExpired(t *testing.T) {
	t.Run(
		"true", func(t *testing.T) {
			got := IsOTPExpired(time.Now().Add(-time.Minute * 4))
			assert.Equal(t, false, got)
		},
	)
	t.Run(
		"false", func(t *testing.T) {
			got := IsOTPExpired(time.Now().Add(-time.Minute * 6))
			assert.Equal(t, true, got)
		},
	)
}
