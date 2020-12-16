package repository_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	
	"GoRest/configuration"
	"GoRest/repository"
)

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}


var _ = BeforeSuite(func() {
	config := configuration.ParseConfiguration()
	repository.CreateDriver(config)
})

var _ = AfterSuite(func() {
    repository.CloseDriver()
})
