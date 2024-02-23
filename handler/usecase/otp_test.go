package usecase

import (
	"burmese_jewellery/config"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func Test_GenOTP(t *testing.T) {
	rList := []string{}
	for i := 1; i < 100; i++ {
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
	config.InitMock()

	tList := map[string]struct {
		params time.Time
		want   bool
	}{
		"ok/2mins": {
			params: time.Now().Add(-time.Minute * 2),
			want:   false,
		},
		"ok/now": {
			params: time.Now(),
			want:   false,
		},
		"ok/4mins": {
			params: time.Now().Add(-time.Minute * 4),
			want:   true,
		},
	}

	for name, tt := range tList {
		t.Run(name, func(t *testing.T) {
			got := IsOTPExpired(tt.params)
			assert.Equal(t, tt.want, got)
		})
	}
}
