package main

import (
	"context"
	"flag"
	"godevicefarm/client"
	"godevicefarm/domain"
	"godevicefarm/services"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
	"github.com/k0kubun/pp"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))

	if err != nil {
		pp.Fatalf("unable to load SDK config, %v", err.Error())
	}

	typeFlag := flag.NewFlagSet("mobile", flag.PanicOnError)
	inputStruct := domain.CliInput{
		ProjectName:               typeFlag.String("projectName", "", "projectName"),
		DevicePoolName:            typeFlag.String("devicePoolName", "", "devicePoolName"),
		TestSpecType:              typeFlag.String("testSpecType", "", "testSpecType"),
		TestSpecName:              typeFlag.String("testSpecName", "", "testSpecName"),
		TestSpecPath:              typeFlag.String("testSpecPath", "", "testSpecPath"),
		TestSpecConfigurationName: typeFlag.String("testSpecConfigurationName", "", "testSpecConfigurationName"),
		TestSpecConfigurationType: typeFlag.String("testSpecConfigurationType", "", "testSpecConfigurationType"),
		TestSpecConfigurationPath: typeFlag.String("testSpecConfigurationPath", "", "testSpecConfigurationPath"),
		AppPath:                   typeFlag.String("appPath", "", "appPath"),
		AppName:                   typeFlag.String("appName", "", "appName"),
		AppType:                   aws.String(""),
		TestName:                  typeFlag.String("testName", "Test From CLI", "testName"),
	}

	if len(os.Args) > 1 && os.Args[1] == "mobile" {
		err = typeFlag.Parse(os.Args[2:])
		if err != nil {
			pp.Fatal(err.Error())
		}
	} else {
		pp.Fatal("specify test selection ./godevicefarm mobile -appName=xx")
	}

	err = inputStruct.Analysis()

	if err != nil {
		pp.Fatal(err.Error())
	}
	farmClient := client.DeviceFarmClient{DeviceFarm: devicefarm.NewFromConfig(cfg)}
	cliClient := services.CliService{DeviceFarm: farmClient}

	err = cliClient.StartTestingProcess(os.Args[1], &inputStruct)
	if err != nil {
		pp.Fatal(err.Error())
	}
}
