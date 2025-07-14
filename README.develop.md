# 开发
## 创建新表
* 手动创建数据库（可能存在自动的方法）

## 运行代码
### 运行最新代码
* go run main.go server -c config/settings.yml
### 运行代码，且更新接口数据（可能是Swagger文档）
* go run main.go server -c config/settings.yml -a true

## 更新到开发环境
### 开发环境
* 腾讯云: youkang1988@163.com
* 服务器：区块链系统(111.230.167.45)
* 路径：/opt/go-admin
### 本地上传到git
* git push -u origin main
### 服务器下载代码
* git clone

## 开发环境运行
### 更新docker image
* docker build -t go-admin:latest .
### 根据docker-compose 运行docker image
* docker compose up -d
### 查看docker container中的terminal
* docker logs -f go-admin


## 微信登录token有效期，在config/setting.yml中设置
