package route

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/dracher/autorhvhprovison/provision"
	"github.com/dracher/autorhvhprovison/utils/cobbler"
	"github.com/dracher/autorhvhprovison/utils/power/idrac"
	"github.com/gin-gonic/gin"
)

const (
	argsTpl = "inst.ks=http://%s:%s/ks/%s "
)

type provisionParams struct {
	ImgURL      string `json:"img_url"`
	KsName      string `json:"ks_name"`
	BkrName     string `json:"bkr_name"`
	ProfileName string `json:"profile_name"`
	SSHPK       string `json:"ssh_pk"`
	HostType    string `json:"host_type"`
}

// ProvisionHandler is
func ProvisionHandler(c *gin.Context) {
	var params provisionParams
	cfg := c.MustGet("cfg").(*provision.AutoConfig)

	if c.BindJSON(&params) == nil {
		provision.ParseKickstarts(params.KsName, params.ImgURL, cfg.IP, cfg.PORT, params.BkrName, params.SSHPK)
		cb := cobbler.NewCobbler(cfg.Common.CobblerAPI, cfg.Common.CobblerCred[0], cfg.Common.CobblerCred[1])

		nic, _ := cfg.GetSystemNic(params.BkrName)
		cb.NewSystem(
			params.BkrName, params.ProfileName, "auto-testing", "testing",
			fmt.Sprintf(argsTpl, cfg.IP, cfg.PORT, params.KsName),
			nic)

		if params.HostType == "idrac" {
			for _, m := range cfg.Pool {
				if m.BkrName == params.BkrName {
					host := idrac.NewHost(m.BkrName, m.IdracUser, m.IdracPass)
					err := host.SetPowerState("GracefulRestart")
					if err != nil {
						c.String(http.StatusBadRequest, err.Error())
						return
					}
					break
				}
			}
		} else if params.HostType == "beaker" {
			log.Warn("Deprecated")
		} else {
			// TODO FIXME
		}

		c.String(http.StatusOK, "OK")
		return
	}
	c.String(http.StatusBadRequest, "Wrong")
}

// ProvisionDoneHandler is
func ProvisionDoneHandler(c *gin.Context) {
	cfg := c.MustGet("cfg").(*provision.AutoConfig)

	ip := c.Param("ip")
	bkrName := c.Param("bkrname")
	log.Info(ip, bkrName)
	cb := cobbler.NewCobbler(
		cfg.Common.CobblerAPI, cfg.Common.CobblerCred[0], cfg.Common.CobblerCred[1])
	cb.RemoveSystem(bkrName)

	c.String(http.StatusOK, "OK")
}

// FetchCobblerProfilesHandler is
func FetchCobblerProfilesHandler(c *gin.Context) {
	cfg := c.MustGet("cfg").(*provision.AutoConfig)
	c.JSON(http.StatusOK, cfg.Profiles)
}

// FetchRHVHBuildsHandler is
func FetchRHVHBuildsHandler(c *gin.Context) {
	cfg := c.MustGet("cfg").(*provision.AutoConfig)
	c.JSON(http.StatusOK, cfg.Builds)
}

/* ================= Route Section =========================== */

// RegisteRoute is
func RegisteRoute(r *gin.RouterGroup) {
	r.POST("/provision/start", ProvisionHandler)
	r.GET("/provision/done/:ip/:bkrname", ProvisionDoneHandler)
	r.GET("/provision/profiles", FetchCobblerProfilesHandler)
	r.GET("/provision/builds", FetchRHVHBuildsHandler)
}
