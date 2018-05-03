package cobbler

const (
	modifyNicTpl = `<?xml version='1.0'?><methodCall><methodName>modify_system</methodName><params><param><value><string>%s</string></value></param><param><value><string>modify_interface</string></value></param><param><value><struct><member><name>%s</name><value><string>%s</string></value></member></struct></value></param><param><value><string>%s</string></value></param></params></methodCall>`
	kargsTpl     = ""
)
