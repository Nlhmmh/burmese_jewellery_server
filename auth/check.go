package auth

import "strings"

var (
	whiteList = []string{
		"/api", // TODO: remove after google testing
		"/api/health_check",

		"/api/admin/login",
		"/api/auth/email/register",
		"/api/auth/email/login",
		"/api/auth/google/login",
		"/api/auth/google/callback",

		"/api/category",
		"/api/gem",
		"/api/material",

		"/api/jewellery",
		"/api/jewellery/:jewellery_id",

		"/api/faq",
	}

	containWhiteList = []string{
		"/api/file/media",
	}

	adminList = []string{
		"/api/admin/account_admin/:account_admins_id",

		"/api/admin/account",
		"/api/admin/account/:account_id",
	}

	adminRoleAdminOnlyList = []string{
		"/api/admin/account_admin",
	}

	userList = []string{
		"/api/profile",
		"/api/favourite",
		"/api/cart",
		"/api/order",
		"/api/order/:account_order_id",
	}
)

func CheckWhiteList(path string) bool {
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

func CheckUserList(path string) bool {
	for _, p := range userList {
		if path == p {
			return true
		}
	}
	return false
}
