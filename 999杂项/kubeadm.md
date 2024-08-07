在主节点执行初始化的命令

kubeadm init --apiserver-advertise-address $(hostname -i) --pod-network-cidr 10.5.0.0/16

apiserver-advertise-address用来指定API服务器的广播地址`$(hostname -i)` 会输出当前节点的 IP 地址

指定Pod网络的CIDR范围

pod-network-cidr 10.5.0.0/16 【网络分配为10.5.0.0-10.5.255.255 去除全零和全一。可以为pod分配的ip地址为】