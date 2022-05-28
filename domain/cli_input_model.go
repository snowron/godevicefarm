package domain

import (
	"errors"
	"path"
)

type CliInput struct {
	ProjectName               *string
	DevicePoolName            *string
	AppPath                   *string
	TestSpecConfigurationType *string
	TestSpecConfigurationPath *string
	TestSpecConfigurationName *string
	TestSpecPath              *string
	TestSpecName              *string
	TestSpecType              *string
	AppName                   *string
	AppType                   *string
	TestName                  *string
}

func (c *CliInput) Analysis() error {

	if *c.TestSpecName == "" {
		if *c.TestSpecPath == "" && *c.TestSpecType == "" {
			return errors.New("-testSpecPath or -testSpecName must be exist")
		}
	}

	if *c.TestSpecConfigurationName == "" {
		if *c.TestSpecConfigurationPath == "" && *c.TestSpecConfigurationType == "" {
			return errors.New("-testSpecConfigurationName or -testSpecConfigurationName must be exist")
		}
	}

	if *c.TestSpecPath != "" {
		exist := c.IsTestTypeExist()
		if !exist {
			return errors.New("not exist test spec type")
		}
	}

	if *c.TestSpecConfigurationPath != "" {
		exist := c.IsTestConfigurationTypeExist()
		if !exist {
			return errors.New("not exist test spec configuration type")
		}
	}

	if *c.AppName == "" && *c.AppPath == "" {
		return errors.New("-appName or -appPath must be exist")
	}

	if *c.AppPath != "" {
		fileExtension := path.Ext(*c.AppPath)
		switch fileExtension {
		case ".apk":
			*c.AppType = "ANDROID_APP"
		case ".ipa":
			*c.AppType = "IOS_APP"
		default:
			return errors.New("app file does not have .ipa or .apk extension " + *c.AppPath)
		}
	}
	return nil
}

func (c *CliInput) IsTestTypeExist() bool {
	fileTypesOnDeviceFarm := []string{"ANDROID_APP",
		"IOS_APP",
		"WEB_APP",
		"EXTERNAL_DATA",
		"APPIUM_JAVA_JUNIT_TEST_PACKAGE",
		"APPIUM_JAVA_TESTNG_TEST_PACKAGE",
		"APPIUM_PYTHON_TEST_PACKAGE",
		"APPIUM_NODE_TEST_PACKAGE",
		"APPIUM_RUBY_TEST_PACKAGE",
		"APPIUM_WEB_JAVA_JUNIT_TEST_PACKAGE",
		"APPIUM_WEB_JAVA_TESTNG_TEST_PACKAGE",
		"APPIUM_WEB_PYTHON_TEST_PACKAGE",
		"APPIUM_WEB_NODE_TEST_PACKAGE",
		"APPIUM_WEB_RUBY_TEST_PACKAGE",
		"CALABASH_TEST_PACKAGE",
		"INSTRUMENTATION_TEST_PACKAGE",
		"UIAUTOMATION_TEST_PACKAGE",
		"UIAUTOMATOR_TEST_PACKAGE",
		"XCTEST_TEST_PACKAGE",
		"XCTEST_UI_TEST_PACKAGE",
		"APPIUM_JAVA_JUNIT_TEST_SPEC",
		"APPIUM_JAVA_TESTNG_TEST_SPEC",
		"APPIUM_PYTHON_TEST_SPEC",
		"APPIUM_NODE_TEST_SPEC",
		"APPIUM_RUBY_TEST_SPEC",
		"APPIUM_WEB_JAVA_JUNIT_TEST_SPEC",
		"APPIUM_WEB_JAVA_TESTNG_TEST_SPEC",
		"APPIUM_WEB_PYTHON_TEST_SPEC",
		"APPIUM_WEB_NODE_TEST_SPEC",
		"APPIUM_WEB_RUBY_TEST_SPEC",
		"INSTRUMENTATION_TEST_SPEC",
		"XCTEST_UI_TEST_SPEC"}

	for _, testType := range fileTypesOnDeviceFarm {
		if *c.TestSpecType == testType {
			return true
		}
	}

	return false
}

func (c *CliInput) IsTestConfigurationTypeExist() bool {
	fileTypesOnDeviceFarm := []string{"ANDROID_APP",
		"IOS_APP",
		"WEB_APP",
		"EXTERNAL_DATA",
		"APPIUM_JAVA_JUNIT_TEST_PACKAGE",
		"APPIUM_JAVA_TESTNG_TEST_PACKAGE",
		"APPIUM_PYTHON_TEST_PACKAGE",
		"APPIUM_NODE_TEST_PACKAGE",
		"APPIUM_RUBY_TEST_PACKAGE",
		"APPIUM_WEB_JAVA_JUNIT_TEST_PACKAGE",
		"APPIUM_WEB_JAVA_TESTNG_TEST_PACKAGE",
		"APPIUM_WEB_PYTHON_TEST_PACKAGE",
		"APPIUM_WEB_NODE_TEST_PACKAGE",
		"APPIUM_WEB_RUBY_TEST_PACKAGE",
		"CALABASH_TEST_PACKAGE",
		"INSTRUMENTATION_TEST_PACKAGE",
		"UIAUTOMATION_TEST_PACKAGE",
		"UIAUTOMATOR_TEST_PACKAGE",
		"XCTEST_TEST_PACKAGE",
		"XCTEST_UI_TEST_PACKAGE",
		"APPIUM_JAVA_JUNIT_TEST_SPEC",
		"APPIUM_JAVA_TESTNG_TEST_SPEC",
		"APPIUM_PYTHON_TEST_SPEC",
		"APPIUM_NODE_TEST_SPEC",
		"APPIUM_RUBY_TEST_SPEC",
		"APPIUM_WEB_JAVA_JUNIT_TEST_SPEC",
		"APPIUM_WEB_JAVA_TESTNG_TEST_SPEC",
		"APPIUM_WEB_PYTHON_TEST_SPEC",
		"APPIUM_WEB_NODE_TEST_SPEC",
		"APPIUM_WEB_RUBY_TEST_SPEC",
		"INSTRUMENTATION_TEST_SPEC",
		"XCTEST_UI_TEST_SPEC"}

	for _, testType := range fileTypesOnDeviceFarm {
		if *c.TestSpecConfigurationType == testType {
			return true
		}
	}

	return false
}
