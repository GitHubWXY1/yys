StatefulSet相比Deployment

StatefulSet是和ReplicaSet相比较的。Deployment可以管理RS和无状态pod，Deployment是RS的高级管理工具。

spec.serviceName设置

StatefulSet特性：

- 稳定的、唯一的网络标识符。
- 稳定的、持久的存储。
- 有序的、优雅的部署和扩缩。
- 有序的、自动的滚动更新

StatefulSet 示例

* 如何创建StatefulSet
* StatefulSet怎样管理它的pod
* 如何删除StatefulSet
* 如何对StatefulSet进行扩容/缩容
* 如何更新一个StatefulSet的pod

