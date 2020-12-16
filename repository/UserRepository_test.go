package Repositories_test

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
	"GoNeo/Repositories"
	"GoNeo/Entities"
)

var _ = Describe("User repository", func() {

	Describe("User not found", func() {
		Context("When using non-existed user's identityID", func() {
			It("Should return empty Entities.User struct", func() {
				By("Using identityID not_exist", func() {
					Expect(Repositories.ReadUserByIdentityID("not_exist")).To(Equal(Entities.User{}))
				})
				By("Using identityID testIdentityID", func() {
					Expect(Repositories.ReadUserByIdentityID("testIdentityID")).To(Equal(Entities.User{}))
				})
			})
		})
	})

	Describe("User exists", func() {
		Context("When using valid user's identityID", func() {
			It("Should not return empty Entities.User struct", func() {
				By("Using identityID amadmin", func() {
					Expect(Repositories.ReadUserByIdentityID("amadmin")).ShouldNot(Equal(Entities.User{}))
				})
				By("Using other valid identityID", func() {
					Expect(Repositories.ReadUserByIdentityID("fooagencygroupuser@fooagencygroup.com")).ShouldNot(Equal(Entities.User{}))
				})
			})
		})
	})
})
