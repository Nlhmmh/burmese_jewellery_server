package auth

import "strings"

var (
	whiteList = []string{
		"/api/admin/login",
	}

	adminList = []string{
		"/api/admin/account_admin",
		"/api/admin/account_admin/:user_id",
	}

	adminRoleAdminOnlyList = []string{
		"/api/admin/account_admin",
		"/api/admin/account_admin/:user_id",
	}

	containWhiteList = []string{
		"/api/file/media",
	}
)

func CheckJWTWhiteList(path string) bool {
	for _, p := range whiteList {
		if path == p {
			return true
		}
	}
	return false
}

func CheckContainWhiteList(path string) bool {
	for _, p := range containWhiteList {
		if strings.Contains(path, p) {
			return true
		}
	}
	return false
}

func CheckAdminList(path string) bool {
	for _, p := range adminList {
		if path == p {
			return true
		}
	}
	return false
}

func CheckAdminRoleAdminOnlyList(path string) bool {
	for _, p := range adminRoleAdminOnlyList {
		if path == p {
			return true
		}
	}
	return false
}