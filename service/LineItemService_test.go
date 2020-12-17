package service_test

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
	"GoRest/entity"
	"GoRest/service"
)

var _ = Describe("Line Item service", func() {
	
	Describe("Return user not found message", func() {
		Context("When using non-existed user's identityID", func() {
			It("Should return entity.Message", func() {
				By("Using identityID not_exist", func() {
					message := service.ReadLineItemsByUserIdentityID("not_exist").(entity.Message)
					Expect(message.Content).To(Equal("User not found"))
				})
			})
		})
	})
	
	
	Describe("Return list of entity.LineItem", func() {
		Context("When using identityID amadmin", func() {
			It("Should not return empty []entity.LineItem list", func() {
				Expect(len( service.ReadLineItemsByUserIdentityID("amadmin").([]entity.LineItem) )).ShouldNot(Equal(0))
			})
			It("Should have IdentityID=amadmin", func() {
				Expect(service.ReadLineItemsByUserIdentityID("amadmin").([]entity.LineItem)[0].IdentityID).To(Equal("amadmin"))
			})
		})
		Context("When using other valid identityID", func() {
			It("Should not return empty []entity.LineItem list", func() {
				Expect(len(service.ReadLineItemsByUserIdentityID("fooagencyadmin@fooagency.com").([]entity.LineItem) )).ShouldNot(Equal(0))
			})
			It("Should have IdentityID = this identityID", func() {
				By("IdentityID = this identity", func() {
					Expect(service.ReadLineItemsByUserIdentityID("fooagencyadmin@fooagency.com").([]entity.LineItem)[0].IdentityID).To(Equal("fooagencyadmin@fooagency.com"))
				})
				By("IdentityID != different identityID", func() {
					Expect(service.ReadLineItemsByUserIdentityID("fooagencyadmin@fooagency.com").([]entity.LineItem)[0].IdentityID).ShouldNot(Equal("amadmin"))
				})
			})
		})
	})

	
})
