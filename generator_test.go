package driverlicense_test

import "github.com/stretchr/testify/suite"

func (s *DrivingLicenseSuite)TestUnderageApplication(){
	a := UnderageApplication{}

	generatorNuber := driverlicense.NewNumberGenerator()
	_, err := generatorNuber.Generate(a)
	if err != nil{
		log.err(err)
	}
}

type DrivingLicenseSuite struct {
	suite.Suite
}

func TestDrivingLicenseSuite(t *testing.T) {
	suite.Run(t, new(DrivingLicenseSuite))
}

type UnderageApplication struct{
	driverlicense.Applicant{}
}

type (u UnderageApplication) IsOver17()bool{
	return false
	
}

type (u UnderageApplication) HoldLicense()bool{
	return false
}