package auth

import (
	"slices"
	"strings"
)

var (
	whiteList = []string{
		"/api", // TODO: remove after google testing
		"/api/health_check",
		"/api/enums",

		"/api/admin/login",
		"/api/auth/email/register",
		"/api/auth/email/otp",
		"/api/auth/email/otp/resend",
		"/api/auth/email/login",
		"/api/auth/google/login",
		"/api/auth/google/callback",

		"/api/category",
		"/api/gem",
		"/api/material",

		"/api/jewellery",
		"/api/jewellery/:jewellery_id",

		"/api/faq",
		"/api/file/:file_name",
	}

	containWhiteList = []string{}

	adminList = []string{
		"/api/admin/account_admin/:account_admins_id",

		"/api/admin/account",
		"/api/admin/account/:account_id",

		"/api/admin/jewellery",
		"/api/admin/jewellery/:jewellery_id",
	}

	adminRoleAdminOnlyList = []string{
		"/api/admin/account_admin",
		"/api/admin/file/:file_name",
	}

	userList = []string{
		"/api/profile",
		"/api/favourite",
		"/api/cart",
		"/api/order",
		"/api/order/:account_order_id",
	}

	bothAdminUserList = []string{
		"/api/file",
	}
)

func CheckWhiteList(path string) bool {
	return slices.Contains(whiteList, path)
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
	return slices.Contains(adminList, path)
}

func CheckAdminRoleAdminOnlyList(path string) bool {
	return slices.Contains(adminRoleAdminOnlyList, path)
}

func CheckUserList(path string) bool {
	return slices.Contains(userList, path)
}

func CheckBothAdminUserList(path string) bool {
	return slices.Contains(bothAdminUserList, path)
}
