package entities

import (
	"errors"

	"github.com/cabralfbenja/segurointeligente/internal/dtos"
)

var (
	ErrInvalidField = errors.New("invalid field")
)

type Insurance struct {
	ID            int64
	UserID        int64
	InsuranceType string
	TimeFrom      string
	TimeTo        string
	Value         float64
}

func NewInsurance(userID int64, input dtos.InsuranceDto) (*Insurance, error) {
	if input.InsuranceType == "" ||
		input.FromTime == "" ||
		input.ToTime == "" {
		return nil, ErrInvalidField
	}

	return &Insurance{
		UserID:        userID,
		InsuranceType: input.InsuranceType,
		TimeFrom:      input.FromTime,
		TimeTo:        input.ToTime,
        Value:         input.Value,
	}, nil
}
