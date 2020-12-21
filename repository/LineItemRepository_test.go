package repository_test

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
	"GoRest/repository"
)

var _ = Describe("Line Item repository", func() {
	// Init list of arrays as empty string map
	testArrayParams := make(map[string]string)
	
	Describe("No optional params", func() {
		
		Describe("Empty line item list", func() {
			Context("When using non-existed user's identityID", func() {
				It("Should return empty []entity.LineItem list", func() {
					By("Using identityID not_exist", func() {
						Expect(len(repository.ReadLineItemsByUserIdentityID("not_exist", testArrayParams))).To(Equal(0))
					})
					By("Using identityID testIdentityID", func() {
						Expect(len(repository.ReadLineItemsByUserIdentityID("testIdentityID", testArrayParams))).To(Equal(0))
					})
				})
			})
		})
		
		Describe("Line item list contains number of elements", func() {
			Context("When using identityID amadmin", func() {
				It("Should not return empty []entity.LineItem list", func() {
					Expect(len(repository.ReadLineItemsByUserIdentityID("amadmin", testArrayParams))).ShouldNot(Equal(0))
				})
				It("Should have IdentityID=amadmin", func() {
					Expect(repository.ReadLineItemsByUserIdentityID("amadmin", testArrayParams)[0].IdentityID).To(Equal("amadmin"))
				})
			})
			Context("When using other valid identityID", func() {
				It("Should not return empty []entity.LineItem list", func() {
					Expect(len(repository.ReadLineItemsByUserIdentityID("fooagencyadmin@fooagency.com",testArrayParams))).ShouldNot(Equal(0))
				})
				It("Should have IdentityID = this identityID", func() {
					By("IdentityID = this identity", func() {
						Expect(repository.ReadLineItemsByUserIdentityID("fooagencyadmin@fooagency.com", testArrayParams)[0].IdentityID).To(Equal("fooagencyadmin@fooagency.com"))
					})
					By("IdentityID != different identityID", func() {
						Expect(repository.ReadLineItemsByUserIdentityID("fooagencyadmin@fooagency.com", testArrayParams)[0].IdentityID).ShouldNot(Equal("amadmin"))
					})
				})
			})
		})
	})
	
	Describe("External optional params", func() {
		Describe("Empty line item list", func() {
			Context("When using valid identityID and 1 invalid external parameter", func() {
				It("Should return empty []entity.LineItem list", func() {
					By("Using brandName not_exist", func() {
						testArrayParams["brandName"] = "not_exist"
						Expect(len(repository.ReadLineItemsByUserIdentityID("amadmin", testArrayParams))).To(Equal(0))
					})
					By("Using initiativeName not_exist", func() {
						testArrayParams["initiativeName"] = "not_exist"
						Expect(len(repository.ReadLineItemsByUserIdentityID("amadmin", testArrayParams))).To(Equal(0))
					})
					By("Using creatorCompanyName not_exist", func() {
						testArrayParams["initiativeName"] = "not_exist"
						Expect(len(repository.ReadLineItemsByUserIdentityID("amadmin", testArrayParams))).To(Equal(0))
					})
				})
			})
			Context("When using valid identityID and 1 valid external parameter", func() {
				It("Should return empty []entity.LineItem list", func() {
					By("Using brandName not_exist", func() {
						testArrayParams["brandName"] = "not_exist"
						Expect(len(repository.ReadLineItemsByUserIdentityID("amadmin", testArrayParams))).To(Equal(0))
					})
					By("Using initiativeName not_exist", func() {
						testArrayParams["initiativeName"] = "not_exist"
						Expect(len(repository.ReadLineItemsByUserIdentityID("amadmin", testArrayParams))).To(Equal(0))
					})
					By("Using creatorCompanyName not_exist", func() {
						testArrayParams["initiativeName"] = "not_exist"
						Expect(len(repository.ReadLineItemsByUserIdentityID("amadmin", testArrayParams))).To(Equal(0))
					})
				})
			})
			Context("When using valid identityID and 2 or more invalid external parameters", func() {
				It("Should return empty []entity.LineItem list", func() {
					By("Using brandName not_exist, initiativeName not_exist, creatorCompanyName not_exist", func() {
						testArrayParams["brandName"] = "not_exist"
						testArrayParams["initiativeName"] = "not_exist"
						testArrayParams["initiativeName"] = "not_exist"
						Expect(len(repository.ReadLineItemsByUserIdentityID("amadmin", testArrayParams))).To(Equal(0))
					})
				})
			})
		})
		
	})

})
