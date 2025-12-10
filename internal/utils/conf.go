package utils

import "flag"

var (
	PASSWD_FILE *string
	BASIC_FILE  *string
)

func Conf() {
	BASIC_FILE = flag.String("passwd", "passwd", "local auth file")
	PASSWD_FILE = flag.String("ocpasswd", "/etc/ocserv/ocpasswd", "ocserv password file")
	flag.Parse()
}
