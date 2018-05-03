package provision

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
)

const (
	postScript01 = `
EM1IP=$(ip -o -4 addr show {} | awk -F '[ /]+' '/global/ {{print $4}}')
curl -s http://%s:%s/api/v1/provision/done/$EM1IP/%s`

	postScript02 = `
EM1IP=$(ip -4 a show | awk -F " " '/inet/ { if (match($2, /^10.*/)) print $2 }' | awk -F "/" '{print $1}')
curl -s http://%s:%s/api/v1/provision/done/$EM1IP/%s`

	ksTplPath  = "static/tpl/"
	ksAutoPath = "static/"
)

// KsParams is
type KsParams struct {
	LiveImgURL   string
	PostScript01 string
	PostScript02 string
}

// ParseKickstarts is
func ParseKickstarts(ksName, imgURL, ip, port, bkrName string) {
	tpl, err := template.ParseFiles(ksTplPath + ksName)
	if err != nil {
		log.Error(err)
	}
	fp, err := os.OpenFile(ksAutoPath+ksName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Error(err)
	}
	defer fp.Close()
	if strings.Contains(ksName, "atv_bonda") {
		ksParams := KsParams{imgURL, fmt.Sprintf(postScript02, ip, port, bkrName), ""}
		tpl.Execute(fp, ksParams)
	}
}
