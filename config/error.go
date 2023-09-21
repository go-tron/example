package config

import "github.com/go-tron/base-error"

var (
	ErrorAuthorize     = baseError.New("01", "Please login")
	ErrorUserNotExists = baseError.New("01", "Account not exists")

	ErrorSystem                 = baseError.SystemFactory("9999", "系统错误{}")
	ErrorAddressCountReachLimit = baseError.New("2011", "You can have 10 address at most")
	ErrorUserStatus             = baseError.New("2023", "账户已停用")
)
