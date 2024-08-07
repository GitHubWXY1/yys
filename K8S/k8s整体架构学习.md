* Pod：面向**应用**【容器】的**逻辑主机**模型。pod中容器共享网络、IPC、UTS、volume。保证容器间的数据共享和通信。

* Node：Pod真正运行的主机【物理机、虚拟机】。管理Pod的进程有:container runtime、kubelet、kube-proxy。

假设你有一个停车场（Kubernetes 集群），里面有多个停车位（Nodes）。你可以在管理系统中添加一个停车位的记录（Node 对象），但是这个停车位是否真的存在并且可用（物理存在的停车位）需要实际检查。如果检查发现这个停车位不存在或者被堵住了（节点不可用），那么你就不会让车子（Pod）停在那里。

**Node实体**是外部创建的，k8s只是为了方便管理node而抽象了一个node对象概念

**Node 实体**：由外部系统创建和管理的实际服务器。

**Node 对象**：Kubernetes 中对 Node 实体的抽象和描述。

**Kubernetes 管理**：使用 Node 对象来管理资源和调度 Pod，但实际节点的创建和维护由外部系统负责。

* Lable：关联到**对象**【Pod等对象】上的key/value，可以用selector对标签进行选择。

	k8s中的对象有：以下列举的内容都是 kubernetes 中的 Object，这些对象都可以在 yaml 文件中作为一种 API 类型来配置。
	这些对象通常以 YAML 或 JSON 格式定义，并通过 kubectl 命令行工具或 Kubernetes API 创建和管理。每个对象都有以下基本部分：
	
	apiVersion：定义对象的 API 版本。
	kind：定义对象的类型。
	metadata：包含对象的名称、命名空间、标签和注释等信息。
	spec：定义对象的期望状态，包含对象的详细配置。
	status（可选）：包含对象的当前状态，由 Kubernetes 系统填充。

| 类别     | 名称                                                         |
| -------- | ------------------------------------------------------------ |
| 资源对象 | Pod、ReplicaSet、ReplicationController、Deployment、StatefulSet、Job、CronJob、HorizontalPodAutoscaling |
| 配置对象 | Node、Namespace、Service、Secret、ConfigMap、Ingress、Label、ThirdPartyResource、ServiceAccount |
| 存储对象 | Volume、Persistent Volume                                    |
| 策略对象 | SecurityContext、ResourceQuota、LimitRange                   |

* Service：一类Pod的逻辑分组。创建时会创建一个同名的endpoint，记录service关联的pod和IP端口信息

![image-20240802142515989](C:\Users\47212\AppData\Roaming\Typora\typora-user-images\image-20240802142515989.png)

* Replication Controller、ReplicaSet、Deployment: RC保证有指定Pod副本正在运行-->RS对RC 在Selector上有升级，支持集合selector --> Deployment为Pod和RS提供声明式更新。

  无状态应用可以、有状态应用？

* 一个pod的创建流程：

通过RS创建pod：1.kubectl向apiServer请求 **Create RS** 2.apiServer处理 **Create RS**请求，将RS信息写入etcd

3.etcd通知/发送 apiserver 一个 **RS create event** 4.apiServer通知Controller manager 一个 **RS create event**

5.Controller Manager 处理 **RS Create event**  并调用 apiServer **create Pod** 6.apiServer 处理 **create Pod** 请求，将pod信息写入etcd 7.etcd通知apiserver  **pod create Event** 8.apiserver 通知 scheduler **Pod create Evenet** 9.scheduler为pod调动节点，通知apiserver 更新 **Pod信息** 10.apiServer更新Pod信息到etcd 11.etcd通知apiServer **Pod Update Event** 12.apiServer通知kubelet **Pod Create Event** 创建pod 13.kubectl创建pod

 
