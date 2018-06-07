package idrac

import (
	"fmt"

	"github.com/levigross/grequests"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

// Allowed Power Actions [On ForceOff GracefulRestart PushPowerButton Nmi]

// Host is
type Host struct {
	URL     string
	Options *grequests.RequestOptions
}

// NewHost is
func NewHost(ip, user, pass string) *Host {
	return &Host{
		URL:     fmt.Sprintf("https://%s/redfish/v1/Systems/System.Embedded.1/", ip),
		Options: &grequests.RequestOptions{Auth: []string{user, pass}, InsecureSkipVerify: true},
	}
}

// GetCurrentPowerState is
func (h Host) GetCurrentPowerState() (string, error) {
	queryPath := "PowerState"

	resp, err := grequests.Get(h.URL, h.Options)
	if err != nil {
		log.Error(err)
		return "", err
	}
	ret := gjson.Get(resp.String(), queryPath)

	return ret.String(), nil
}

// GetAllowablePowerActions ["On", "ForceOff", "GracefulRestart", "PushPowerButton", "Nmi"]
func (h Host) GetAllowablePowerActions() (actions []string, err error) {
	queryPath := "Actions.#ComputerSystem\\.Reset.ResetType@Redfish\\.AllowableValues"

	resp, err := grequests.Get(h.URL, h.Options)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	ret := gjson.Get(resp.String(), queryPath)

	for _, r := range ret.Array() {
		actions = append(actions, r.Str)
	}
	return
}

// SetPowerState is
func (h Host) SetPowerState(powerState string) (err error) {
	url := fmt.Sprintf("%sActions/ComputerSystem.Reset", h.URL)

	h.Options.Headers = map[string]string{"content-type": "application/json"}
	h.Options.JSON = map[string]string{"ResetType": powerState}

	resp, err := grequests.Post(url, h.Options)
	if err != nil {
		log.Error(err)
		return err
	}

	if resp.StatusCode == 204 {
		log.Infof("Power action pass, host power state change to %s", powerState)
		return nil
	}
	log.Errorf("Power action Failed, return code is %d", resp.StatusCode)
	return fmt.Errorf("%s", resp.String())
}
