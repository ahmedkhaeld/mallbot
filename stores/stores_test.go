package stores

import (
	"fmt"
	"testing"

	"github.com/cucumber/godog"
)

var storeName = ""

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "progress",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}
	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func aStoreCalledExists(name string) error {
	if storeName != name {
		return fmt.Errorf("store does not exist: %s", name)
	}
	return nil
}

func aValidStoreOwner() error {
	return nil
}

func iCreateTheStoreCalled(name string) error {
	storeName = name
	return nil
}

func noStoreCalledExists(name string) error {
	if storeName == name {
		return fmt.Errorf("store does exist: %s", name)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a store called "([^"]*)" exists$`, aStoreCalledExists)
	ctx.Step(`^a valid store owner$`, aValidStoreOwner)
	ctx.Step(`^I create the store called "([^"]*)"$`, iCreateTheStoreCalled)
	ctx.Step(`^no store called "([^"]*)" exists$`, noStoreCalledExists)
}
