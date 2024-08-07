# architecture

Docker uses a client-server architecture. 客户-服务端架构

client :又称：（docker）  use commands such as docker run 、docker build 、docker pull 

host ：**docker daemon**守护进程 又称：(dockerd) listens for Docker API request and manages **Docker object** such as images,containers, networks, and volumes.

Docker Desktop：包含了**dockerd**、**docker**、**docker Compose**、**docker content trust**、**kubernetes** and **Credential Helper**。

Docker registry：存储（stores）docker images 的地方。docker hub 是任何人都可以使用的 docker registry。docker default looks for images on Docker Hub . 

docker pull : 从**configured registry** 【registry 注册表】中获取 image。docker run ：运行image。docker push：将image push 到 **configured registry**

>registry 是可以配置的，默认配置是Docker Hub

当在使用Docker时，实质就是在使用 **Docker Object**

包含：

## Image

镜像，即一个只读模板。一个image 是基于另一个image并进行一些额外的配置。

## Container

容器: a container is a runnable instance of an image. You can create, start,stop,move,or delete a container using the Docker API or     CLI.

容器由创建它的镜像和创建时的参数定义。**容器被删除时**，任何未被存储在 persistent storage (持久存储)的状态更改都会disappear.

>docker run -i -t ubuntu /bin/bash
>
>运行Ubuntu容器，以交互方式 附加到本地命令行会话 ，并运行 /bin/bash
>
>