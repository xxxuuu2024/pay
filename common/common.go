package common

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"sort"
	"strings"
)

//ASCII排序 并链接成字符串
func AsciiSort(params []byte) (map[string]string, string, error) {
	data := make(map[string]string, 8)
	if err := json.Unmarshal(params, &data); err != nil {
		return nil, "", err
	}
	var key []string
	var content string
	for k, _ := range data {
		key = append(key, k)
	}
	sort.Strings(key)
	for _, v := range key {
		if data[v] == "" {
			continue
		}
		content += v + "=" + data[v] + "&"
	}
	content = strings.TrimRight(content, "&")
	return data, content, nil
}

//错误信息
func ErrMsg(msg string) error {

	return errors.New(msg)

}

func PrivateKeyDecode(filepath string) (*rsa.PrivateKey, error) {

	priByte, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(priByte)
	if block == nil {
		return nil, ErrMsg("DecodeKey Decode err")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}
func PubKeyDecode(filepath string) (*rsa.PublicKey, error) {

	priByte, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(priByte)
	if block == nil {
		return nil, ErrMsg("DecodePem Decode err")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	return pub.(*rsa.PublicKey), err
}

func SHA256Sign(param []byte, key *rsa.PrivateKey) (string, error) {
	h := sha256.New()
	h.Write(param)
	digest := h.Sum(nil)
	s, err := rsa.SignPKCS1v15(nil, key, crypto.SHA256, digest)
	if err != nil {
		return "", err
	}
	data := base64.StdEncoding.EncodeToString(s)
	return data, nil
}
func SHA256SignVerify(param []byte, key *rsa.PublicKey, sign string) error {
	h := sha256.New()
	h.Write(param)
	digest := h.Sum(nil)
	decodeSign, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	err = rsa.VerifyPKCS1v15(key, crypto.SHA256, digest, decodeSign)
	if err != nil {
		return err
	}
	return nil
}

func SHASign(param []byte, key *rsa.PrivateKey) (string, error) {
	h := sha1.New()
	h.Write(param)
	digest := h.Sum(nil)
	s, err := rsa.SignPKCS1v15(nil, key, crypto.SHA1, digest)
	if err != nil {
		return "", err
	}
	data := base64.StdEncoding.EncodeToString(s)
	return data, nil
}
