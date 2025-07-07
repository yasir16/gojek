package store

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"parking_lot/errors"
	"parking_lot/schema"
)

var _ = Describe("leave parking store test", func() {
	var (
		connection Store
	)
	connection = NewStore()
	It("Tear Down Store Data", func() {
		TearDown()
	})
	Context("leave store execute", func() {
		TearDown()
		It("leave help", func() {
			cmd := &schema.Command{
				Command:   "leave",
				Arguments: []string{"help"},
			}
			res, err := connection.Leave().Execute(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal(schema.CMDLeaveHint))
		})
		It("No parking lot available", func() {
			cmd := &schema.Command{
				Command:   "leave",
				Arguments: []string{"1"},
			}
			res, err := connection.Leave().Execute(cmd)

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
		It("leave success", func() {
			cmd := &schema.Command{
				Command:   "leave",
				Arguments: []string{"1"},
			}
			res, err := connection.Leave().Execute(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal("Slot number 1 is free"))
		})
	})
})
