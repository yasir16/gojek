package store

import (
	"parking_lot/errors"
	"parking_lot/schema"
	"parking_lot/utils"
	"strconv"
	"strings"
)

type getSlotNumberByRegisNumber struct {
	*store
}

func NewGetSlotNumberByRegisNumber(store *store) *getSlotNumberByRegisNumber {
	return &getSlotNumberByRegisNumber{store}
}

func (s *getSlotNumberByRegisNumber) IsHelp(arg string) (string, bool) {
	if arg == string(schema.CMDHelp) {
		return schema.CMDSlotNumberWithRegistNumHint, true
	}
	return "", false
}

func (s *getSlotNumberByRegisNumber) Execute(cmd *schema.Command) (string, error) {
	if res, isHelp := s.IsHelp(cmd.Arguments[0]); isHelp {
		return res, nil
	}
	if err := validateRegisNo(cmd.Arguments); err != nil {
		return "", err
	}
	if ParkingLot == nil {
		return "", errors.ErrNoParkingLot
	}
	var slotNumber string
	slot := ParkingLot.GetSlotIDByRegisNum(strings.ToUpper(cmd.Arguments[0]))
	if slot != nil {
		slotNumber = strconv.FormatUint(uint64(slot.GetID()), 10)
	} else {
		slotNumber = "Not found"
	}
	return slotNumber, nil
}
func validateRegisNo(args []string) error {
	if !utils.IsRegNoValid(args[0]) {
		return errors.ErrInvalidRegNo
	}
	return nil
}
