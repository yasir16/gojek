package schema

import (
	"parking_lot/errors"
	"time"
)

// ParkingLot struct holds all the parking lot information and parking history
type ParkingLot struct {
	Name        string         `json:"name"`
	Floor       string         `json:"floor"`
	TotalBlocks int            `json:"total_blocks"`
	BlockHeight int            `json:"block_height"`
	TotalSlots  int            `json:"total_slots"`
	Address     string         `json:"address"`
	Pincode     string         `json:"pincode"`
	Slots       []*Slot        `json:"slots"`
	ParkHistory []*ParkHistory `json:"park_history"`
}

// ParkHistory holds the parking information
type ParkHistory struct {
	SlotID             uint
	RegistrationNumber string
	Colour             string
	CreatedAt          time.Time
}

// FirstAvailableSlot returns the first available slot to park Vehicle
func (pl *ParkingLot) FirstAvailableSlot() (*Slot, error) {
	for _, slot := range pl.Slots {
		if slot.IsSlotAvailable() {
			return slot, nil
		}
	}

	return nil, errors.ErrParkingSlotsFull
}

func (pl *ParkingLot) LeaveSlot(slotID uint) error {
	for _, slot := range pl.Slots {
		if slot.ID == slotID {
			err := slot.ExitPark()
			if err != nil {
				return err
			}
			break
		}
	}
	return nil
}

func (pl *ParkingLot) GetRegisNumberByColor(color string) []*Slot {
	var newSlots []*Slot
	for _, slot := range pl.Slots {
		if slot.Vehicle.GetColour() == color {
			newSlots = append(newSlots, slot)
		}
	}
	return newSlots
}
func (pl *ParkingLot) GetSlotNumByColor(color string) []*Slot {
	var newSlots []*Slot
	for _, slot := range pl.Slots {
		if slot.Vehicle.GetColour() == color {
			newSlots = append(newSlots, slot)
		}
	}
	return newSlots
}

func (pl *ParkingLot) GetSlotNumbersByColor(color string) []*Slot {
	var newSlots []*Slot
	for _, slot := range pl.Slots {
		if slot.Vehicle.GetColour() == color {
			newSlots = append(newSlots, slot)
		}
	}
	return newSlots
}

func (pl *ParkingLot) GetSlotIDByRegisNum(regNum string) *Slot {
	for _, slot := range pl.Slots {
		if slot.Vehicle.GetRegNumber() == regNum {
			return slot
		}
	}
	return nil
}
