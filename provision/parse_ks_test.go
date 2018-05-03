package provision

import (
	"testing"
)

func TestParseKickstarts(t *testing.T) {
	ParseKickstarts(
		"atv_bonda_02.ks",
		"http://10.66.8.175:5060/crawled.rhvh4x_img/redhat-virtualization-host-4.2-20180420.0/redhat-virtualization-host-4.2-20180420.0.x86_64.liveimg.squashfs",
		"10.66.8.150",
		"3000",
		"dell-per515-01.lab.eng.pek2.redhat.com",
	)
}
