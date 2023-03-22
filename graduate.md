# 毕业总结
补充一些云原生相关的问题和自己整理的答案：（持续更新中～

## 一个pod从apply deploy.yaml到running，经历了哪些过程？有什么组件起到了作用？

1. 解析部署文件：kube-apiserver接收到了kubectl apply deploy.yaml命令之后，会对部署文件进行解析和验证，包括版本、镜像、容器、资源需求等信息。
2. 创建PodSpec：根据部署文件中定义的内容，kube-apiserver创建PodSpec对象
3. 创建pod： kube-scheduler根据pod spec中定义的tolerant或affinity等信息，将pod调度到可用的node上
4. 容器镜像拉取：被调度到的node上的kubelet组件会检查pod spec中定义的容器镜像是否存在，如果不存在node上，则拉取
5. 创建容器：kubelet组件根据PodSpec中定义的容器配置信息配置容器，包括limit和request之类的资源限制，并将其加入pod中
6. 容器状态检查：liveness probe和readiness probe等，kubelet还会周期性检查容器状态，看容器是否存活、是否在运行中
7. pod状态更新：kubelet在probe成功后，会上报给apiserver。当所有容器都running，则pod状态是running。如果有readiness探针，则会在探测成功后，将容器准备状态设置为true，如果所有容器的准备状态都是true，kubelet就会将pod的状态更新为ready，并上报给apiserver。但如果readiness探针返回失败，容器也许会仍处于running状态，这种情况下，容器的状态确实是running，但是kubelet会将容器的准备状态设为false，该容器在运行但不能被其他容器访问。总的来说，container和pod都会出现running但不ready的情况。

## kubevirt是如何以容器形式运行虚拟机的？

kubevirt在k8s上以容器形式运行虚拟机，可以让虚拟机被k8s管控，享受到资源隔离、故障发现等机制。kubevirt也是借助了k8s CRD的形式，他自定义了一个资源叫VirtualMachineInstance，这个对象包含了虚拟机所需要的所有配置，但一般我们创建的是VM。

VM： 提供一些管理功能，比如启停虚拟机，

VMI ReplicaSet: 启动指定数量的虚拟机（这个我们没用，用户需要起几个我们就创建几个虚拟机）

kubevirt利用virt-controller和virt-api这两个集群组件与apiServer交互来创建我们需要的vm资源。在kubevirt的namespace下可以看到virt-controller和virt-api是独立出来的，virt-launcher是每个vm都有一个

- virt-api 因为crd模式管理vm 的pod，这个virt-api就是操作入口，可以启停vm
- virt-controller根据vmi的CRD去生成virt-launcher 的pod，监听CRD对象状态
- virt-handler是一个damonset 部署在每个节点上，监听vm实例的状态，根据状态变化进行响应，把它变回用户要求的状态；通知virt-launcher去使用它本地的libevirtd实例来启动vmi
- virt-launcher是在pod中每个vmi都有，他背后是libvirtd和qemu，这两个就是vm的虚拟化底层实现。launcher会管理vmi，pod生命周期结束时也会去通知vmi终止。
- 真正管理vm生命周期的是libvirtd。
- virtctl工具可以越过了virt-lancher的pod这层直接管理vm，启停一般用命令行比较好实现。但我们也使用了修改crd中的spec的running状态并重新apply去实现。pause的话是根据core api直接调用api暂停的。

创建流程：

k8s api应用vm crd对象

virt-controller监听vmi创建，生成pod spec文件，创建virt-launcher pods

virt-controller更新vmi crd状态

virt-handler监听到vm状态变更，通知virt-launcher去真正创建虚拟机，并负责生命周期管理。

- CDI Containerized Data Importer可以让PVC用于虚拟机（*PVC存在的问题）

cloud-init:在vmi launch时注入数据。vmi实例启动时执行。因为我们是自己搭的私有云，所以需要使用cloudinitnocloud方式，给虚拟机挂载一个文件系统到虚拟机里面，然后自己定制启动时的任务。会固定生成一个/dev/vdb的设备，把它挂载到/mnt下，有三个文件：

- meta-data 只有instance-id local-hostname
- network-config 网络信息，address 子网，网关，mac地址等
- user-data 用户自定义的一些信息，我们就是在userdata里写入对用户名密码的设置，以及设置DNS解析和一些基础的路由。
- 在networkData里设置静态ip，这个是从我们规划的网段中选择可用的ip绑定给这台vm。

存储：

- dataVolume在虚拟机启动流程中自动导入pvc，无需用户自己提前准备

目前使用的比较多的就是redhat的kubevirt和一个叫virtlet的。

## Calico和Cilium的区别以及各自的优缺点

- 工作的网络层不同，Calico在3层，cilium在7层（其实是工作在L3L4L7都可以的），所以cilium可以支持应用层协议。
- 实现方式不同：Calico基于BGP协议，无需额外封包；Cilium基于eBPF技术，这是一种内核态的技术，所以需要node的内核版本比较高。还可以做追踪。

## cgroup和namespace的区别

cgroup（Control Groups）和namespace（命名空间）都是Linux内核提供的一种机制，用于隔离进程和资源，但它们的目的和实现方式是不同的。

cgroup主要用于控制和限制进程组的资源使用，如CPU、内存、磁盘I/O等，它允许管理员将系统资源划分为若干个组，然后将进程分配到相应的组中，以达到对进程资源的管理和限制。cgroup的实现是通过在内核中创建一个层次结构的cgroup树，每个cgroup节点对应一个资源组，并设置相应的资源限制和控制参数。

namespace主要用于隔离进程的命名空间，包括PID、网络、挂载点、IPC等，每个命名空间提供了一个独立的环境，使得同一主机上的进程可以在不同的命名空间中运行，从而实现进程之间的隔离。namespace的实现是通过在内核中创建多个独立的命名空间，每个命名空间都有自己的PID、网络、文件系统等，进程在不同的命名空间中运行时只能访问相应的资源，无法访问其他命名空间中的资源。

总之，cgroup和namespace都是用于进程隔离和资源限制的机制，但是cgroup主要关注资源限制和控制，而namespace主要关注隔离和命名空间。

## 进程间通信

进程间通信（IPC）是指在操作系统中不同进程之间交换数据或协调活动的机制。常用的进程间通信方式包括以下几种：

1.管道（pipe）：管道是一种单向的通信方式，它可以将一个进程的输出和另一个进程的输入相连。管道通常用于父子进程之间或者兄弟进程之间的通信。

2.命名管道（named pipe）：命名管道是一种特殊的管道，它是通过文件系统中的一个特殊文件实现的。多个进程可以通过打开同一个文件来实现进程间通信。

3.信号（signal）：信号是一种异步通信方式，进程可以向另一个进程发送信号，接收进程可以选择处理或忽略信号。信号通常用于进程间的通知和中断处理。

4.共享内存（shared memory）：共享内存是一种高效的进程间通信方式，它允许多个进程访问同一块物理内存，从而避免了数据的复制。共享内存通常用于需要高效地传递大量数据的场景。

5.消息队列（message queue）：消息队列是一种基于消息的通信方式，进程可以向消息队列发送消息，接收进程可以从消息队列中读取消息。消息队列通常用于进程间的异步通信。

6.套接字（socket）：套接字是一种通用的进程间通信方式，它可以在本地或网络上进行通信。套接字可以用于各种进程间通信场景，如客户端-服务器、进程间的远程过程调用等

## 在k8s中使用nginx ingress controller和在物理机上使用nginx的区别和好处是什么

1. 功能区别：

Nginx Ingress Controller是Kubernetes中的一种扩展机制，它使用Ingress资源来管理HTTP和HTTPS路由，并将流量路由到Kubernetes集群内的Service资源。它支持多种负载均衡算法和流量控制策略，并可以与其他Kubernetes资源（如ConfigMap、Secret等）集成，提供高度可定制化的功能。

而在物理机上使用Nginx时，需要手动编写Nginx配置文件，并通过配置文件来管理HTTP和HTTPS路由，从而实现负载均衡和流量控制。这种方式需要更多的人工操作，并且不具备Kubernetes所提供的自动化和可扩展性。

1. 管理方式区别：

使用Nginx Ingress Controller可以通过Kubernetes API进行管理，从而实现自动化配置、自动扩缩容等功能，大大降低了运维的难度和成本。

而在物理机上使用Nginx时，需要手动管理Nginx的配置文件，需要更多的人力物力来维护，而且不如Kubernetes的自动化管理方式高效和稳定。

总之，使用Nginx Ingress Controller比在物理机上使用Nginx的好处在于它提供了更多的自动化和可定制化的功能，以及更高效、更稳定的管理方式。同时，Nginx Ingress Controller也是Kubernetes生态系统中的一部分，与其他Kubernetes组件（如Service、Deployment等）可以更好地集成，从而实现整个应用的自动化管理。