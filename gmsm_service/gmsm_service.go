package gmsm_service

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"math/big"
	"strings"
)

var (
	defaultUid = []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38}
)

// SM3WithSM2Sign SM3WithSM2签名 Hex ToUpper
func SM3WithSM2Sign(privateKey *sm2.PrivateKey, forSignStr string, uid []byte) (string, error) {

	if uid == nil {
		uid = defaultUid
	}

	r, s, err := sm2.Sm2Sign(privateKey, []byte(forSignStr), uid, rand.Reader)
	if err != nil {
		return "", err
	}

	rBytes, sBytes := r.Bytes(), s.Bytes()
	if rLen := len(rBytes); rLen < 32 {
		rBytes = append(make([]byte, 32-rLen), rBytes...)
	}
	if sLen := len(sBytes); sLen < 32 {
		sBytes = append(make([]byte, 32-sLen), sBytes...)
	}

	var buffer bytes.Buffer
	buffer.Write(rBytes)
	buffer.Write(sBytes)

	return strings.ToUpper(hex.EncodeToString(buffer.Bytes())), nil
}

// SM3Sum SM3摘要 Hex ToUpper
func SM3Sum(forSignStr string) string {
	h := sm3.New()
	h.Write([]byte(forSignStr))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

// TransHexToSm2PrivateKey 将16进制私钥转换成sm2私钥实例
func TransHexToSm2PrivateKey(HexPrivateKey string) *sm2.PrivateKey {

	d := new(big.Int)
	d.SetString(HexPrivateKey, 16)
	privateKey := new(sm2.PrivateKey)
	privateKey.D = d
	curve := sm2.P256Sm2()
	privateKey.PublicKey.Curve = curve
	privateKey.PublicKey.X, privateKey.PublicKey.Y = curve.ScalarBaseMult(d.Bytes())

	return privateKey
}
