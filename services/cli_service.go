package services

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"godevicefarm/domain"
)

type CliService struct {
	DeviceFarm DeviceFarm
}

type DeviceFarm interface {
	FindProjectArnWithName(name string) (string, error)
	FindDevicePoolArnWithName(projectArn string, name string) (string, error)
	FindUpload(projectArn string, fileName string) (string, error)
	UploadFile(projectArn string, fileType string, path string) (string, error)
	GetTestSpecType(fileArn string) (string, error)
	StartRun(args devicefarm.ScheduleRunInput) error
}

func (c CliService) StartTestingProcess(section string, cliInput *domain.CliInput) error {
	if section == "mobile" {
		// Check If provided project name exist
		projectArn, _ := c.DeviceFarm.FindProjectArnWithName(*cliInput.ProjectName)

		if projectArn != "" {
			input := devicefarm.ScheduleRunInput{
				ProjectArn: aws.String(projectArn),
				Test: &types.ScheduleRunTest{
					Type:           "",
					TestPackageArn: nil,
					TestSpecArn:    nil,
				},
				AppArn:        nil,
				DevicePoolArn: nil,
				Name:          cliInput.TestName,
			}

			// Check If provided device pool name exist
			err := c.IsDevicePoolExist(projectArn, cliInput, &input)
			if err != nil {
				return err
			}
			// Find which app parameter(name,path) provided find or upload it
			appFileArn, err := c.GetUploadArn(
				projectArn,
				*cliInput.AppPath,
				*cliInput.AppName,
				*cliInput.AppType,
			)

			if err != nil {
				return err
			}

			input.AppArn = aws.String(appFileArn)

			// Find which app parameter(name,path) provided find or upload it
			testSpecFileArn, err := c.GetUploadArn(
				projectArn,
				*cliInput.TestSpecPath,
				*cliInput.TestSpecName,
				*cliInput.TestSpecType,
			)
			if err != nil {
				return err
			}
			input.Test.TestPackageArn = aws.String(testSpecFileArn)

			// Find which app parameter(name,path) provided find or upload it
			testSpecConfArn, err := c.GetUploadArn(
				projectArn,
				*cliInput.TestSpecConfigurationPath,
				*cliInput.TestSpecConfigurationName,
				*cliInput.TestSpecConfigurationType,
			)

			if err != nil {
				return err
			}

			input.Test.TestSpecArn = aws.String(testSpecConfArn)

			err = c.GetTestSpecType(&input)

			if err != nil {
				return err
			}

			err = c.DeviceFarm.StartRun(input)

			if err != nil {
				return err
			}
		}
		return errors.New("Unable to find project with provided name 😡 " + *cliInput.ProjectName)
	}
	return errors.New("section is not allowed")
}

func (c CliService) GetUploadArn(projectArn, fieldPath, fieldName, fileType string) (string, error) {
	if fieldPath != "" {
		fileArn, err := c.DeviceFarm.UploadFile(projectArn, fileType, fieldPath)

		if err != nil {
			return "", err
		}

		return fileArn, nil
	}

	fileArn, err := c.DeviceFarm.FindUpload(projectArn, fieldName)

	if err != nil {
		return "", err
	}

	return fileArn, nil
}

func (c CliService) IsDevicePoolExist(projectArn string, cliInput *domain.CliInput, scheduleInput *devicefarm.ScheduleRunInput) error {
	existPoolArn, err := c.DeviceFarm.FindDevicePoolArnWithName(projectArn, *cliInput.DevicePoolName)

	if err != nil {
		return err
	}

	scheduleInput.DevicePoolArn = aws.String(existPoolArn)

	return nil
}

func (c CliService) GetTestSpecType(input *devicefarm.ScheduleRunInput) error {
	testType, err := c.DeviceFarm.GetTestSpecType(*input.Test.TestPackageArn)

	if err != nil {
		return err
	}

	input.Test.Type = types.TestType(testType)

	return nil
}
