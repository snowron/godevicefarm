
# AWS Device Farm CLI Tool

Basically it's a cli tool which helps to automatize our e2e mobile tests workloads on CI pipelines.

[How to run mobile test step-by-step with aws-cli](https://docs.aws.amazon.com/devicefarm/latest/developerguide/how-to-create-test-run.html#how-to-create-test-run-cli)

This project makes these steps simple.




## Prerequisites

A existed mobile project service and configured device pool on device farm.


## Environment Variables

To run this project, you will need to export the following environment variables to your terminal session.

`AWS_ACCESS_KEY_ID`

`AWS_SECRET_ACCESS_KEY`

`AWS_DEFAULT_REGION` is already set in project source code. Device Farm servivces are only available in us-west-2 region.

## How to run

Clone the project

Install dependencies

```bash
  go mod download
```

Start the test process

There are two way to run tests. Device farm needs test package and related configuration file.

If already uploaded these files then get names from aws console(ui) and put remote bash executation line.
```bash
  ./godevicefarm mobile -testSpecName remoteSpecName.yml -testSpecConfigurationName testSpecConfigurationName.yml -devicePoolName devicePool -testName TheTestFromCli -appName androidApp 
```

If you didn't upload test package and configuration file then use below bash executation. Fill with right paths. It will upload these files.
```bash
  ./godevicefarm mobile -testSpecType "APPIUM_NODE_TEST_PACKAGE" -testSpecPath ./testSpecPath.zip -testSpecConfigurationType "APPIUM_NODE_TEST_SPEC" -testSpecConfigurationName ./testSpecConfigurationName.yml -devicePoolName devicePool -testName TheTestFromCli -appPath ./android.apk
```


## Running Tests

To run tests, run the following command

```bash
  go test ./...
```


## Authors

- [Murat TURAN](https://www.linkedin.com/in/muratturan0/)

