package Repositories_test

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
	"GoNeo/Repositories"
)

var _ = Describe("Line Item repository", func() {
	
	Describe("Empty line item list", func() {
		Context("When using non-existed user's identityID", func() {
			It("Should return empty []Entities.LineItem list", func() {
				By("Using identityID not_exist", func() {
					Expect(len(Repositories.ReadLineItemsByUserIdentityID("not_exist"))).To(Equal(0))
				})
				By("Using identityID testIdentityID", func() {
					Expect(len(Repositories.ReadLineItemsByUserIdentityID("testIdentityID"))).To(Equal(0))
				})
			})
		})
	})
	
	Describe("Not empty line item list", func() {
		Context("When using identityID amadmin", func() {
			It("Should not return empty []Entities.LineItem list", func() {
				Expect(len(Repositories.ReadLineItemsByUserIdentityID("amadmin"))).ShouldNot(Equal(0))
			})
			It("Should have IdentityID=amadmin", func() {
				Expect(Repositories.ReadLineItemsByUserIdentityID("amadmin")[0].IdentityID).To(Equal("amadmin"))
			})
		})
		Context("When using other valid identityID", func() {
			It("Should not return empty []Entities.LineItem list", func() {
				Expect(len(Repositories.ReadLineItemsByUserIdentityID("fooagencyadmin@fooagency.com"))).ShouldNot(Equal(0))
			})
			It("Should have IdentityID = this identityID", func() {
				By("IdentityID = this identity", func() {
					Expect(Repositories.ReadLineItemsByUserIdentityID("fooagencyadmin@fooagency.com")[0].IdentityID).To(Equal("fooagencyadmin@fooagency.com"))
				})
				By("IdentityID != different identityID", func() {
					Expect(Repositories.ReadLineItemsByUserIdentityID("fooagencyadmin@fooagency.com")[0].IdentityID).ShouldNot(Equal("amadmin"))
				})
			})
		})
	})

	
})
