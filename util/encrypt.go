package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
)

const modulus = "00e0b509f6259df8642dbc35662901477df22677ec152b5ff68ace615bb7b725152b3ab17a876aea8a5aa76d2e417629ec4ee341f56135fccf695280104e0312ecbda92557c93870114af6c9d05c4f7f0c3685b7a46bee255932575cce10b424d813cfe4875d3e82047b97ddef52741d546b8e289dc6935b3ece0462db0a22b8e7"
const nonce = "0CoJUm6Qyw8W8jud"
const pubKey = "010001"
const iv = "0102030405060708"

// EncryptForm 加密表单参数
func EncryptForm(param string) (string, string, error) {
	secKey := randStr(16) // 生成16位随机字符串
	// aes encrypt for twice
	encText, err := aesEncrypt(param, nonce)
	if err != nil {
		return "", "", err
	}
	encText, err = aesEncrypt(encText, secKey)
	if err != nil {
		return "", "", err
	}
	encSecKey := rsaEncrypt(secKey, pubKey, modulus)
	return encText, encSecKey, nil
}

// MD5Sum return the md5 sum value of input string.
func MD5Sum(s string) string {
	r := md5.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

// aesEncrypt AES加密流程
func aesEncrypt(plain string, key string) (string, error) {
	iv := []byte(iv)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	padding := block.BlockSize() - len([]byte(plain))%block.BlockSize()
	src := append([]byte(plain), bytes.Repeat([]byte{byte(padding)}, padding)...)
	mode := cipher.NewCBCEncrypter(block, iv)
	c := make([]byte, len(src))
	mode.CryptBlocks(c, src)
	return base64.StdEncoding.EncodeToString(c), nil
}

// rsaEncrypt RSA加密流程
func rsaEncrypt(key string, pubKey string, modulus string) string {
	key = reverseStr([]byte(key)) // 翻转 key
	hexRKey := ""
	for _, char := range []rune(key) {
		hexRKey += fmt.Sprintf("%x", int(char))
	}
	bigRKey, _ := big.NewInt(0).SetString(hexRKey, 16)
	bigPubKey, _ := big.NewInt(0).SetString(pubKey, 16)
	bigModulus, _ := big.NewInt(0).SetString(modulus, 16)
	bigRs := bigRKey.Exp(bigRKey, bigPubKey, bigModulus)
	hexRs := fmt.Sprintf("%x", bigRs)
	return addPadding(hexRs, modulus)
}

// addPadding 补零
func addPadding(encText string, modulus string) string {
	ml := len(modulus)
	for i := 0; ml > 0 && modulus[i:i+1] == "0"; i++ {
		ml--
	}
	num := ml - len(encText)
	prefix := ""
	for i := 0; i < num; i++ {
		prefix += "0"
	}
	return prefix + encText
}

// randStr 生成指定长度随机字符串
func randStr(size int) (str string) {
	const rStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < size; i++ {
		idx := rand.Intn(len(rStr))
		str += rStr[idx : idx+1]
	}
	return
}

// reverseStr 翻转字符串
func reverseStr(bs []byte) string {
	for i, j := 0, len(bs)-1; i < len(bs)/2; i++ {
		bs[i], bs[j-i] = bs[j-i], bs[i]
	}
	return string(bs)
}
