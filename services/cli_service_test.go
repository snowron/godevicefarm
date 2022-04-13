package services

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mocks "godevicefarm/.mocks"
	"godevicefarm/domain"
	"testing"
)

func TestShouldFindFileArnAndReturnItWhenUserProvidesOnlyNameParameter(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	awsDeviceFarm := mocks.NewMockDeviceFarm(mockController)

	projectArn := "projectArnTesting"
	fieldPath := ""
	fieldName := "apk-app.apk"
	fileType := ""

	awsDeviceFarm.EXPECT().FindUpload(projectArn, fieldName).Return("awsArnData", nil).Times(1)
	service := CliService{DeviceFarm: awsDeviceFarm}
	arn, err := service.GetUploadArn(projectArn, fieldPath, fieldName, fileType)

	assert.Empty(t, err)
	assert.Equal(t, arn, "awsArnData")
}
func TestShouldThrowErrorIfItIsNotExistWhenUserProvidesOnlyNameParameter(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	awsDeviceFarm := mocks.NewMockDeviceFarm(mockController)

	projectArn := "projectArnTesting"
	fieldPath := ""
	fieldName := "apk-app.apk"
	fileType := ""
	fakeError := errors.New("cant find any match file: " + fieldName)

	awsDeviceFarm.EXPECT().FindUpload(projectArn, fieldName).Return("", fakeError).Times(1)
	service := CliService{DeviceFarm: awsDeviceFarm}

	arn, err := service.GetUploadArn(projectArn, fieldPath, fieldName, fileType)

	assert.Empty(t, arn)
	assert.Error(t, err)
	assert.Equal(t, fakeError, err)
}
func TestShouldUploadAndReturnArWhenUserProvidesPathAndTypeParameter(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	awsDeviceFarm := mocks.NewMockDeviceFarm(mockController)

	projectArn := "projectArnTesting"
	fieldPath := "./apk-app.apk"
	fieldName := ""
	fileType := "ANDROID_APP"

	awsDeviceFarm.EXPECT().UploadFile(projectArn, fileType, fieldPath).Return("awsArnData", nil).Times(1)
	service := CliService{DeviceFarm: awsDeviceFarm}

	arn, err := service.GetUploadArn(projectArn, fieldPath, fieldName, fileType)

	assert.Empty(t, err)
	assert.Equal(t, arn, "awsArnData")
}
func TestShouldThrowErrorWhenUserProvidesPathAndTypeParameterButFileDoesNotExist(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	awsDeviceFarm := mocks.NewMockDeviceFarm(mockController)

	projectArn := "projectArnTesting"
	fieldPath := "./apk-app.apkkk"
	fieldName := ""
	fileType := "ANDROID_APP"
	fakeError := errors.New("cant upload file: " + fieldPath)
	awsDeviceFarm.EXPECT().UploadFile(projectArn, fileType, fieldPath).Return("", fakeError).Times(1)
	service := CliService{DeviceFarm: awsDeviceFarm}

	arn, err := service.GetUploadArn(projectArn, fieldPath, fieldName, fileType)

	assert.Empty(t, arn)
	assert.Error(t, err)
	assert.Equal(t, fakeError, err)
}
func TestShouldUploadAndReturnArnFromPathWhenUserProvidesPathAndTypeAndNameParameter(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	awsDeviceFarm := mocks.NewMockDeviceFarm(mockController)

	projectArn := "projectArnTesting"
	fieldPath := "./apk-app.apk"
	//TODO hangisi dogru law
	fieldName := "apk-app.apk"
	fileType := "ANDROID_APP"

	awsDeviceFarm.EXPECT().UploadFile(projectArn, fileType, fieldPath).Return("awsArnData", nil).Times(1)
	service := CliService{DeviceFarm: awsDeviceFarm}

	arn, err := service.GetUploadArn(projectArn, fieldPath, fieldName, fileType)

	assert.Empty(t, err)
	assert.Equal(t, arn, "awsArnData")
}

func TestShouldThrowErrorWhenProjectNameDoesNotExistOnDeviceFarm(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	awsDeviceFarm := mocks.NewMockDeviceFarm(mockController)

	cliInput := domain.CliInput{
		ProjectName:    aws.String("projectArnTesting"),
		DevicePoolName: aws.String("testPoolForOneDevice"),
	}

	awsDeviceFarm.EXPECT().FindProjectArnWithName("projectArnTesting").Return("", nil).Times(1)
	service := CliService{DeviceFarm: awsDeviceFarm}

	err := service.StartTestingProcess("mobile", &cliInput)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Unable to find project with provided name 😡 projectArnTesting")
}

func TestShouldGetDevicePoolExistStatusSuccessfullyWhenDevicePoolNameExistOnDeviceFarm(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	awsDeviceFarm := mocks.NewMockDeviceFarm(mockController)

	cliInput := domain.CliInput{
		ProjectName:    aws.String("projectArnTesting"),
		DevicePoolName: aws.String("testPoolForOneDevice"),
	}
	scheduleRunInput := devicefarm.ScheduleRunInput{
		DevicePoolArn: nil,
	}
	awsDeviceFarm.EXPECT().FindDevicePoolArnWithName("projectArnTesting", "testPoolForOneDevice").Return("awsArnData", nil).Times(1)
	service := CliService{DeviceFarm: awsDeviceFarm}

	err := service.IsDevicePoolExist("projectArnTesting", &cliInput, &scheduleRunInput)

	assert.Empty(t, err)
	assert.Equal(t, scheduleRunInput.DevicePoolArn, aws.String("awsArnData"))
}
func TestShouldThrowErrorWhenDevicePoolNameDoesNotExistOnDeviceFarm(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	awsDeviceFarm := mocks.NewMockDeviceFarm(mockController)

	cliInput := domain.CliInput{
		ProjectName:    aws.String("projectArnTesting"),
		DevicePoolName: aws.String("testPoolForOneDevice"),
	}
	scheduleRunInput := devicefarm.ScheduleRunInput{
		DevicePoolArn: nil,
	}
	fakeError := errors.New("cant find any match for device pool")
	awsDeviceFarm.EXPECT().FindDevicePoolArnWithName("projectArnTesting", "testPoolForOneDevice").Return("", fakeError).Times(1)
	service := CliService{DeviceFarm: awsDeviceFarm}

	err := service.IsDevicePoolExist("projectArnTesting", &cliInput, &scheduleRunInput)

	assert.Error(t, err)
	assert.Equal(t, fakeError, err)
}

func TestShouldGetTestSpecTypeSuccessfullyWhenUserProvidesTestSpecReturnType(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	awsDeviceFarm := mocks.NewMockDeviceFarm(mockController)

	scheduleRunInput := devicefarm.ScheduleRunInput{
		Test: &types.ScheduleRunTest{
			TestPackageArn: aws.String("TestArnTestPackage"),
		},
	}
	awsDeviceFarm.EXPECT().GetTestSpecType("TestArnTestPackage").Return("APPIUM_NODE", nil).Times(1)
	service := CliService{DeviceFarm: awsDeviceFarm}

	err := service.GetTestSpecType(&scheduleRunInput)

	assert.Empty(t, err)
	assert.Equal(t, scheduleRunInput.Test.Type, types.TestType("APPIUM_NODE"))
}
func TestShouldThrowErrorWhenUserProvidesTestSpecButItsTypeNotFitScheduleTypeInput(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	awsDeviceFarm := mocks.NewMockDeviceFarm(mockController)

	scheduleRunInput := devicefarm.ScheduleRunInput{
		Test: &types.ScheduleRunTest{
			TestPackageArn: aws.String("TestArnTestPackage"),
		},
	}
	fakeError := errors.New("incompatible type for schedule testInCompatible")
	awsDeviceFarm.EXPECT().GetTestSpecType("TestArnTestPackage").Return("", fakeError).Times(1)
	service := CliService{DeviceFarm: awsDeviceFarm}

	err := service.GetTestSpecType(&scheduleRunInput)

	assert.Error(t, err)
	assert.Equal(t, err, fakeError)
}

func TestShouldStartTestingProcessSuccessfullyWhenUserProvidesTestRequirementsWithOnlyNames(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	awsDeviceFarm := mocks.NewMockDeviceFarm(mockController)

	cliInput := domain.CliInput{
		ProjectName:               aws.String("testProjectName"),
		DevicePoolName:            aws.String("testDevicePoolName"),
		TestSpecType:              aws.String(""),
		TestSpecConfigurationType: aws.String(""),
		AppPath:                   aws.String(""),
		TestSpecConfigurationPath: aws.String(""),
		TestSpecPath:              aws.String(""),
		AppName:                   aws.String("testAppNameOnDeviceFarm.apk"),
		AppType:                   aws.String(""),
		TestSpecConfigurationName: aws.String("testTestSpecConfigurationName"),
		TestSpecName:              aws.String("testTestSpecName"),
		TestName:                  aws.String("testNameDeviceFarm"),
	}

	awsDeviceFarm.EXPECT().FindProjectArnWithName(gomock.Any()).Return("testProjectName", nil).Times(1)
	awsDeviceFarm.EXPECT().FindDevicePoolArnWithName(gomock.Any(), gomock.Any()).Return("arnDevicePool", nil).Times(1)
	awsDeviceFarm.EXPECT().FindUpload(gomock.Any(), gomock.Any()).Return("arnFile", nil).Times(3)
	awsDeviceFarm.EXPECT().GetTestSpecType(gomock.Any()).Return("APPIUM_NODE", nil).Times(1)
	awsDeviceFarm.EXPECT().StartRun(gomock.Any()).Return(nil).Times(1)

	service := CliService{DeviceFarm: awsDeviceFarm}

	err := service.StartTestingProcess("mobile", &cliInput)

	assert.Empty(t, err)
}
func TestShouldStartTestingProcessSuccessfullyWhenUserProvidesTestRequirementsWithOnlyPaths(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	awsDeviceFarm := mocks.NewMockDeviceFarm(mockController)

	cliInput := domain.CliInput{
		ProjectName:               aws.String("testProjectName"),
		DevicePoolName:            aws.String("testDevicePoolName"),
		TestSpecType:              aws.String("APPIUM_NODE_TEST_PACKAGE"),
		TestSpecConfigurationType: aws.String("APPIUM_NODE_SPEC_TYPE"),
		AppPath:                   aws.String("./testAppNameOnDeviceFarm.apk"),
		TestSpecConfigurationPath: aws.String("./appium-spec.yml"),
		TestSpecPath:              aws.String("./mytests.zip"),
		AppName:                   aws.String(""),
		AppType:                   aws.String(""),
		TestSpecConfigurationName: aws.String(""),
		TestSpecName:              aws.String(""),
		TestName:                  aws.String("testNameDeviceFarm"),
	}

	awsDeviceFarm.EXPECT().FindProjectArnWithName(gomock.Any()).Return("arn", nil).Times(1)
	awsDeviceFarm.EXPECT().FindDevicePoolArnWithName(gomock.Any(), gomock.Any()).Return("arnDevicePool", nil).Times(1)
	awsDeviceFarm.EXPECT().UploadFile(gomock.Any(), gomock.Any(), gomock.Any()).Return("arnFile", nil).Times(3)
	awsDeviceFarm.EXPECT().GetTestSpecType(gomock.Any()).Return("APPIUM_NODE", nil).Times(1)
	awsDeviceFarm.EXPECT().StartRun(gomock.Any()).Return(nil).Times(1)

	service := CliService{DeviceFarm: awsDeviceFarm}

	err := service.StartTestingProcess("mobile", &cliInput)

	assert.Empty(t, err)
}

func TestShouldThrowErrorWhenSectionIsNotAllowed(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	cliInput := domain.CliInput{
		ProjectName:               aws.String("testProjectName"),
		DevicePoolName:            aws.String("testDevicePoolName"),
		TestSpecType:              aws.String(""),
		TestSpecConfigurationType: aws.String(""),
		AppPath:                   aws.String(""),
		TestSpecConfigurationPath: aws.String(""),
		TestSpecPath:              aws.String(""),
		AppName:                   aws.String("testAppNameOnDeviceFarm.apk"),
		AppType:                   aws.String(""),
		TestSpecConfigurationName: aws.String("testTestSpecConfigurationName"),
		TestSpecName:              aws.String("testTestSpecName"),
		TestName:                  aws.String("testNameDeviceFarm"),
	}

	awsDeviceFarm := mocks.NewMockDeviceFarm(mockController)
	service := CliService{DeviceFarm: awsDeviceFarm}
	err := service.StartTestingProcess("", &cliInput)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "section is not allowed")
}
