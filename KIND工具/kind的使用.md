kind的使用:

​	**集群操作**

* 创建

  > kind create cluster
  >
  > 例子：kind create cluster --config cluster.yaml --name 1c3w

  cluster.yaml的例子

  ```yaml
  kind: Cluster
  apiVersion: kind.x-k8s.io/v1alpha4
  nodes:
  - role: control-plane
    extraPortMappings:
    - containerPort: 31000  # 将主机 31000 端口映射到容器的 31000 端口
      hostPort: 31000
      listenAddress: "0.0.0.0" # Optional, defaults to "0.0.0.0"
      protocol: tcp # Optional, defaults to tcp
  - role: worker
  - role: worker
  - role: worker
  
  # 配置网络代理 【失败】
  kind: Cluster
  apiVersion: kind.x-k8s.io/v1alpha4
  networking:
    disableDefaultCNI: false
    kubeProxy:
      extraArgs:
        proxy-args: "http://172.27.14.119:7890"
      
  kubelet:
    extraEnv:
      - name: HTTPS_PROXY
        value: "https://172.27.14.119:7890"
      - name: HTTP_PROXY
        value: "http://172.27.14.119:7890"
        
  nodes:
  - role: control-plane
    extraPortMappings:
    - containerPort: 31000  # 将主机 31000 端口映射到容器的 31000 端口
      hostPort: 31000
      listenAddress: "0.0.0.0" # Optional, defaults to "0.0.0.0"
      protocol: tcp # Optional, defaults to tcp
  - role: worker
  - role: worker
  - role: worker
  ```

* 查看

  > kind get clusters

* 销毁

  > kind delete cluster --name 

* 进入

  > kind

### kubectl 

第一个命令:集群信息、上下文是kind-kind

> kubectl cluster-info --context kind-kind

切换管理集群上下文

> kubectl config use-context kind-集群名

获取node节点信息

> kubectl get node

运行一个nginx Deployment并将其暴露出来

> kubectl run --image=nginx nginx-app --port=80 --env="DOMAIN=CLUSTER"

**在宿主机中将docker镜像传入集群各个节点**

> kind load docker-image nginx:my-latest --name kind-1c3w