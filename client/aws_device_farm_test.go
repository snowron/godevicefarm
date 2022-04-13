package client

import (
	"errors"
	mocks "godevicefarm/.mocks"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnProjectArnSuccessfully(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockPaginator := mocks.NewMockListProjectsPaginator(mockController)
	output := devicefarm.ListProjectsOutput{
		NextToken: nil,
		Projects: []types.Project{
			{
				Arn:  aws.String("testArn"),
				Name: aws.String("test project"),
			},
		},
	}

	gomock.InOrder(
		mockPaginator.EXPECT().HasMorePages().Return(true).Times(1),
		mockPaginator.EXPECT().NextPage(gomock.Any()).Return(&output, nil).Times(1),
	)
	projectArn, err := searchProjectWithPaginator(mockPaginator, aws.String("test project"))

	assert.Nil(t, err)
	assert.Equal(t, "testArn", projectArn)
}
func TestShouldThrowErrorWhenThereIsNoMatchingProjectName(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockPaginator := mocks.NewMockListProjectsPaginator(mockController)

	gomock.InOrder(
		mockPaginator.EXPECT().HasMorePages().Return(true).Times(1),
		mockPaginator.EXPECT().NextPage(gomock.Any()).Return(nil, errors.New("cant find any project")).Times(1),
	)
	projectArn, err := searchProjectWithPaginator(mockPaginator, aws.String("test project"))

	assert.Error(t, err)
	assert.Equal(t, "", projectArn)
}
func TestShouldReturnFileArnSuccessfully(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockPaginator := mocks.NewMockListUploadsPaginator(mockController)

	output := devicefarm.ListUploadsOutput{
		Uploads: []types.Upload{
			{
				Name:   aws.String("Test File"),
				Arn:    aws.String("Aws test arn"),
				Status: "SUCCEEDED",
			},
		},
	}

	gomock.InOrder(
		mockPaginator.EXPECT().HasMorePages().Return(true).Times(1),
		mockPaginator.EXPECT().NextPage(gomock.Any()).Return(&output, nil).Times(1),
	)
	projectArn, err := searchUploadWithPaginator(mockPaginator, aws.String("Test File"))

	assert.Nil(t, err)
	assert.Equal(t, "Aws test arn", projectArn)
}
func TestShouldThrowErrorWhenThereIsNoMatchingFileName(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockPaginator := mocks.NewMockListUploadsPaginator(mockController)
	output := devicefarm.ListUploadsOutput{
		Uploads: []types.Upload{
			{
				Name:   aws.String("Test File"),
				Arn:    aws.String("Aws test arn"),
				Status: "SUCCEEDED",
			},
		},
	}
	gomock.InOrder(
		mockPaginator.EXPECT().HasMorePages().Return(true).Times(1),
		mockPaginator.EXPECT().NextPage(gomock.Any()).Return(&output, nil).Times(1),
		mockPaginator.EXPECT().HasMorePages().Return(false).Times(1),
	)
	_, err := searchUploadWithPaginator(mockPaginator, aws.String("test file name"))

	assert.Error(t, err)
	assert.Equal(t, "cant find any file with that name on aws test file name", err.Error())
}
func TestShouldReturnDevicePoolArnSuccessfully(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockPaginator := mocks.NewMockListDevicePoolPaginator(mockController)

	output := devicefarm.ListDevicePoolsOutput{
		DevicePools: []types.DevicePool{
			{
				Name: aws.String("Test Device Pool Name"),
				Arn:  aws.String("Aws test arn"),
			},
		},
	}

	gomock.InOrder(
		mockPaginator.EXPECT().HasMorePages().Return(true).Times(1),
		mockPaginator.EXPECT().NextPage(gomock.Any()).Return(&output, nil).Times(1),
	)
	projectArn, err := searchDevicePoolWithPaginator(mockPaginator, aws.String("Test Device Pool Name"))

	assert.Nil(t, err)
	assert.Equal(t, "Aws test arn", projectArn)
}
func TestShouldThrowErrorWhenThereIsNoMatchingDevicePool(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockPaginator := mocks.NewMockListDevicePoolPaginator(mockController)

	gomock.InOrder(
		mockPaginator.EXPECT().HasMorePages().Return(true).Times(1),
		mockPaginator.EXPECT().NextPage(gomock.Any()).Return(nil, errors.New("cant find any match for device pool")).Times(1),
	)
	_, err := searchDevicePoolWithPaginator(mockPaginator, aws.String("test project"))

	assert.Error(t, err)
	assert.Equal(t, "cant find any match for device pool", err.Error())
}

func TestShouldThrowErrorWhenCantGetUploadDataFromAws(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockAwsFarmDevice := mocks.NewMockProviderDeviceFarmInterface(mockController)

	input := devicefarm.GetUploadInput{
		Arn: aws.String("FileArnTest"),
	}

	mockAwsFarmDevice.EXPECT().GetUpload(gomock.Any(), &input).Return(nil, errors.New("error from aws")).Times(1)

	client := DeviceFarmClient{mockAwsFarmDevice}

	_, err := client.GetTestSpecType("FileArnTest")

	assert.Error(t, err)
}
func TestShouldReturnSpecTypeWhenMatchingFile(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockAwsFarmDevice := mocks.NewMockProviderDeviceFarmInterface(mockController)

	input := devicefarm.GetUploadInput{
		Arn: aws.String("FileArnTest"),
	}

	output := devicefarm.GetUploadOutput{
		Upload: &types.Upload{
			Type: "APPIUM_RUBY_TEST_PACKAGE",
		},
	}

	mockAwsFarmDevice.EXPECT().GetUpload(gomock.Any(), &input).Return(&output, nil).Times(1)

	client := DeviceFarmClient{mockAwsFarmDevice}

	specType, err := client.GetTestSpecType("FileArnTest")

	assert.Nil(t, err)
	assert.Equal(t, specType, "APPIUM_RUBY")
}
func TestShouldThrowErrorWhenTheFileHasUnknownOrWrongType(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockAwsFarmDevice := mocks.NewMockProviderDeviceFarmInterface(mockController)

	input := devicefarm.GetUploadInput{
		Arn: aws.String("FileArnTest"),
	}

	output := devicefarm.GetUploadOutput{
		Upload: &types.Upload{
			Type: "APPIUM_RUBY_TEST_PACKAGE_WOW",
		},
	}

	mockAwsFarmDevice.EXPECT().GetUpload(gomock.Any(), &input).Return(&output, nil).Times(1)

	client := DeviceFarmClient{mockAwsFarmDevice}

	_, err := client.GetTestSpecType("FileArnTest")

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Incompatible type for schedule APPIUM_RUBY_TEST_PACKAGE_WOW")
}

func TestShouldThrowErrorWhenCantGetJobsDataFromAws(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockAwsFarmDevice := mocks.NewMockProviderDeviceFarmInterface(mockController)

	input := devicefarm.ListJobsInput{
		Arn: aws.String("RunArnTest"),
	}

	mockAwsFarmDevice.EXPECT().ListJobs(gomock.Any(), &input).Return(nil, errors.New("error from aws")).Times(1)

	client := DeviceFarmClient{mockAwsFarmDevice}

	_, err := client.ListJobsInOneLine("RunArnTest")

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "error from aws")
}
func TestShouldReturnJobInfoSuccessfullyWhenJobCreatedAndRunning(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockAwsFarmDevice := mocks.NewMockProviderDeviceFarmInterface(mockController)

	input := devicefarm.ListJobsInput{
		Arn: aws.String("RunArnTest"),
	}

	output := devicefarm.ListJobsOutput{
		Jobs: []types.Job{
			{
				Name:   aws.String("Samsung Galaxy S9"),
				Result: "RUNNING",
				Status: "RUNNING",
			},
		},
	}

	mockAwsFarmDevice.EXPECT().ListJobs(gomock.Any(), &input).Return(&output, nil).Times(1)

	client := DeviceFarmClient{mockAwsFarmDevice}

	cliOutput, err := client.ListJobsInOneLine("RunArnTest")

	assert.Nil(t, err)
	assert.Equal(t, "Samsung Galaxy S9 RUNNING RUNNING\n", cliOutput)
}
func TestShouldReturnJobsInfoSuccessfullyWhenJobsCreatedAndRunning(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockAwsFarmDevice := mocks.NewMockProviderDeviceFarmInterface(mockController)

	input := devicefarm.ListJobsInput{
		Arn: aws.String("RunArnTest"),
	}

	output := devicefarm.ListJobsOutput{
		Jobs: []types.Job{
			{
				Name:   aws.String("Samsung Galaxy S9"),
				Result: "RUNNING",
				Status: "RUNNING",
			}, {
				Name:   aws.String("Samsung Galaxy S20"),
				Result: "RUNNING",
				Status: "RUNNING",
			},
		},
	}

	mockAwsFarmDevice.EXPECT().ListJobs(gomock.Any(), &input).Return(&output, nil).Times(1)

	client := DeviceFarmClient{mockAwsFarmDevice}

	cliOutput, err := client.ListJobsInOneLine("RunArnTest")

	assert.Nil(t, err)
	assert.Equal(t, cliOutput, "Samsung Galaxy S9 RUNNING RUNNING\nSamsung Galaxy S20 RUNNING RUNNING\n")
}

func TestShouldThrowErrorWhenCantScheduleRun(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockAwsFarmDevice := mocks.NewMockProviderDeviceFarmInterface(mockController)

	input := devicefarm.ScheduleRunInput{
		ProjectArn:                   nil,
		Test:                         nil,
		AppArn:                       nil,
		Configuration:                nil,
		DevicePoolArn:                nil,
		DeviceSelectionConfiguration: nil,
		ExecutionConfiguration:       nil,
		Name:                         nil,
	}

	mockAwsFarmDevice.EXPECT().ScheduleRun(gomock.Any(), &input).Return(nil, errors.New("error from aws")).Times(1)

	client := DeviceFarmClient{mockAwsFarmDevice}

	err := client.StartRun(input)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "error from aws")
}
func TestShouldContinueWhenJobStatusIsSCHEDULING(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockAwsFarmDevice := mocks.NewMockProviderDeviceFarmInterface(mockController)

	input := devicefarm.ScheduleRunInput{
		ProjectArn:                   aws.String("ProjectArn"),
		Test:                         nil,
		AppArn:                       nil,
		Configuration:                nil,
		DevicePoolArn:                nil,
		DeviceSelectionConfiguration: nil,
		ExecutionConfiguration:       nil,
		Name:                         nil,
	}

	scheduleRunOutput := devicefarm.ScheduleRunOutput{
		Run: &types.Run{
			Arn: aws.String("TestRunArnOnJob:run:123/654"),
		},
	}

	getRunOutput := devicefarm.GetRunOutput{
		Run: &types.Run{
			Arn:    aws.String("TestRunArnOnJob:run:123/654"),
			Status: "SCHEDULING",
		},
	}
	listJobsOutput := devicefarm.ListJobsOutput{
		Jobs: []types.Job{
			{
				Name:   aws.String("Samsung Galaxy S9"),
				Result: "SCHEDULING",
				Status: "SCHEDULING",
			},
		},
	}

	mockAwsFarmDevice.EXPECT().ScheduleRun(gomock.Any(), &input).Return(&scheduleRunOutput, nil).Times(1)
	mockAwsFarmDevice.EXPECT().GetRun(gomock.Any(), gomock.Any()).Return(&getRunOutput, nil).Times(1)
	mockAwsFarmDevice.EXPECT().ListJobs(gomock.Any(), gomock.Any()).Return(&listJobsOutput, nil).Times(1)

	// For Break Test
	mockAwsFarmDevice.EXPECT().GetRun(gomock.Any(), gomock.Any()).Return(nil, errors.New("")).Times(1)

	client := DeviceFarmClient{mockAwsFarmDevice}

	_ = client.StartRun(input)
}
func TestShouldContinueWhenJobStatusIsRUNNING(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockAwsFarmDevice := mocks.NewMockProviderDeviceFarmInterface(mockController)

	input := devicefarm.ScheduleRunInput{
		ProjectArn:                   aws.String("ProjectArn"),
		Test:                         nil,
		AppArn:                       nil,
		Configuration:                nil,
		DevicePoolArn:                nil,
		DeviceSelectionConfiguration: nil,
		ExecutionConfiguration:       nil,
		Name:                         nil,
	}

	scheduleRunOutput := devicefarm.ScheduleRunOutput{
		Run: &types.Run{
			Arn: aws.String("TestRunArnOnJob:run:123/654"),
		},
	}

	getRunOutput := devicefarm.GetRunOutput{
		Run: &types.Run{
			Arn:    aws.String("TestRunArnOnJob:run:123/654"),
			Status: "RUNNING",
		},
	}
	listJobsOutput := devicefarm.ListJobsOutput{
		Jobs: []types.Job{
			{
				Name:   aws.String("Samsung Galaxy S9"),
				Result: "RUNNING",
				Status: "RUNNING",
			},
		},
	}

	mockAwsFarmDevice.EXPECT().ScheduleRun(gomock.Any(), &input).Return(&scheduleRunOutput, nil).Times(1)
	mockAwsFarmDevice.EXPECT().GetRun(gomock.Any(), gomock.Any()).Return(&getRunOutput, nil).Times(1)
	mockAwsFarmDevice.EXPECT().ListJobs(gomock.Any(), gomock.Any()).Return(&listJobsOutput, nil).Times(1)

	// For Break Test
	mockAwsFarmDevice.EXPECT().GetRun(gomock.Any(), gomock.Any()).Return(nil, errors.New("")).Times(1)

	client := DeviceFarmClient{mockAwsFarmDevice}

	_ = client.StartRun(input)
}

func TestShouldReturnAWSLinkOfRunSuccessfullyWhenJobCreated(t *testing.T) {
	link := GetAwsLinkOfRun(aws.String("TestRunArn:run:123/654"))
	assert.Equal(t, "https://us-west-2.console.aws.amazon.com/devicefarm/home?region=eu-west-2#/projects/123/runs/654", link)
}
