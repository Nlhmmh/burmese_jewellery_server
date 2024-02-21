package auth

import "strings"

var (
	whiteList = []string{
		"/api/admin/login",
	}

	adminList = []string{
		"/api/admin/account_admin/:account_admins_id",
	}

	adminRoleAdminOnlyList = []string{
		"/api/admin/account_admin",
		"/api/admin/account_admin/:account_admins_id",
	}

	containWhiteList = []string{
		"/api", // TODO: remove after google testing

		"/api/file/media",
		"/api/health_check",
		"/api/category",
		"/api/gem",
		"/api/material",
		"/api/jewellery",
		"/api/jewellery/:jewellery_id",
		"/api/faq",
		"/api/profile",
		"/api/auth/email/register",
		"/api/auth/email/login",
		"/api/auth/google/login",
		"/api/auth/google/callback",
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
		if p == "/api/file/media" && strings.Contains(path, p) {
			return true
		}
		if path == p {
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
