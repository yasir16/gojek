package store

import (
	"parking_lot/errors"
	"parking_lot/schema"
	"parking_lot/utils"
	"strings"
)

type getStore struct {
	*store
}

func NewGetStore(st *store) *getStore {
	return &getStore{st}
}

func (gs *getStore) IsHelp(arg string) (string, bool) {
	if arg == string(schema.CMDHelp) {
		return schema.CMDRegisNumForCarsWithColorHint, true
	}
	return "", false
}

func (gs *getStore) Execute(cmd *schema.Command) (string, error) {
	if res, isHelp := gs.IsHelp(cmd.Arguments[0]); isHelp {
		return res, nil
	}
	if err := validateGetReq(cmd.Arguments); err != nil {
		return "", err
	}
	if ParkingLot == nil {
		return "", errors.ErrNoParkingLot
	}
	var regisNumbers []string
	slots := ParkingLot.GetRegisNumberByColor(strings.ToLower(cmd.Arguments[0]))
	for _, slot := range slots {
		regisNumbers = append(regisNumbers, slot.Vehicle.GetRegNumber())
	}

	return strings.Join(regisNumbers, ", "), nil
}

func validateGetReq(args []string) error {
	if !utils.IsValidString(args[0]) {
		return errors.ErrInvalidColour
	}
	return nil
}
