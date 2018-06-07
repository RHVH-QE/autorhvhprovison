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
		"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDK8/OjonxbM1wePPIm00RdY8CmHkeLE00fLFRvB+56fQksEUHqWB/QgKta4yyv4k5mbz40/4zS/mZ4x5+dVigi+n2/ErrheVsvVnk0yAdUCZ5BiPPZzMbDiTrilvzmdhezHewkWypqsda+cK+iKNI3Ci8hMti0aUjmhztTAhk7ocyvIIPMVGMGJVBTUOTQv/EFrJPNN+te9gLisAMA6q7CXtB15naielIFflAGlsKLNXiYMGSx6nNRg2HIUbfoaLZ7MSZbm/N7uCsFdYo7m/u1hc75rClbLbFj2kr3tHFi8rTcuOVzC8wJbYjfODE8uPI9fwHLt/OJFwTnJ3jcND4b dracher@dhcp-9-25.nay.redhat.com",
	)
}
