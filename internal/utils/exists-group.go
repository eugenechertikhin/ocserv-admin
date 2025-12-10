package utils

import "ocserv-admin/internal/model"

func ExistGroup(base *[]model.Base, s, grp string) bool {
	for _, r := range *base {
		if r.Group == grp && r.User == s {
			return true
		}
	}
	return false
}
