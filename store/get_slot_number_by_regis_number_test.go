package store

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"parking_lot/errors"
	"parking_lot/schema"
)

var _ = Describe("get slot number by regis number parking store test", func() {
	var (
		connection Store
	)
	connection = NewStore()
	It("Tear Down Store Data", func() {
		TearDown()
	})
	Context("get slot store execute", func() {
		TearDown()
		It("slot_number_for_registration_number help", func() {
			cmd := &schema.Command{
				Command:   "slot_number_for_registration_number",
				Arguments: []string{"help"},
			}
			res, err := connection.GetSlotNubByRegisNo().Execute(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal(schema.CMDSlotNumberWithRegistNumHint))
		})
		It("No parking lot available", func() {
			cmd := &schema.Command{
				Command:   "slot_number_for_registration_number",
				Arguments: []string{"TN-24-AJ-8462"},
			}
			res, err := connection.GetSlotNubByRegisNo().Execute(cmd)

			Expect(err).To(Equal(errors.ErrNoParkingLot))
			Expect(res).To(Equal(""))
		})
		It("Create a parking lot with 5 slots", func() {
			cmd := &schema.Command{
				Command:   "create_parking_lot",
				Arguments: []string{"1"},
			}
			res, err := connection.CreateParkingLot().Execute(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal(fmt.Sprintf(ParkinglotCreatedInfo, 1)))
		})
		It("park a vehicle", func() {
			cmd := &schema.Command{
				Command:   "park",
				Arguments: []string{"TN-24-AJ-8462", "Red"},
			}
			res, err := connection.Park().Execute(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal("Allocated slot number: 1"))
		})
		It("get slot num by regis num success ", func() {
			cmd := &schema.Command{
				Command:   "slot_number_for_registration_number",
				Arguments: []string{"TN-24-AJ-8462"},
			}
			res, err := connection.GetSlotNubByRegisNo().Execute(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal("1"))
		})
		It("get slot num by regis num not success ", func() {
			cmd := &schema.Command{
				Command:   "slot_number_for_registration_number",
				Arguments: []string{"TN-24-AJ-8461"},
			}
			res, err := connection.GetSlotNubByRegisNo().Execute(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal("Not found"))
		})
	})
})
