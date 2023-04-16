package driverlicense_test

import (
	"testing"

	driverlicense "module"

	"github.com/stretchr/testify/suite"
)

func (s *DrivingLicenseSuite) TestUnderageApplication() {
	a := UnderageApplication{}

	lg := driverlicense.NewNumberGenerator()
	_, err := lg.Generate(a)
	s.Error(err)
	s.Contains(err.Error(), "Underaged")

}

type DrivingLicenseSuite struct {
	suite.Suite
}

func TestDrivingLicenseSuite(t *testing.T) {
	suite.Run(t, new(DrivingLicenseSuite))
}

type UnderageApplication struct{}

func (u UnderageApplication) IsOver17() bool {
	return false
}

func (u UnderageApplication) HoldLicense() bool {
	return false
}

func (s *DrivingLicenseSuite) TestNoSecondLicense() {

}

type LicenseHolderApplicat struct {
}
