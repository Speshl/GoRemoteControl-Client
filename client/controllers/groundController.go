package controllers

import (
	"log"
	"math"

	"github.com/Speshl/GoRemoteControl/models"
)

func (c *Controller) mapGroundState() models.StateIface {
	cfg := c.cfg.Config.(GroundConfig)
	returnState := models.GroundState{
		Steer:     c.getConfigEntryValue(cfg.Steer),
		Gas:       c.getConfigEntryValue(cfg.Gas),
		Brake:     c.getConfigEntryValue(cfg.Brake),
		Clutch:    c.getConfigEntryValue(cfg.Clutch),
		HandBrake: c.getConfigEntryValue(cfg.HandBrake),
		Pan:       c.getConfigEntryValue(cfg.Pan),
		Tilt:      c.getConfigEntryValue(cfg.Tilt),
	}

	returnState.NumGears = len(cfg.Gears) - 1 //Don't count reverse
	for gear, gearEntry := range cfg.Gears {
		if c.getButtonEntryValue(gearEntry) > 0 {
			if gear == len(cfg.Gears)-1 {
				returnState.Gear = -1
			} else {
				returnState.Gear = gear + 1
			}
		}
	}

	for pos, auxButton := range cfg.Aux {
		if c.getButtonEntryValue(auxButton) > 0 {
			returnState.Aux[pos] = true
		}
	}
	return returnState
}

func (c *Controller) getConfigEntryValue(entry ConfigEntry) int {
	if entry.Axis != nil {
		return c.getAxisEntryValue(*entry.Axis)
	} else if entry.Button != nil {
		return c.getButtonEntryValue(*entry.Button)
	} else {
		return 0
	}
}

func (c *Controller) getAxisEntryValue(entry AxisEntry) (value int) {
	if entry.DeviceID > len(c.internalStates) {
		log.Printf("Warning: Device not found - %+v\n", entry)
		return 0
	}

	value = c.internalStates[entry.DeviceID].AxisData[entry.AxisID]
	if entry.Inverted {
		value = value * -1
	}
	return value
}

func (c *Controller) getButtonEntryValue(entry ButtonEntry) (value int) {
	baseMin := -32768
	baseMax := 32768
	maxBitValue := uint32(math.Pow(2, float64(entry.MaxID)))

	if entry.DeviceID > len(c.internalStates) {
		log.Printf("Warning: Device not found - %+v\n", entry)
		return 0
	}

	if c.internalStates[entry.DeviceID].Buttons&maxBitValue > 0 {
		value = baseMax
	} else {
		value = 0
	}

	if entry.MinID != nil {
		minBitValue := uint32(math.Pow(2, float64(*entry.MinID)))
		if c.internalStates[entry.DeviceID].Buttons&minBitValue > 0 {
			value = baseMin
		}
	}
	return value
}
