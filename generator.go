package driverlicense

import "errors"

type Applicant interface {
	isOver17() bool
	HoldLicense() bool
}

type NumberGenerator struct {
}

func (g NumberGenerator) Generate(a Applicant) (string, error) {
	return "", errors.New("Underaged Applicat, vou must be 17 for license")
}

func NewNumberGenerator() NumberGenerator {
	return NewNumberGenerator()
}
