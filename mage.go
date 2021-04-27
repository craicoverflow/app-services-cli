//+build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Complete mg.Namespace
type Prepare mg.Namespace
type Test mg.Namespace
type Build mg.Namespace

var (
	Default     = Complete.Run
	directories = []string{"./cmd/...", "./pkg/...", "./internal/..."}
	binaryname  = "rhoas"
	debug = getEnvOrDefault("DEBUG", "false")
	// build configuration
	rhoasVersion         = getEnvOrDefault("RHOAS_VERSION", "dev")
	repositoryOwner      = getEnvOrDefault("REPOSITORY_OWNER", "redhat-developer")
	repositoryName       = getEnvOrDefault("REPOSITORY_NAME", "app-services-cli")
	termsReviewEventCode = getEnvOrDefault("TERMS_REVIEW_EVENT_CODE", "onlineServices")
	termsReviewSiteCode  = getEnvOrDefault("TERMS_REVIEW_EVENT_CODE", "ocm")

	buildFlags = []string{}
)

func init() {
	//Enable Go modules
	os.Setenv("GO111MODULE", "on")

	// if debug == "true" {
	// 	buildFlags = append(buildFlags, []string{
	// 		`-gcflags "all=-N -l"`,
	// 	})
	// }
}

func (Complete) Run() {
	mg.Deps(
		Prepare.Run,
		Test.Run,
		Build.Local,
	)
}

func (Build) Run() error {
	mg.Deps(
		Test.Lint,
		Build.Local,
	)
	return nil
}

func (Prepare) Run() error {
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	return sh.Run("go", "install", "./...")
}

func (Test) Run() error {
	mg.Deps(
		Test.Lint,
		Test.Unit,
	)
	return nil
}

func (Test) Unit() error {
	testArgs := append([]string{"test"}, directories...)
	return sh.Run("go", testArgs...)
}

func (Test) Lint() error {
	lintArgs := append([]string{"run"}, directories...)
	return sh.Run("golangci-lint", lintArgs...)
}

func (Build) Local() error {
	buildArgs := []string{
		"build",
		// ...buildFlags,
		"-o",
		binaryname,
		"./cmd/" + binaryname,
	}
	return sh.Run("go", buildArgs...)
}

func getEnvOrDefault(env string, defaultV string) string {
	v, ok := os.LookupEnv(env)
	if !ok {
		return defaultV
	}
	return v
}
