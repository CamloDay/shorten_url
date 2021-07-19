package main

import (
	"flag"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"os"
	"shorten_url/step_definitions"
	"testing"
)

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress",
}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func TestMain(m *testing.M) {
	fmt.Println("TestMain run")

	flag.Parse()

	status := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: step_definitions.InitializeTestSuite,
		ScenarioInitializer:  step_definitions.InitializeScenario,
		Options:              &opts,
	}.Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
