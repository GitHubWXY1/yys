# Service和HeadlessService

service有三个类型：ClusterIP[默认类型]，NodePort、LoadBalancer

每一个service对象会有一个全域限定名，创建后会在**DNS**内创建一个 A记录【Address Record】=> [全限定名，clusterIP]

创建的流程：编写service的yaml配置，之后apply，然后k8s会为service分配一个Cluster IP，会在DNS组件里面添加一个A记录。

访问的流程：可以通过servicename访问，经过DNS解析，然后拿着ClusterIP去节点里面找到相应绑定有Pod的节点【ClusterIP->Node 此流程是k8s网络模型吗？】，然后在节点上通过kube-proxy进行**负载均衡**将流量送到合适的pod上。

kube-proxy的负载均衡策略是：**IPtables** 和 **IPVS** 



HeadlessService：不分配ClusterIP。设置是spec.clusterIP=None。广泛用于有状态服务

#### 名词解释

##### 全域限定名

格式一般是这样： <service-name>.<namespace>.svc.cluster.local