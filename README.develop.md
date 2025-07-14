# 开发信息记录

## 运行代码
### 运行最新代码
```bash
go run main.go server -c config/settings.yml
````
### 运行代码，且更新接口数据（可能是Swagger文档）
```bash
go run main.go server -c config/settings.yml -a true
````
## 备注信息
* 创建新表后，手动创建数据库（可能存在自动的方法）
* app/admin 中后台接口
* app/wechat 微信端接口
* config/settings.yml 使用的配置文件
* secrets.txt 文件用于消除git中的“秘密”信息（不消除不让上传到git）
  (git filter-repo --force --replace-text secrets.txt)

---------------------------------
## 更新到开发环境
### 开发环境
* 腾讯云: youkang1988@163.com
* 服务器：区块链系统(111.230.167.45)
* 路径：/opt/go-admin
### 首次迁移
* 连接服务器（服务器已经配置了SSH密钥，无需输入密码）(本地运行)
```bash
ssh root@111.230.167.45
```
* 项目路径(服务器运行)
```bash
cd /opt
```
* git下载代码(服务器运行)
* (国内服务器通过https下载代码会被墙，可以使用ssh的方式，需要在服务器设置好SSH key（ssh-keygen -t ed25519 -C "your_email@example.com"
  ），然后把公钥（cat ~/.ssh/id_ed25519.pub）复制到github设置ssh的页面（https://github.com/settings/keys）)
```bash
git clone git@github.com:FeiyangTan/go-admin.git
````
* 下载go（如果服务器没有go或者go版本不对）
```bash
  cd /tmp
  wget https://golang.google.cn/dl/go1.24.5.linux-arm64.tar.gz
  
```
* go可执行文件编译(服务器运行)
```bash
go mod tidy
GOOS=linux GOARCH=amd64 go build -o go-admin main.go
```


### 本地上传到git
* git push -u origin main
### 服务器下载代码
* git clone

## 开发环境运行
### 更新docker image
```bash
docker build -t go-admin:latest .
```
### 根据docker-compose 运行docker image
* docker compose up -d
### 查看docker container中的terminal
* docker logs -f go-admin


## 微信登录token有效期，在config/setting.yml中设置
