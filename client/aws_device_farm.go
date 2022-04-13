package client

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/k0kubun/pp"
)

type DeviceFarmClient struct {
	DeviceFarm ProviderDeviceFarmInterface
}

type ProviderDeviceFarmInterface interface {
	devicefarm.ListProjectsAPIClient
	devicefarm.ListUploadsAPIClient
	devicefarm.ListDevicePoolsAPIClient
	devicefarm.ListJobsAPIClient

	GetUpload(ctx context.Context, params *devicefarm.GetUploadInput,
		optFns ...func(*devicefarm.Options)) (*devicefarm.GetUploadOutput, error)
	ScheduleRun(ctx context.Context, params *devicefarm.ScheduleRunInput,
		optFns ...func(*devicefarm.Options)) (*devicefarm.ScheduleRunOutput, error)
	GetRun(ctx context.Context, params *devicefarm.GetRunInput,
		optFns ...func(*devicefarm.Options)) (*devicefarm.GetRunOutput, error)
	ListArtifacts(ctx context.Context, params *devicefarm.ListArtifactsInput,
		optFns ...func(*devicefarm.Options)) (*devicefarm.ListArtifactsOutput, error)
	CreateUpload(ctx context.Context, params *devicefarm.CreateUploadInput,
		optFns ...func(*devicefarm.Options)) (*devicefarm.CreateUploadOutput, error)
}

const SUCCEEDED = "SUCCEEDED"

type ListProjectsPaginator interface {
	HasMorePages() bool
	NextPage(context.Context, ...func(*devicefarm.Options)) (*devicefarm.ListProjectsOutput, error)
}

type ListUploadsPaginator interface {
	HasMorePages() bool
	NextPage(context.Context, ...func(*devicefarm.Options)) (*devicefarm.ListUploadsOutput, error)
}

type ListDevicePoolPaginator interface {
	HasMorePages() bool
	NextPage(context.Context, ...func(*devicefarm.Options)) (*devicefarm.ListDevicePoolsOutput, error)
}

func (a DeviceFarmClient) FindProjectArnWithName(name string) (string, error) {
	p := devicefarm.NewListProjectsPaginator(a.DeviceFarm, nil)
	projectArn, err := searchProjectWithPaginator(p, aws.String(name))
	if err != nil {
		return "", err
	}
	return projectArn, nil
}

func searchProjectWithPaginator(p ListProjectsPaginator, name *string) (string, error) {
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())

		if err != nil {
			return "", err
		}

		for _, project := range page.Projects {
			if strings.Compare(*project.Name, *name) == 0 {
				pp.Println("👍 The project has found and gathering ARN 👍")
				return *project.Arn, nil
			}
		}
	}
	return "", errors.New("cant find any project")
}

func (a DeviceFarmClient) FindUpload(projectArn, fileName string) (string, error) {
	input := devicefarm.ListUploadsInput{
		Arn: aws.String(projectArn),
	}

	p := devicefarm.NewListUploadsPaginator(a.DeviceFarm, &input)
	fileArn, err := searchUploadWithPaginator(p, aws.String(fileName))

	if err != nil {
		return "", err
	}
	return fileArn, nil
}

func searchUploadWithPaginator(p ListUploadsPaginator, name *string) (string, error) {
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return "", err
		}
		for _, file := range page.Uploads {
			if strings.Compare(*file.Name, *name) == 0 && file.Status == SUCCEEDED {
				pp.Println("👍 The file has found and gathering ARN 👍 ", *file.Name)
				return *file.Arn, nil
			}
		}
	}
	return "", errors.New("cant find any file with that name on aws " + *name)
}

func (a DeviceFarmClient) FindDevicePoolArnWithName(projectArn, name string) (string, error) {
	input := &devicefarm.ListDevicePoolsInput{
		Arn: aws.String(projectArn),
	}

	p := devicefarm.NewListDevicePoolsPaginator(a.DeviceFarm, input)
	devicePoolArn, err := searchDevicePoolWithPaginator(p, aws.String(name))

	if err != nil {
		return "", err
	}

	return devicePoolArn, nil
}

func searchDevicePoolWithPaginator(p ListDevicePoolPaginator, name *string) (string, error) {
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return "", err
		}
		for _, pool := range page.DevicePools {
			if strings.Compare(*pool.Name, *name) == 0 {
				pp.Println("👍 The device pool has found and gathering ARN 👍")
				return *pool.Arn, nil
			}
		}
	}
	return "", errors.New("cant find any match for device pool")
}

func (a DeviceFarmClient) GetTestSpecType(fileArn string) (string, error) {
	upload, err := a.DeviceFarm.GetUpload(context.TODO(), &devicefarm.GetUploadInput{
		Arn: aws.String(fileArn),
	})
	if err != nil {
		return "", err
	}

	convertTypeToRunType := map[string]string{
		"APPIUM_NODE_TEST_PACKAGE":        "APPIUM_NODE",
		"APPIUM_RUBY_TEST_PACKAGE":        "APPIUM_RUBY",
		"APPIUM_PYTHON_TEST_PACKAGE":      "APPIUM_PYTHON",
		"APPIUM_JAVA_JUNIT_TEST_PACKAGE":  "APPIUM_JAVA_JUNIT",
		"APPIUM_JAVA_TESTNG_TEST_PACKAGE": "APPIUM_JAVA_TESTING",
		"UIAUTOMATION_TEST_PACKAGE":       "UIAUTOMATION",
		"UIAUTOMATOR_TEST_PACKAGE":        "UIAUTOMATOR",
		"XCTEST_TEST_PACKAGE":             "XCTEST",
		"XCTEST_UI_TEST_PACKAGE":          "XCTEST_UI",
	}

	pp.Println("👍 The test spec file's type has gathered successfully  👍", convertTypeToRunType[string(upload.Upload.Type)])

	if convertTypeToRunType[string(upload.Upload.Type)] == "" {
		return "", errors.New(string("Incompatible type for schedule " + upload.Upload.Type))
	}

	return convertTypeToRunType[string(upload.Upload.Type)], nil
}

func (a DeviceFarmClient) UploadFile(projectArn, fileType, path string) (string, error) {
	baseFile := filepath.Base(path)
	input := devicefarm.CreateUploadInput{
		Name:       aws.String(baseFile),
		ProjectArn: aws.String(projectArn),
		Type:       types.UploadType(fileType),
	}
	upload, err := a.DeviceFarm.CreateUpload(context.TODO(), &input)

	if err != nil {
		return "", errors.New("cant upload file: " + path + err.Error())
	}

	err = uploadFileToAWS(*upload.Upload.Url, path)

	if err != nil {
		return "", errors.New("cant upload file: " + path + err.Error())
	}

	for _ = range time.Tick(time.Second * 5) {
		uploadStatus, err := a.DeviceFarm.GetUpload(context.TODO(), &devicefarm.GetUploadInput{
			Arn: upload.Upload.Arn,
		})

		if err != nil {
			return "", err
		}

		if uploadStatus.Upload.Status == SUCCEEDED {
			pp.Println("👍 The file has uploaded successfully and gathering ARN 👍", path)
			return *upload.Upload.Arn, nil
		} else if uploadStatus.Upload.Status == "FAILED" {
			return "", errors.New("a problem occurred network problem or file and its type doesn't match")
		}
	}
	return "", err
}

func uploadFileToAWS(url, filePath string) error {
	data, err := os.Open(filePath)
	pp.Println("👍 Starting to read file from disk 👍", filePath)

	if err != nil {
		return errors.New("error occurred when read file from disk 😡 " + filePath)
	}
	defer data.Close()

	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(string(content)))
	if err != nil {
		return errors.New("error occurred when read file from disk 😡 " + err.Error())
	}

	req.Header.Del("Transfer-Encoding")

	client := &http.Client{}

	pp.Println("👍 Starting to upload file to aws 😎 ", filePath)

	res, err := client.Do(req)
	if err != nil {
		return errors.New("error occurred when sending request to server 😡 " + err.Error())
	}

	defer res.Body.Close()

	return nil
}

func (a DeviceFarmClient) StartRun(args devicefarm.ScheduleRunInput) error {
	pp.Println("👍 Starting a run for provided app on device farm 👍")
	run, err := a.DeviceFarm.ScheduleRun(context.TODO(), &args)
	if err != nil {
		return err
	}
	pp.Println("👍 Here is the link to follow on Aws Device Farm 👍 ")
	fmt.Println(GetAwsLinkOfRun(run.Run.Arn))
	for _ = range time.Tick(time.Second * 5) {
		getRunResult, err2 := a.DeviceFarm.GetRun(context.TODO(), &devicefarm.GetRunInput{Arn: run.Run.Arn})
		if err2 != nil {
			return err2
		}
		switch getRunResult.Run.Status {
		case types.ExecutionStatusCompleted:
			pp.Println("Your test run has done, Device Farm:", getRunResult.Run.Status)

			pp.Println("Your test job result is, Device Farm:", getRunResult.Run.Result)

			artifacts, _ := a.DeviceFarm.ListArtifacts(context.TODO(), &devicefarm.ListArtifactsInput{
				Arn:  run.Run.Arn,
				Type: "FILE",
			})

			pp.Println(artifacts.Artifacts)
			pp.Printf("Device Farm Total Time: %v | Total jobs: %v | Passed Jobs: %v | Failed Jobs: %v",
				*getRunResult.Run.DeviceMinutes.Total,
				*getRunResult.Run.Counters.Total,
				*getRunResult.Run.Counters.Passed,
				*getRunResult.Run.Counters.Failed)

			if getRunResult.Run.Result != "PASSED" {
				return errors.New("tests are failed")
			}

			return nil
		case types.ExecutionStatusRunning:
			line, _ := a.ListJobsInOneLine(*getRunResult.Run.Arn)

			pp.Printf("\033[2K\r %s %s %s %s",
				"Your test run is running, Device Farm:", getRunResult.Run.Status,
				time.Now().Format("01-02-2006 15:04:05 Monday"),
				line)

		case types.ExecutionStatusScheduling:
			line, _ := a.ListJobsInOneLine(*getRunResult.Run.Arn)

			pp.Printf("\033[2K\r %s %s %s %s",
				"Your test run added queue, Device Farm:",
				getRunResult.Run.Status,
				time.Now().Format("01-02-2006 15:04:05 Monday"),
				line)
		}
	}
	return nil
}

func GetAwsLinkOfRun(runArn *string) string {
	formattedString := strings.ReplaceAll(strings.Split(*runArn, ":run:")[1], "/", "/runs/")
	link := "https://us-west-2.console.aws.amazon.com/devicefarm/home?region=eu-west-2#/projects/" + formattedString
	return link
}

func (a DeviceFarmClient) ListJobsInOneLine(runArn string) (string, error) {
	jobs, err := a.DeviceFarm.ListJobs(context.TODO(), &devicefarm.ListJobsInput{
		Arn: aws.String(runArn),
	})
	if err != nil {
		return "", err
	}
	oneLine := ""
	for _, job := range jobs.Jobs {
		oneLine += *job.Name + " " + string(job.Result) + " " + string(job.Status) + "\n"
	}
	return oneLine, nil
}
