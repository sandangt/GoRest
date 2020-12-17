package repository_test

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
	"GoRest/repository"
	"GoRest/entity"
)

var _ = Describe("User repository", func() {

	Describe("User not found", func() {
		Context("When using non-existed user's identityID", func() {
			It("Should return empty entity.User struct", func() {
				By("Using identityID not_exist", func() {
					Expect(repository.ReadUserByIdentityID("not_exist")).To(Equal(entity.User{}))
				})
				By("Using identityID testIdentityID", func() {
					Expect(repository.ReadUserByIdentityID("testIdentityID")).To(Equal(entity.User{}))
				})
			})
		})
	})

	Describe("User exists", func() {
		Context("When using valid user's identityID", func() {
			It("Should not return empty entity.User struct", func() {
				By("Using identityID amadmin", func() {
					Expect(repository.ReadUserByIdentityID("amadmin")).ShouldNot(Equal(entity.User{}))
				})
				By("Using other valid identityID", func() {
					Expect(repository.ReadUserByIdentityID("fooagencygroupuser@fooagencygroup.com")).ShouldNot(Equal(entity.User{}))
				})
			})
		})
	})
})
