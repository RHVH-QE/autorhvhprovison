liveimg --url={{.LiveImgURL}}

clearpart --all

autopart --type=thinp

rootpw --plaintext redhat

timezone --utc Asia/Harbin

zerombr

text

reboot

%post --erroronfail
imgbase layout --init
{{.PostScript01}}
{{.PostScript02}}
%end
