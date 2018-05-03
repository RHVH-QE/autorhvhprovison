package pre

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func beakerClientCheck() {
	log.Info("start checking beaker-client")
	_, err := exec.LookPath("bkr")
	if err != nil {
		log.Error("can not found command `bkr` in path, install beaker-client first")
		os.Exit(-1)
	}
	log.Info("found beaker client...done")

	log.Info("start checking if valid krb ticket exists")
	out, err := exec.Command("bkr", "whoami").CombinedOutput()
	if err != nil {
		log.Errorf("%s", out)
		os.Exit(-1)
	}
	log.Info("found valid krb ticket...done")
}

// Check is
func Check() {
	beakerClientCheck()
}
