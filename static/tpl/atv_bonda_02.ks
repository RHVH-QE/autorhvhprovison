#
# KS for vdsm bond anaconda test on dell-per515-01
#
### Language ###
lang en_US.UTF-8

### Timezone ###
timezone Asia/Shanghai

### Keyboard ###
keyboard --vckeymap=us --xlayouts='us'

### Kdump ###

### Security ###

### User ###
rootpw --plaintext redhat
auth --enableshadow --passalgo=md5

### Misc ###
services --enabled=sshd
selinux --enforcing

### Installation mode ###
install
#liveimg url will substitued by autoframework
liveimg --url={{.LiveImgURL}}
text
reboot

### Network ###
network --device=bond0 --bootproto=static --ip=10.73.73.17 --netmask=255.255.252.0 --gateway=10.73.75.254 --bondslaves=em1,em2 --bondopts=mode=active-backup,primary=em2,miimon=100

### Partitioning ###
ignoredisk --only-use=/dev/disk/by-id/scsi-360a9800050334c33424b41762d726954
zerombr
clearpart --all
bootloader --location=mbr
autopart --type=thinp

### Pre deal ###

### Post deal ###
%post --erroronfail
imgbase layout --init

{{.PostScript01}}
{{.PostScript02}}

%end

