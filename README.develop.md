# 开发信息记录

## 运行代码
### 运行最新代码
```bash
# 本地
go run main.go server -c config/settings.yml
````
### 运行代码，且更新接口数据（可能是Swagger文档）
```bash
# 本地
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
* 连接服务器（服务器已经配置了SSH密钥，无需输入密码）
```bash
# 本地
ssh root@111.230.167.45
```
* 项目路径
```bash
# 服务器
cd /opt
```
* git下载代码 (国内服务器通过https下载代码会被墙，可以使用ssh的方式，需要在服务器设置好SSH key（ssh-keygen -t ed25519 -C "your_email@example.com"
  ），然后把公钥（cat ~/.ssh/id_ed25519.pub）复制到github设置ssh的页面（https://github.com/settings/keys）)
```bash
# 服务器
git clone git@github.com:FeiyangTan/go-admin.git
````
* 同步git管理的以外文件".env"用于标注环境变量
```bash
# 本地
scp .env root@111.230.167.45:/opt/go-admin
````
* 下载go（如果服务器没有go或者go版本不对）
```bash
# 服务器
  cd /tmp
  wget https://golang.google.cn/dl/go1.24.5.linux-arm64.tar.gz
  
```
* 编译go可执行文件
```bash
# 服务器
go mod tidy
GOOS=linux GOARCH=amd64 go build -o go-admin main.go
```
* 把本地构建好的docker image上传到服务器（服务器构建不了docker image）
1. 本地构建镜像（适用于服务器linux/amd64的image）
```bash
# 本地
# 设置buildx(只需设置一次)
docker buildx create --use --name amd64builder
# 如果之前已经设置过了（切换buildx模式）
（
# 查看当前所有 buildx 实例
docker buildx ls
# 切换到已有的 amd64builder
docker buildx use amd64builder
）

# 在 M1/M2 上构建 amd64 镜像并直接输出成 tar 文件
docker buildx build \
  --platform linux/amd64 \
  -t aixiaoqi-server:latest \
  --output type=docker,dest=./aixiaoqi-server.tar \
  .
  
# 上传到服务器
scp aixiaoqi-server.tar root@111.230.167.45:/opt/go-admin

```
2. 服务器上运行镜像
```bash
# 服务器
cd /opt/go-admin
docker load -i aixiaoqi-server.tar
docker image ls #查看image列表
docker compose up -d
docker logs -f qixiaoqi-server 
```

### 更新代码
* 本地上传到git
```bash
git push -u origin main
````
### 服务器下载代码
```bash
git pull
````

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
