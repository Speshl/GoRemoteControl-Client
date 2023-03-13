package models

type GroundState struct {
	State
	Steer     int
	Gas       int
	Brake     int
	Clutch    int
	HandBrake int
	Pan       int
	Tilt      int
	Gear      int
	NumGears  int
	Aux       [8]bool
}

func (s GroundState) GetType() ControlSchema {
	return s.Schema
}

// Only used by server
func (s GroundState) GetBytes() []byte {
	return nil
}
