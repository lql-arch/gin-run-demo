package service

import (
	"douSheng/sql"
)

// LoginReset 监视登录状态,用于置空状态
func LoginReset() {
	sql.Reset()
}
