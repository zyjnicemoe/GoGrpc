# 生成含有SAN的证书

## 第一种pem方式

> ca.conf

```conf
[ req ]
default_bits       = 4096
distinguished_name = req_distinguished_name
 
[ req_distinguished_name ]
countryName                 = Country Name (2 letter code)
countryName_default         = CN
stateOrProvinceName         = State or Province Name (full name)
stateOrProvinceName_default = HuBei
localityName                = Locality Name (eg, city)
localityName_default        = WuHan
organizationName            = Organization Name (eg, company)
organizationName_default    = zyjblogs.cn
commonName                  = Common Name (e.g. server FQDN or YOUR name)
commonName_max              = 64
commonName_default          = Ted CA Test
```

> client.conf

```conf
[ req ]
default_bits       = 2048
distinguished_name = req_distinguished_name
req_extensions     = req_ext
 
[ req_distinguished_name ]
countryName                 = Country Name (2 letter code)
countryName_default         = CN
stateOrProvinceName         = State or Province Name (full name)
stateOrProvinceName_default = HuBei
localityName                = Locality Name (eg, city)
localityName_default        = WuHan
organizationName            = Organization Name (eg, company)
organizationName_default    = zyjblogs.cn
commonName                  = Common Name (e.g. server FQDN or YOUR name)
commonName_max              = 64
commonName_default          = zyjblogs.cn
 
[ req_ext ]
basicConstraints = CA:TRUE  # 表明要签发CA证书
subjectAltName = @alt_names
 
[alt_names]
DNS.1   = zyjblogs.cn
IP      = 127.0.0.1
```

> server.conf

```conf
[ req ]
default_bits       = 2048
distinguished_name = req_distinguished_name
req_extensions     = req_ext
 
[ req_distinguished_name ]
countryName                 = Country Name (2 letter code)
countryName_default         = CN
stateOrProvinceName         = State or Province Name (full name)
stateOrProvinceName_default = HuBei
localityName                = Locality Name (eg, city)
localityName_default        = WuHan
organizationName            = Organization Name (eg, company)
organizationName_default    = zyjblogs.cn
commonName                  = Common Name (e.g. server FQDN or YOUR name)
commonName_max              = 64
commonName_default          = zyjblogs.cn
 
[ req_ext ]
subjectAltName = @alt_names
 
[alt_names]
DNS.1   = zyjblogs.cn
IP      = 127.0.0.1
```

生产带密码的证书

```bash
#生成ca秘钥，得到ca.key                       
openssl genrsa  -aes256 -passout pass:123456 -out ca.key 4096 
#生成ca证书签发请求，得到ca.csr
openssl req -new -sha256  -passin pass:123456  -out ca.csr  -key ca.key  -config ca.conf
#生成ca根证书，得到ca.crt -sha256
openssl x509  -req  -days 3650 -in ca.csr -passin pass:123456 -signkey ca.key -out ca.crt

#生成秘钥，得到server.key
openssl genrsa -aes256 -passout pass:server -out server.key 2048
#生成证书签发请求，得到server.csr 
openssl req -new  -sha256 -passin pass:server -out server.csr -key server.key -config server.conf
#用CA证书生成终端用户证书，得到server.crt
openssl x509 -req -days 3650 -passin pass:123456 -CA ca.crt -CAkey ca.key  -CAcreateserial -in server.csr  -out server.pem -extensions req_ext -extfile server.conf

#生成秘钥，得到client.key
openssl genrsa -aes256 -passout pass:client -out client.key 2048
#生成证书签发请求，得到client.csr
openssl req -new  -sha256 -passin pass:client -out client.csr  -key client.key -config client.conf
#用CA证书生成终端用户证书，得到client.crt
openssl x509 -req -days 3650  -passin pass:123456 -CA ca.crt -CAkey ca.key -CAcreateserial -in client.csr  -out client.pem -extensions req_ext -extfile client.conf
```

不带密码的证书

```bash

#生成ca秘钥，得到ca.key                       
openssl genrsa -out ca.key 4096 
#生成ca证书签发请求，得到ca.csr
openssl req -new  -sha256  -out ca.csr  -key ca.key  -config ca.conf
#生成ca根证书，得到ca.crt -sha256
openssl x509  -req  -days 3650 -in ca.csr -signkey ca.key -out ca.crt

#生成秘钥，得到server.key
openssl genrsa -out server.key 2048
#生成证书签发请求，得到server.csr 
openssl req -new  -sha256 -out server.csr -key server.key -config server.conf
#用CA证书生成终端用户证书，得到server.crt
openssl x509 -req -days 3650 -CA ca.crt -CAkey ca.key  -CAcreateserial -in server.csr  -out server.pem -extensions req_ext -extfile server.conf

#生成秘钥，得到client.key
openssl genrsa -out client.key 2048
#生成证书签发请求，得到client.csr
openssl req -new -sha256 -out client.csr  -key client.key -config client.conf
#用CA证书生成终端用户证书，得到client.crt
openssl x509 -req -days 3650 -CA ca.crt -CAkey ca.key -CAcreateserial -in client.csr  -out client.pem -extensions req_ext -extfile client.conf
```



## 第二种crt方式

> ca.conf

```conf
# ca.conf
[ req ]
default_bits = 2048
default_keyfile = privkey.pem
distinguished_name = req_distinguished_name
# 生成v3版本带扩展属性的证书
req_extensions = v3_req

# 设置默认域名
[ req_distinguished_name ]
# Minimum of 4 bytes are needed for common name
commonName         = www.examples.com
commonName_default = *.zyjblogs.cn
commonName_max     = 64

# 设置两位国家代码
# ISO2 country code only
countryName         = China
countryName_default = CN

# 设置州 或者 省的名字
# State is optional, no minimum limit
stateOrProvinceName         = Province
stateOrProvinceName_default = HuBei

# 设置城市的名字
# City is required
localityName         = City
localityName_default = WuHan

# 设置公司或组织机构名称
# Organization is optional
organizationName         = Organization
organizationName_default = ca

# 设置部门名称
# Organization Unit is optional
organizationalUnitName         = ca
organizationalUnitName_default = ca

# 设置联系邮箱
# Email is optional
emailAddress         = Email
emailAddress_default = 1317453947@qq.com

# 拓展信息配置
[ v3_req ]
#basicConstraints = CA:FALSE # 表明要签发终端证书
basicConstraints = CA:TRUE # 表明要签发CA证书
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
subjectAltName = @alt_names

# 要配置的域名
[alt_names]
DNS.1 = www.zyjblogs.cn
DNS.2 = *.zyjblogs.cn
DNS.3 = localhost
DNS.4 = zyjblogs.cn
IP    = 127.0.0.1
```

> example.com.conf

```conf
# example.com.conf
[ req ]
default_bits = 2048
default_keyfile = privkey.pem
distinguished_name = req_distinguished_name
# 生成v3版本带扩展属性的证书
req_extensions = v3_req

# 设置默认域名
[ req_distinguished_name ]
# Minimum of 4 bytes are needed for common name
commonName         = www.examples.com
commonName_default = *.zyjblogs.cn
commonName_max     = 64


# 设置两位国家代码
# ISO2 country code only
countryName         = China
countryName_default = CN

# 设置州 或者 省的名字
# State is optional, no minimum limit
stateOrProvinceName         = Province
stateOrProvinceName_default = HuBei

# 设置城市的名字
# City is required
localityName         = City
localityName_default = WuHan

# 设置公司或组织机构名称
# Organization is optional
organizationName         = Organization
organizationName_default = zyjblogs

# 设置部门名称
# Organization Unit is optional
organizationalUnitName         = zyjblogs
organizationalUnitName_default = zyjblogs

# 设置联系邮箱
# Email is optional
emailAddress         = Email
emailAddress_default = 1317453947@qq.com

# 拓展信息配置
[ v3_req ]
basicConstraints = CA:FALSE # 表明要签发终端证书
#basicConstraints = CA:TRUE # 表明要签发CA证书
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
subjectAltName = @alt_names

# 要配置的域名
[alt_names]
DNS.1 = www.zyjblogs.cn
DNS.2 = *.zyjblogs.cn
DNS.3 = localhost
DNS.4 = zyjblogs.cn
IP    = 127.0.0.1
```

> client.ext

```conf
extendedKeyUsage=clientAuth
```

```bash
# CA生成
openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -config ca.conf -days 5000 -out ca.crt

# 服务端证书
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -config example.com.conf
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000 -extensions v3_req -extfile example.com.conf

# 客户端证书
openssl genrsa -out client.key 2048
openssl req -new -key client.key -config example.com.conf -out client.csr
# 1、创建文件client.ext 内容：extendedKeyUsage=clientAuth
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -extfile client.ext -out client.crt -days 5000 -extensions v3_req -extfile example.com.conf  # 必须要加-extensions v3_req -extfile example.com.conf

# 查看数字证书内容
openssl x509 -text -in client.crt -noout # 查看client.crt内容
openssl x509 -text -in server.crt -noout
```



