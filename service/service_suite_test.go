package service_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	
	"GoRest/configuration"
	"GoRest/repository"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}

var _ = BeforeSuite(func() {
	config := configuration.ParseConfiguration()
	repository.CreateDriver(config)
})

var _ = AfterSuite(func() {
    repository.CloseDriver()
})
