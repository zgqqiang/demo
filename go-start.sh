#!/bin/bash

# 下载Go安装包
wget https://golang.google.cn/dl/go1.21.4.linux-amd64.tar.gz

# 解压安装包到/usr/local
tar -C /usr/local -zxf go1.21.4.linux-amd64.tar.gz

# 在/etc/profile文件末尾添加环境变量配置
echo 'export GOROOT=/usr/local/go' >> /etc/profile
echo 'export PATH=$PATH:$GOROOT/bin' >> /etc/profile

# 刷新/etc/profile文件
source /etc/profile

# 配置GOPROXY和GO111MODULE
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=on

echo "Go安装和配置完成!"

# 下载Helm版本，以3.5.3为例
wget https://get.helm.sh/helm-v3.5.3-linux-amd64.tar.gz
 
# 解压缩
tar -zxf helm-v3.5.3-linux-amd64.tar.gz
 
# 移动并给予执行权限
sudo mv linux-amd64/helm /usr/local/bin/helm
sudo chmod +x /usr/local/bin/helm
echo "helm安装和配置完成!"