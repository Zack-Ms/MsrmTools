package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"io"
)

func Md5Encrypt(plainText, salt string) string {
	md5 := md5.New()
	md5.Write([]byte(plainText + salt))
	cipherStr := md5.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

/**
 * RSA加密字符串
 */
func RsaEncrypt(value string, publicKey string) (string, error) {
	key, _ := base64.StdEncoding.DecodeString(publicKey)
	pubKey, _ := x509.ParsePKIXPublicKey(key)
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey.(*rsa.PublicKey), []byte(value))
	return base64.StdEncoding.EncodeToString(encryptedData), err
}

/**
 * RSA解密字符串
 */
func RsaDecrypt(value string, privateKey string) (string, error) {
	encryptedDecodeBytes, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}
	key, _ := base64.StdEncoding.DecodeString(privateKey)
	prvKey, _ := x509.ParsePKCS1PrivateKey(key)
	originalData, err := rsa.DecryptPKCS1v15(rand.Reader, prvKey, encryptedDecodeBytes)
	return string(originalData), err
}

// AES 加密 CBC
func AesEncryptCBC(origData, key string) string {
	plaintext, secureKey := []byte(origData), []byte(key)
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(secureKey)
	blockSize := block.BlockSize()                                    // 获取秘钥块的长度
	plaintext = PKCS5Padding(plaintext, blockSize)                    // 补全码
	blockMode := cipher.NewCBCEncrypter(block, secureKey[:blockSize]) // 加密模式
	encrypted := make([]byte, len(plaintext))                         // 创建数组
	blockMode.CryptBlocks(encrypted, plaintext)                       // 加密
	return base64.StdEncoding.EncodeToString(encrypted)
}

// AES 解密 CBC
func AesDecryptCBC(encrypted, key string) string {
	plaintext, _ := base64.StdEncoding.DecodeString(encrypted)
	secureKey := []byte(key)
	block, _ := aes.NewCipher(secureKey)                              // 分组秘钥
	blockSize := block.BlockSize()                                    // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, secureKey[:blockSize]) // 加密模式
	decrypted := make([]byte, len(plaintext))                         // 创建数组
	blockMode.CryptBlocks(decrypted, plaintext)                       // 解密
	decrypted = PKCS5UnPadding(decrypted)                             // 去除补全码
	return string(decrypted)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncryptCFB(origData []byte, key []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted
}
func AesDecryptCFB(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted
}
