package Services_test

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
	"GoNeo/Entities"
	"GoNeo/Services"
)

var _ = Describe("Line Item service", func() {
	
	Describe("Return user not found message", func() {
		Context("When using non-existed user's identityID", func() {
			It("Should return Entities.Message", func() {
				By("Using identityID not_exist", func() {
					message := Services.ReadLineItemsByUserIdentityID("not_exist").(Entities.Message)
					Expect(message.Content).To(Equal("User not found"))
				})
			})
		})
	})
	
	
	Describe("Return list of Entities.LineItem", func() {
		Context("When using identityID amadmin", func() {
			It("Should not return empty []Entities.LineItem list", func() {
				Expect(len( Services.ReadLineItemsByUserIdentityID("amadmin").([]Entities.LineItem) )).ShouldNot(Equal(0))
			})
			It("Should have IdentityID=amadmin", func() {
				Expect(Services.ReadLineItemsByUserIdentityID("amadmin").([]Entities.LineItem)[0].IdentityID).To(Equal("amadmin"))
			})
		})
		Context("When using other valid identityID", func() {
			It("Should not return empty []Entities.LineItem list", func() {
				Expect(len(Services.ReadLineItemsByUserIdentityID("fooagencyadmin@fooagency.com").([]Entities.LineItem) )).ShouldNot(Equal(0))
			})
			It("Should have IdentityID = this identityID", func() {
				By("IdentityID = this identity", func() {
					Expect(Services.ReadLineItemsByUserIdentityID("fooagencyadmin@fooagency.com").([]Entities.LineItem)[0].IdentityID).To(Equal("fooagencyadmin@fooagency.com"))
				})
				By("IdentityID != different identityID", func() {
					Expect(Services.ReadLineItemsByUserIdentityID("fooagencyadmin@fooagency.com").([]Entities.LineItem)[0].IdentityID).ShouldNot(Equal("amadmin"))
				})
			})
		})
	})

	
})
