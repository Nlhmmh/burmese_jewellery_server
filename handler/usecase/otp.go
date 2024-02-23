package usecase

import (
	"burmese_jewellery/config"
	cryptoRand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

const (
	max = 999999
	min = 000000
)

func GenSafeOTP() (string, error) {
	randomNum, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(900000))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", randomNum.Int64()+100000), nil
}

func GenOTP() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06d", r.Intn(max-min)+min)
}

func IsOTPExpired(otpUpdatedAt time.Time) bool {
	return time.Now().After(
		otpUpdatedAt.Add(
			time.Minute * time.Duration(config.Get().OTPExpiredMinutes),
		),
	)
}
