package provision

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"

	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"

	"github.com/dracher/autorhvhprovison/utils/cobbler"
	"github.com/dracher/helpers"
)

const (
	colMachinePool     = "machine_pool"
	colCommon          = "common"
	rhvhSearchPattern  = "RHVH-4.[0-9]"
	buildSearchPattern = ""
	buildPath          = "/var/resources/crawled.rhvh4x_img"
)

type (
	// CommonParams is
	CommonParams struct {
		CobblerAPI      string   `bson:"cobber_api"`
		CobblerCred     []string `bson:"cobber_credential"`
		PostScript01    string   `bson:"post_script_01"`
		PostScript02    string   `bson:"post_script_02"`
		CoverageSnippet string   `bson:"coverage_snippet"`
	}

	// Machine is
	Machine struct {
		BkrName   string `bson:"bkr_name"`
		NicName   string `bson:"nic_name"`
		Mac       string `bson:"mac"`
		Used      bool   `bson:"used"`
		IdracUser string `bson:"idrac_user"`
		IdracPass string `bson:"idrac_pass"`
	}

	// AutoConfig is
	AutoConfig struct {
		Common   CommonParams
		Pool     []Machine
		Profiles []string
		Builds   []string
		IP       string
		PORT     string
		db       *mgo.Database
		lock     *sync.Mutex
	}
)

// NewAutoConfig is
func NewAutoConfig(db *mgo.Database, ip, port string) *AutoConfig {
	ac := &AutoConfig{db: db, lock: &sync.Mutex{}, IP: ip, PORT: port}
	ac.initAutoConfig()
	ac.updateProfiles()
	ac.updateBuilds()
	return ac
}

func (ac *AutoConfig) initAutoConfig() {
	err := ac.db.C(colCommon).Find(nil).One(&ac.Common)
	if err != nil {
		log.Panic(err)
	}
	err = ac.db.C(colMachinePool).Find(nil).All(&ac.Pool)
	if err != nil {
		log.Panic(err)
	}
}

// GetSystemNic is
func (ac AutoConfig) GetSystemNic(bkrName string) ([]string, error) {
	for _, m := range ac.Pool {
		if m.BkrName == bkrName {
			return []string{m.NicName, m.Mac}, nil
		}
	}
	return []string{}, fmt.Errorf("can not found bkr name %s", bkrName)
}

// UpdateProfiles is
func (ac *AutoConfig) updateProfiles() {
	cb := cobbler.NewCobbler(
		ac.Common.CobblerAPI, ac.Common.CobblerCred[0], ac.Common.CobblerCred[1])
	ac.lock.Lock()
	ac.Profiles = cb.ListProfiles(regexp.MustCompile(rhvhSearchPattern))
	ac.lock.Unlock()
}

// UpdateProfiles is
func (ac *AutoConfig) UpdateProfiles() {
	ticker := time.NewTicker(30 * 60 * time.Second)
	log.Info("start to updating profiles every 30min")
	go func() {
		for t := range ticker.C {
			log.Infof("update cobbler profile at %s", t)
			ac.updateProfiles()
		}
	}()
}

func (ac *AutoConfig) updateBuilds() {
	if !helpers.FileExists(buildPath) {
		log.Errorf("%s not exists", buildPath)
		ac.Builds = []string{"Error To Get Builds"}
		return
	}
	ac.Builds = []string{}
	err := filepath.Walk(buildPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Errorf("prevent panic by handling failure accessing a path %q: %v\n", buildPath, err)
			return err
		}
		if info.IsDir() && info.Name() != "crawled.rhvh4x_img" {
			ac.Builds = append(ac.Builds, info.Name())
		}
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", buildPath, err)
	}
}

// UpdateBuilds is
func (ac *AutoConfig) UpdateBuilds() {
	ticker := time.NewTicker(30 * 60 * time.Second)
	log.Info("start to updating builds every 30min")
	go func() {
		for t := range ticker.C {
			log.Infof("update RHVH builds at %s", t)
			ac.updateBuilds()
		}
	}()
}
