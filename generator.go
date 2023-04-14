package driverlicense

type Applicant interface {
	isOver17() bool
	HoldLicense() bool
}

type NumberGenerator struct {
}

func (g NumberGenerator) Generate(a Applicant) (string, error) {
	return "", nil
}

func NewNumberGenerator() NumberGenerator {
	return NewNumberGenerator()
}
