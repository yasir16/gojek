package store

import (
	"parking_lot/errors"
	"parking_lot/schema"
	"parking_lot/utils"
	"strconv"
	"strings"
)

type getSlotNumberByColorStore struct {
	*store
}

func NewGetSlotNumberByColorStore(store *store) *getSlotNumberByColorStore {
	return &getSlotNumberByColorStore{store}
}

func (s *getSlotNumberByColorStore) IsHelp(arg string) (string, bool) {
	if arg == string(schema.CMDHelp) {
		return schema.CMDSlotNumberWithColorHint, true
	}
	return "", false
}

func (s *getSlotNumberByColorStore) Execute(cmd *schema.Command) (string, error) {
	if res, isHelp := s.IsHelp(cmd.Arguments[0]); isHelp {
		return res, nil
	}
	if err := validateColor(cmd.Arguments); err != nil {
		return "", err
	}
	if ParkingLot == nil {
		return "", errors.ErrNoParkingLot
	}
	var slotNumbers []string
	slots := ParkingLot.GetSlotNumbersByColor(strings.ToLower(cmd.Arguments[0]))
	for _, slot := range slots {
		slotNumbers = append(slotNumbers, strconv.FormatUint(uint64(slot.GetID()), 10))
	}
	return strings.Join(slotNumbers, ", "), nil
}

func validateColor(args []string) error {
	if !utils.IsValidString(args[0]) {
		return errors.ErrInvalidColour
	}
	return nil
}
