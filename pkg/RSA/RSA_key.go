package RSA

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

// 读取PKCS1格式私钥
func ReadRSAPKCS1PrivateKey(path string) (*rsa.PrivateKey, error) {
	// 读取文件
	context, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// pem解码
	pemBlock, _ := pem.Decode(context)
	// x509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	return privateKey, err
}

// 读取公钥(包含PKCS1和PKCS8)
func ReadRSAPublicKey(path string) (*rsa.PublicKey, error) {
	var err error
	// 读取文件
	readFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// 使用pem解码
	pemBlock, _ := pem.Decode(readFile)
	var pkixPublicKey interface{}
	if pemBlock.Type == "RSA PUBLIC KEY" {
		// -----BEGIN RSA PUBLIC KEY-----
		pkixPublicKey, err = x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	} else if pemBlock.Type == "PUBLIC KEY" {
		// -----BEGIN PUBLIC KEY-----
		pkixPublicKey, err = x509.ParsePKIXPublicKey(pemBlock.Bytes)
	}
	if err != nil {
		return nil, err
	}
	publicKey := pkixPublicKey.(*rsa.PublicKey)
	return publicKey, nil
}
