通过kubectl get pods --watch -l app=nginx查看pod的创建流程

生命周期:

1. pending 【待定】
2. ContainerCreating
3. Running
4. Unkonw 与pod通信中断
5. 