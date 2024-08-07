## 学习目标

1. 

2. 能够搭建与使用docker私有仓库

## Docker
docker是使用容器(Container)技术实现的.

镜像(image)来创造容器，注册中心(Registry)可以是公有(docker hub)或私有。

一个最基础的容器一般是一个OS，such as Ubuntu、centOS

### 在Ubuntu上安装Docker

1. 卸载可能会冲突的包

   ```sh
   for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do sudo apt-get remove $pkg; done
   ```

2. 在apt中安装Docker repository

   ```sh
    # 添加 Docker 的官方 GPG 密钥:
    sudo apt-get update
    sudo apt-get install ca-certificates curl
    sudo install -m 0755 -d /etc/apt/keyrings
    # 官方命令 sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
    # 使用阿里云镜像地址
    sudo curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
    sudo chmod a+r /etc/apt/keyrings/docker.asc

    # 向 Apt source添加repository:
    # 官方命令
    # echo \
    #  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] #https://download.docker.com/linux/ubuntu \
    #  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
    #  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
    # 使用阿里云镜像
    echo \
      "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://mirrors.aliyun.com/docker-ce/linux/ubuntu \
      $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
      sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

    sudo apt-get update
   ```

3. 安装Docker

   ```sh
   sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
   ```

4. 经典helloWorld

   ```sh
   sudo docker run hello-world
   ```
   >此时可能由于docker hub 连接不上，拉取镜像失败
   >
   >这里推荐使用https://atomhub.openatom.cn/进行镜像的拉取
   >
   >```sh
   >sudo docker pull hub.atomgit.com/library/hello-world:latest
   >```
   >
   >然后再运行
   >
   >```sh
   >sudo docker run atomhub.openatom.cn/library/hello-world:latest
   >```
5. uninstall

   * 卸载 Docker Engine、 CLI、 Container 和 Docker Compose 包:

   ```sh
   sudo apt-get purge docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin docker-ce-rootless-extras
   ```

   * 需手动卸载docker object such as 镜像、容器、卷或自定义配置文件

   ```sh
   sudo rm -rf /var/lib/docker
   sudo rm -rf /var/lib/containerd
   ```
   
      

### Docker服务的启动和停止

使用systemctl命令

```sh
# 启动服务
systemctl start docker
# 关闭服务
systemctl stop docker
# 重启服务
systemctl restart docker
# 服务开机自启
systemctl enable docker
# 查看服务状态
systemctl status docker
```

docker内置的一些命令

```sh
#
docker info
#
docker --help
```

### 镜像相关命令

```sh
# 查看本地镜像
docker images
# 搜索远程镜像信息
docker search hello-world
# 拉取远程镜像信息
docker pull hello-world
# 删除镜像 rmi-【remove image】 镜像名称/ID
docker rmi hello-world
# 
docker rmi `docker images -q`
```

### 容器相关命令

* 查看容器

```sh
# 查看当前运行容器 ps-【process status】
docker ps
# 查看历史中所有容器 -a-【--all】包括停止的容器
docker ps -a
# 查看容器相关信息
docker inspect 容器名称/ID
# 查看容器指定相关信息
docker inspect --format='{{.一级属性.二级属性}}' 容器名称/ID
```

* 创建和启动容器

1. 交互式创建

```sh
docker run -it --nname=容器名称 镜像名称:TAG/镜像ID /bin/bash
```

交互式启动，提供命令行界面

在命令行中退出容器

```sh
exit
```

2. 守护式创建

```sh
docker run -di --name=容器名称 镜像名称:TAG/镜像ID 
```

在后台运行

```sh
#	进入此容器
docker exec -it 容器名称/容器ID /bin/bash
```

* 停止容器

  ```sh
  docker stop 容器名称/容器ID
  ```

* 启动容器

  ```sh
  docker start 容器名称/容器ID
  ```

* 文件拷贝 ->目录挂载[映射]

  1. 从宿主机到容器

     ```sh
     docker cp 文件路径 容器名称:容器内路径
	   ```
  
	2. 从容器到宿主机
  
     ```sh
	   docker cp 容器名称:容器内路径 文件路径
	   ```
	
	3. 目录挂载(创建时挂载)，挂载后目录内文件会auto 在容器和宿主机中都有一份
	
	   ```sh
	   docker run -di -v 主机目录:容器内目录 --name=容器名 镜像名:TAG/镜像ID
	   ```
	
* 删除容器
	
	正在运行的容器不能删除
	
	```sh
	docker stop 容器
	docker rm 容器
	```
	
	