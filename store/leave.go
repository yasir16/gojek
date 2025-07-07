package store

import (
	"fmt"
	"parking_lot/errors"
	"parking_lot/schema"
	"strconv"
)

type leaveStore struct {
	*store
}

func NewLeaveStore(st *store) *leaveStore {
	ll := &leaveStore{st}
	return ll
}

func (ll *leaveStore) IsHelp(arg string) (string, bool) {
	if arg == string(schema.CMDHelp) {
		return schema.CMDLeaveHint, true
	}
	return "", false
}

func (ll *leaveStore) Execute(cmd *schema.Command) (string, error) {
	if res, isHelp := ll.IsHelp(cmd.Arguments[0]); isHelp {
		return res, nil
	}
	slotID, err := strconv.Atoi(cmd.Arguments[0])
	if err != nil {
		return "", errors.ErrInvalidInputSlot
	}
	if ParkingLot == nil {
		return "", errors.ErrNoParkingLot
	}
	if err = ParkingLot.LeaveSlot(uint(slotID)); err != nil {

	}
	return fmt.Sprintf(SlotIsFreeInfo, slotID), nil
}
