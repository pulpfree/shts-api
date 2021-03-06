package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

// IntegSuite struct
type IntegSuite struct {
	suite.Suite
	c *Config
}

// SetupTest method
func (suite *IntegSuite) SetupTest() {

	suite.c = &Config{}

	os.Setenv("Stage", "test")
	suite.c.setDefaults()
	suite.c.setEnvVars()
}

// TestSetSSMParams function
// this test assumes that the S3Bucket is set
/* func (suite *IntegSuite) TestSetSSMParams() {

	DBNameBefore := defs.DBName
	err := suite.c.setSSMParams()
	suite.NoError(err)

	DBNameAfter := defs.DBName
	suite.True(strings.Compare(DBNameBefore, DBNameAfter) != 0)
} */

// TestLoad function
func (suite *IntegSuite) TestLoad() {

	suite.Empty(suite.c.AWSRegion)

	err := suite.c.Load()
	suite.NoError(err)
	suite.NotEmpty(suite.c.AWSRegion)
}

// TestLoadProduction
func (suite *IntegSuite) TestLoadProduction() {

	os.Setenv("Stage", "dev")

	err := suite.c.Load()
	suite.NoError(err)
	suite.NotEmpty(suite.c.AWSRegion)
	fmt.Printf("suite.c %+v\n", suite.c)
}

func (suite *IntegSuite) TestSetStageEnv() {
	envBefore := suite.c.GetStageEnv()
	suite.True(envBefore == suite.c.GetStageEnv())
	suite.c.SetStageEnv(string(ProdEnv))
	envAfter := suite.c.GetStageEnv()
	suite.True(envAfter == ProdEnv)
}

// TestIntegrationSuite function
func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegSuite))
}
