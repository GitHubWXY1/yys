systemctl status docker 

## docker 构建镜像

构建命令：docker build [选项] <上下文路径/URL/->   *不常用的：-f 指定Dockerfile路径。默认为上下文路径下的Dockerfile*.

`docker build -t imageName .`

* 上下文路径的作用:指定host中Dockerfile文件内一些命令的源路径上下文

	> copy [源路径] [目标路径]中的前者的上下文

## docker 镜像创建容器



## docker构建服务



## docker容器操作

进入容器: docker exec -it 容器 bash