package core

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

//生成32位md5字串
func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func getGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return getMd5String(base64.URLEncoding.EncodeToString(b))
}

/**
 * This function is used to create unique IDs for trees and nodes.
 *
 * (consult http://www.ietf.org/rfc/rfc4122.txt).
 *
 * @class createUUID
 * @construCtor
 * @return {String} A unique ID.
**/
func CreateUUID() string {
	return getGuid()
	/*
			var s = [];
		    var hexDigits = "0123456789abcdef";
		    for (var i = 0; i < 36; i++) {
		      s[i] = hexDigits.substr(Math.floor(Math.random() * 0x10), 1);
		    }
		    // bits 12-15 of the time_hi_and_version field to 0010
		    s[14] = "4";

		    // bits 6-7 of the clock_seq_hi_and_reserved to 01
		    s[19] = hexDigits.substr((s[19] & 0x3) | 0x8, 1);

		    s[8] = s[13] = s[18] = s[23] = "-";

		    var uuid = s.join("");
		    return uuid;
	*/
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

//定义注册结构map
type RegisterStructMaps struct {
	nodes map[string]NodeCreator
}

func NewRegisterStructMaps() *RegisterStructMaps {
	return &RegisterStructMaps{
		nodes: make(map[string]NodeCreator),
	}
}

//根据名字注册实例
func (rsm *RegisterStructMaps) Register(name string, factory NodeCreator) {
	rsm.nodes[name] = factory
}

func (rsm *RegisterStructMaps) CheckNode(name string) bool {
	if _, ok := rsm.nodes[name]; ok {
		return true
	}
	return false
}

func (rsm *RegisterStructMaps) GetNode(name string) NodeCreator {
	if node, ok := rsm.nodes[name]; ok {
		return node
	}
	return nil
}

type NodeCreator func() IBaseNode
