package util

import (
	"encoding/hex"
	"crypto/md5"
)

func Md5pwd(pwd string) string{
	
	md5ctx := md5.New()
	md5ctx.Write([]byte(pwd))
	cipherStr := md5ctx.Sum(nil)
	md5pwd := hex.EncodeToString(cipherStr)

	return md5pwd
}