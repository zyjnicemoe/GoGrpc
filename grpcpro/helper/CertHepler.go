package helper

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

//服务端证书
func GetServerCreds() credentials.TransportCredentials {
	crt := "conf/server/server.pem"
	key := "conf/server/server.key"
	cert, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	}
	certPool := x509.NewCertPool()

	ca, err := ioutil.ReadFile("conf/ca.crt")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	return creds
}

//客户端证书
func GetClientCreds() credentials.TransportCredentials {
	crt := "conf/client/client.pem"
	key := "conf/client/client.key"
	cert, err := tls.LoadX509KeyPair(crt, key)

	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("conf/ca.crt")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "zyjblogs.cn",
		RootCAs:      certPool,
	})
	return creds
}

func Get() credentials.TransportCredentials {
	crt := "conf/server/server.pem"
	key := "conf/server/server.key"
	return InitHttpsClient(crt, key, "server")
}
func InitHttpsClient(keyPem, certPem, pemPass string) credentials.TransportCredentials {
	// 读取私钥文件
	keyBytes, err := ioutil.ReadFile(keyPem)
	if err != nil {
		panic("Unable to read keyPem")
	}
	// 把字节流转成PEM结构
	block, rest := pem.Decode(keyBytes)
	if len(rest) > 0 {
		panic("Unable to decode keyBytes")
	}
	// 解密PEM
	der, err := x509.DecryptPEMBlock(block, []byte(pemPass))
	if err != nil {
		panic("Unable to decrypt pem block")
	}
	// 解析出其中的RSA 私钥
	key, err := x509.ParsePKCS1PrivateKey(der)
	if err != nil {
		panic("Unable to parse pem block")
	}
	// 编码成新的PEM 结构
	keyPEMBlock := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)
	// 读取证书文件
	certPEMBlock, err := ioutil.ReadFile(certPem)
	if err != nil {
		panic("Unable to read certPem")
	}
	// 生成密钥对
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		panic("Unable to read privateKey")
	}
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(certPEMBlock); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "zyjblogs.cn",
		RootCAs:      certPool,
	})
	return creds
}
