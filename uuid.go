package ghelper

import "github.com/rs/xid"

// GUID 生成分布式字符串 ID，用于全局唯一 ID
func (Gu *GUid) GUID() string {
	return xid.New().String()
}
