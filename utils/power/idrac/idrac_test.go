package idrac

import (
	"testing"
)

var host = NewHost("dell-per730-34-idrac.lab.eng.pek2.redhat.com", "root", "calvin")

func TestHost_GetCurrentPowerState(t *testing.T) {
	_, err := host.GetCurrentPowerState()
	if err != nil {
		t.Error(err)
	}
}

func TestHost_GetAllowablePowerActions(t *testing.T) {
	_, err := host.GetAllowablePowerActions()
	if err != nil {
		t.Error(err)
	}
}
