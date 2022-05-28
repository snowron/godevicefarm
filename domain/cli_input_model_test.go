package domain

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestShouldReturnTrueWhenConfigurationTypeExistOnDeviceFarm(t *testing.T) {
	cliInput := CliInput{
		TestSpecConfigurationPath: aws.String("./wow/test.yml"),
		TestSpecConfigurationType: aws.String("APPIUM_NODE_TEST_SPEC"),
	}

	exist := cliInput.IsTestConfigurationTypeExist()
	assert.True(t, exist)
}
func TestShouldReturnFalseWhenConfigurationTypeDoesNotExistOnDeviceFarm(t *testing.T) {
	cliInput := CliInput{
		TestSpecConfigurationPath: aws.String("./wow/test.yml"),
		TestSpecConfigurationType: aws.String("APPIUM_WEB_NODE_TEST_SPECS"),
	}

	exist := cliInput.IsTestConfigurationTypeExist()
	assert.False(t, exist)
}

func TestShouldReturnTrueWhenSpecTypeExistOnDeviceFarm(t *testing.T) {
	cliInput := CliInput{
		TestSpecPath: aws.String("./wow/test.zip"),
		TestSpecType: aws.String("APPIUM_NODE_TEST_PACKAGE"),
	}

	exist := cliInput.IsTestTypeExist()
	assert.True(t, exist)
}
func TestShouldReturnFalseWhenSpecTypeDoesNotExistOnDeviceFarm(t *testing.T) {
	cliInput := CliInput{
		TestSpecPath: aws.String("./wow/test.zip"),
		TestSpecType: aws.String("NODE_PACKAGE"),
	}

	exist := cliInput.IsTestTypeExist()
	assert.False(t, exist)
}

func TestShouldThrowErrorWhenProvidedSpecTypeIsNotAvailableOnDeviceFarm(t *testing.T) {
	cliInput := CliInput{
		TestSpecName:              aws.String("conf"),
		TestSpecPath:              aws.String("./mytest.zip"),
		TestSpecType:              aws.String("nODe_PACQage"),
		TestSpecConfigurationName: aws.String(""),
		TestSpecConfigurationPath: aws.String("./conf.yml"),
	}

	err := cliInput.Analysis()

	assert.Error(t, err)
	assert.Equal(t, "not exist test spec type", err.Error())
}
func TestShouldThrowErrorWhenProvidedSpecConfTypeIsNotAvailableOnDeviceFarm(t *testing.T) {
	cliInput := CliInput{
		TestSpecName:              aws.String(""),
		TestSpecPath:              aws.String("./mytest.zip"),
		TestSpecType:              aws.String("APPIUM_JAVA_JUNIT_TEST_PACKAGE"),
		TestSpecConfigurationName: aws.String(""),
		TestSpecConfigurationPath: aws.String("./conf.yml"),
		TestSpecConfigurationType: aws.String("faketype"),
	}

	err := cliInput.Analysis()

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "not exist test spec configuration type")
}

func TestShouldThrowErrorWhenSpecNamePathAndTypeNotExist(t *testing.T) {
	cliInput := CliInput{
		TestSpecName: aws.String(""),
		TestSpecPath: aws.String(""),
		TestSpecType: aws.String(""),
	}

	err := cliInput.Analysis()

	assert.Error(t, err)
	assert.Equal(t, "-testSpecPath or -testSpecName must be exist", err.Error())
}
func TestShouldThrowErrorWhenTestSpecConfigurationNamePathAndTypeNotExist(t *testing.T) {
	cliInput := CliInput{
		TestSpecName:              aws.String("Fake Name"),
		TestSpecPath:              aws.String(""),
		TestSpecType:              aws.String(""),
		TestSpecConfigurationName: aws.String(""),
		TestSpecConfigurationPath: aws.String(""),
		TestSpecConfigurationType: aws.String(""),
	}

	err := cliInput.Analysis()

	assert.Error(t, err)
	assert.Equal(t, "-testSpecConfigurationName or -testSpecConfigurationName must be exist", err.Error())
}

func TestShouldThrowErrorWhenAppNameOrAppPathIsEmpty(t *testing.T) {
	cliInput := CliInput{
		TestSpecName:              aws.String(""),
		TestSpecPath:              aws.String("./mytest.zip"),
		TestSpecType:              aws.String("APPIUM_JAVA_JUNIT_TEST_PACKAGE"),
		TestSpecConfigurationName: aws.String(""),
		TestSpecConfigurationPath: aws.String("./conf.yml"),
		TestSpecConfigurationType: aws.String("APPIUM_JAVA_JUNIT_TEST_SPEC"),
		AppName:                   aws.String(""),
		AppPath:                   aws.String(""),
		AppType:                   aws.String(""),
	}
	err := cliInput.Analysis()

	assert.Error(t, err)
	assert.Equal(t, "-appName or -appPath must be exist", err.Error())
}
func TestShouldThrowErrorWhenAppPathDoesNotHaveMobileExtension(t *testing.T) {
	cliInput := CliInput{
		TestSpecName:              aws.String(""),
		TestSpecPath:              aws.String("./mytest.zip"),
		TestSpecType:              aws.String("APPIUM_JAVA_JUNIT_TEST_PACKAGE"),
		TestSpecConfigurationName: aws.String(""),
		TestSpecConfigurationPath: aws.String("./conf.yml"),
		TestSpecConfigurationType: aws.String("APPIUM_JAVA_JUNIT_TEST_SPEC"),
		AppPath:                   aws.String("./app.as"),
		AppType:                   aws.String(""),
		AppName:                   aws.String(""),
	}
	err := cliInput.Analysis()

	assert.Equal(t, "app file does not have .ipa or .apk extension ./app.as", err.Error())
	assert.Error(t, err)
}
func TestShouldSetAsANDROID_APPWhenPathHasDotApkExtension(t *testing.T) {
	cliInput := CliInput{
		TestSpecName:              aws.String(""),
		TestSpecPath:              aws.String("./mytest.zip"),
		TestSpecType:              aws.String("APPIUM_JAVA_JUNIT_TEST_PACKAGE"),
		TestSpecConfigurationPath: aws.String("./conf.yml"),
		TestSpecConfigurationName: aws.String(""),
		TestSpecConfigurationType: aws.String("APPIUM_JAVA_JUNIT_TEST_SPEC"),
		AppPath:                   aws.String("./app.apk"),
		AppType:                   aws.String(""),
		AppName:                   aws.String(""),
	}
	err := cliInput.Analysis()

	assert.Nil(t, err)
	assert.Equal(t, aws.String("ANDROID_APP"), cliInput.AppType)
}
func TestShouldSetAsIOS_APPWhenPathHasDotIpaExtension(t *testing.T) {
	cliInput := CliInput{
		TestSpecName:              aws.String(""),
		TestSpecPath:              aws.String("./mytest.zip"),
		TestSpecType:              aws.String("APPIUM_JAVA_JUNIT_TEST_PACKAGE"),
		TestSpecConfigurationName: aws.String(""),
		TestSpecConfigurationPath: aws.String("./conf.yml"),
		TestSpecConfigurationType: aws.String("APPIUM_JAVA_JUNIT_TEST_SPEC"),
		AppPath:                   aws.String("./app.ipa"),
		AppType:                   aws.String(""),
		AppName:                   aws.String(""),
	}
	err := cliInput.Analysis()

	assert.Nil(t, err)
	assert.Equal(t, aws.String("IOS_APP"), cliInput.AppType)
}
