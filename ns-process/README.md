[此处](https://www.kancloud.cn/infoq/docker-source-code-analysis/80531)对容器创建过程的d描述很棒。

rootfs [下载地址](https://uk.images.linuxcontainers.org/images/centos/7/amd64/)

namespace 以及本文参考：[https://medium.com/@teddyking/namespaces-in-go-basics-e3f0fc1ff69a](https://medium.com/@teddyking/namespaces-in-go-basics-e3f0fc1ff69a)

名称空间 namespace 使你的进程不能干涉到对它不可视的东西，而能起到限制进程使用物理资源的内核特性是 CGroups。

## Let's go

命令`unshare`允许你运行一个子进程并使其从其父进程的 namespace 中脱离开来：

```bash
Usage:
 unshare [options] [<program> [<argument>...]]

Run a program with some namespaces unshared from the parent.

Options:
 -m, --mount[=<file>]      unshare mounts namespace
 -u, --uts[=<file>]        unshare UTS namespace (hostname etc)
 -i, --ipc[=<file>]        unshare System V IPC namespace
 -n, --net[=<file>]        unshare network namespace
 -p, --pid[=<file>]        unshare pid namespace
 -U, --user[=<file>]       unshare user namespace
 -C, --cgroup[=<file>]     unshare cgroup namespace
 -T, --time[=<file>]       unshare time namespace
```

以一个在新 UTS namespace 的 /bin/bash 进程中修改主机名的实验为例，其中的配置并不会影响到其父进程：

```bash
]$ sudo su
]# hostname
iyidongren
]# unshare -u /bin/bash
]# hostname new-hostname
]# hostname
new-hostname
]# exit
exit
]# hostname
iyidongren
```

当我们创建大部分的 namespace 时都需要 root 权限，下面是七个可用的`namespace`：

- Mount：隔离文件系统挂载点；
- UTS：隔离主机名和 domainname；
- IPC：隔离进程间通信资源；
- PID：隔离 PID 号空间；
- Network： 隔离网络接口；
- User：隔离 UID/GID 号空间；
- Cgroup：隔离 cgroup 根目录；

大部分的容器技术的实现都会用到上述中的 namespace 以确保能提供给两个分隔的容器进程最高的隔离性。

虽然`unshare`在简单的脚本文件中使用起来很棒，但却达不到想容器那样对 namespace d精细控制。`go`语言以成为容器实现的首选语言了，一部分原因是因为 Docker 就是用 go 实现的。

这文章的主要目的是理解如何用 go 语言来操作 Linux namespaces，为此我们还需要来构建一个名为`ns-process`的简单应用。

`ns-process`一开始很简单，将在新的 namespace 下创建`/bin/sh`程序。慢慢深入它将进化为更会有趣的东西：

> a program capable of creating unprivileged containers!

你现在还不理解 unprivileged 没关系，`ns-process`的代码都在原作者（Ed King）本人的 [Github 仓库](https://github.com/teddyking/ns-process)中。

```go
package main

import (
        "fmt"
        "os"
        "os/exec"
        "syscall"
)

func main() {
        cmd := exec.Command("/bin/sh")

        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        cmd.Env = []string{"PS1=-[ns-process]- # "}

        cmd.SysProcAttr = &syscall.SysProcAttr{
                Cloneflags: syscall.CLONE_NEWUTS,
        }

        if err := cmd.Run(); err != nil {
                fmt.Printf("Error running the /bin/sh command - %s\n", err)
                os.Exit(1)
        }
}
```

如你所见，我们就是简单地创建了一个`*exec.Cmd`对象，同时将标准输入输出错误都从调用进程转向了它，还为新进程设置了环境变量`PS1`仅仅是为了好区分 shell 界面罢了。

有趣的地方在`cmd.SysProcAttr`，不过在开始前我们还得向深挖掘那写组成 namespaces API 的底层系统调用。

## The namespaces API

在 man 帮助手册中查找 namespace 关键字，

```bash
man -k namespace
# whatis clone
```

其中有三个关键的系统调用：

- clone (2)       - create a child process
- setns (2)       - reassociate thread with a namespace
- unshare (2)      - disassociate parts of the process execution context

当在运行到`exec.Run()`时，`clone()`就会被调用，man 中有描述如下：

```bash
Both clone() and clone3() allow a flags bit mask that modifies their behavior and allows the caller to specify what  is shared between the calling process and the child process.

CLONE_NEWCGROUP (since Linux 4.6)
              Create the process in a new cgroup namespace. If this flag is not set, then (as with fork(2)) the process is created in the same cgroup namespaces as the calling proces
              
CLONE_NEWIPC (since Linux 2.6.19)
              If  CLONE_NEWIPC is set, then create the process in a new IPC namespace.  If this flag is not set, then (as with fork(2)), the process is created in the same IPC namespace as the calling process.
              
CLONE_NEWNET (since Linux 2.6.24)
CLONE_PID
CLONE_NEWUTS
CLONE_NEWUSER

CLONE_NEWNS (since Linux 2.4.19)
              If CLONE_NEWNS is set, the cloned child is started in a new mount namespace, initialized with a copy of the namespace  of  the  parent. 
```

接着时 go 中的`cmd.SysProcAttr`了，它允许我们来为`*exec.Cmd`设置属性；当我们指明了`Cloneflags`属性时，go 将会`CLONE_*`这些标志传递给系统调用`clone()`。

通过 root 权限运行这个程序时，你将陷入一个有着新`UTS namespace`的`/bin/sh`中。

```bash
go build ns_process.go
sudo ./ns_process
```

这此我们以如下方式来验证，它是否运行在新的`UTS namespace`中：

```bash
-[ns-process]- # readlink /proc/self/ns/uts
uts:[4026532851]
-[ns-process]- # exit
exit
➜ readlink /proc/self/ns/uts 
uts:[4026531838]
```

可见`ns-process`中的 UTS inode 号与外层的不同，便表明他们运行在了不同的`UTS namespace`中。

```go
cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWUSER,
	}
```

做出如上修改编译运行后，这次你将陷入到新的 Mount, UTS, IPC, PID, Network 和 User namespace 的`/bin/sh`进程当中。

> 当传递了 usr namespace 后，我们便不需要在通过 root 权限运行了。

虽然目前为止还挺好，但我们还是错过了一些初始化 namespace 时的配置操作：

- 我们请求了一个新的 mount namespace 但是还背负着主机的挂载和根文件系统；
- 我们请求了一个新的 PID namespace 但却没有挂载到一个新的`/proc`文件系统；
- 我们请求了一个新的 network namespace 但却没有在其中创建任何接口；
- 我们请求了一个新的 user namespace 但却没有提供一个 UID/GID 的映射。

## User namespace

尽管我们无需再以 root 权限运行`ns-process`了，但同时我们也丢掉了 root 身份：

```bash
➜ ./ns_process              
-[ns-process]- # whoami
nobody
```

身份的丢失代表着我们忽略了一些重要的配置，仅仅传递`CLONE_NEWUSER`标志是不够的。ID 映射和其与 user namespace 如何关联的内容太过庞大，我们仅需知道如何修复即可。

- User namespace 提供了 UID 和 GID 的隔离；
- 同一主机上随时可以有多个不同的 user namespace；
- 任一 Linux 进程都运行在所有 user namespace 中之一；
- User namespace 允许同一进程在不同 user namespace 中有不同的 UID 号；
- UID/GID mapping 提供了一种在两分隔的 user namespace 中映射 ID 的机制。

其核心实现细节就是，对两个不同的 user namespace 的映射。虽然`ns-process`进程在前一 user namespace 中有当前用户的权限，但当其运行在新 user namespace 后变成了`nobody`。

ID 的映射都是在`cmd.SysProcAttr`中实施的，如下的`SysProcIDMap`结构都可以在 go 的`syscall`包中找到。

```go
type SysProcIDMap struct {
        ContainerID int // Container ID.
        HostID      int // Host ID.
        Size        int // Size.
}
```

前两项的意思都很好理解，`size`决定了 ID 能映射的范围。再来更新一下我们的程序：

```go
cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}
```

这里我们创建了单个的 UID、GID 映射，我们将`ContainerID`设为了 0，同时`HostID`为当前用户的 user ID；也就是说我们将新 user namespace 中的 root ID 映射给了调用`ns-process`进程的用户。

```bash
➜ go build ./ns_process.go 
➜ ./ns_process            
-[ns-process]- # whoami
root
```

成功保留 root 身份。

## Reexec

这部分主要是来理解`reexec`包，这是 Docker 代码的一部分，下面的例子能很好地说明问题。

假设我们需要升级`ns-process`，需要一个随机生成的主机名在其 UTS namespace 中；并且处于安全考虑，我们需要在`/bin/sh`进程运行前就完成主机名的设置。毕竟我们不希望运行在`ns-process`中的进程会发现主机的名字。

go 语言自身并没有提供这样一种方法，因为 name spaces 是在为`*exec.Cmd`设置属性时创建的，这也是指定我们要运行的进程的位置：

```go
cmd := exec.Command("/bin/echo", "Process already running")
cmd.SysProcAttr = &syscall.SysProcAttr{
 Cloneflags: syscall.CLONE_NEWUTS,
}
cmd.Run()
```

一旦`cmd.Run()`被调用，便会开始克隆 name spaces 并且直接开始运行进程。这里没有钩子或其他什么办法来让我们能在 name space 创建后、进程运行前，运行一些代码。re-exec 正是为此而生。

来深入`reexec`包源码看看：

```go
// Register adds an initialization func under the specified name
func Register(name string, initializer func()) {
 # ...
}
```

首先我们的`Register`暴露出来了一种，注册任意函数并将其保留到内存当中的方法。我们将用它来在`ns-process`最处启动时，注册一些 “初始化 name space” 的函数。

```go
// Init is called as the first part of the exec process
// and returns true if an initialization function was called.
func Init() bool {
 # ...
}
```

接下来我们的`Init`，为我们提供了一种机制，该机制可判断一个进程在被 re-exec 后是否还在运行，并且对于正在运行的进程还可以再运行一个 Registered 函数。它通过检查`os.Args[0]`中的名字是否在先前的注册函数里来实现的。

```go
// Command returns *exec.Cmd which have Path as current binary.
// ...
func Command(args ...string) *exec.Cmd {
 return &exec.Cmd{
  Path: Self(),
  Args: args,
  SysProcAttr: &syscall.SysProcAttr{
   Pdeathsig: syscall.SIGTERM,
  },
 }
}
```

`Command`通过创建一个`*exec.Cmd`并且其中 Path 设置为`Self()`，来将所有绑定在一起；这就像是 Linux 上的`/proc/self/exe`机制。我们可以选择一个想要调用的注册函数，放在`reexec`的`args[0]`上即可。

> `/proc/self/exe`是一个指向正在运行的可执行文件路径的连接文件。

整合的第一步就是，创建一个函数并且用`reexec`注册。

```go
# Git repo: https://github.com/teddyking/ns-process
# Git tag: 3.0
# Filename: ns_process.go

# ...
func init() {
	reexec.Register("nsInitialisation", nsInitialisation)
	if reexec.Init() {
		os.Exit(0)
	}
}
# ...
```

这里发生了两件重要的事。第一，我们以 ”nsInitialisation“ 之名注册了`nsInitialisation`函数。过一会我们就加上它。第二件，我们调用了`reexec.Init()`接着当返回为真时`os.Exit(0)`退出；这是保证程序不陷入 exec 循环的重要条件。

```go
# ...
func nsInitialisation() {
	fmt.Printf("\n>> namespace setup code goes here <<\n\n")
	nsRun()
}

func nsRun() {
	cmd := exec.Command("/bin/sh")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = []string{"PS1=-[ns-process]- # "}

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running the /bin/sh command - %s\n", err)
		os.Exit(1)
	}
}
```

在这里我们添加的`nsInitialisation`就是一个简单的占位函数，但它将在我们真正需要配置 name spaces 时越来越重要；此时它只是调用了`nsRun()`。

仅剩下要修改的是`main()`了，将 main 对`/bin/sh`的直接调用改为通过对`reexec`和`nsInitialisation`的调用。

```go
func main() {
    cmd := reexec.Command("nsInitialisation")
    # ...
}
```

将`nsInitialisation`指明为`Command`的第一个参数，实质上就是让`reexec`运行`/proc/self/exec`时带将`os.Args[0]` 设置为`nsInitialisation`。最后，程序一旦被 reexec 了，`Init`将会探测已注册的函数并且真正运行它。

- [参考资料](https://pkg.go.dev/github.com/docker/docker/pkg/reexec?tab=overview)
- [下载镜像](https://github.com/goproxy/goproxy.cn/blob/master/README.zh-CN.md)

## Mount

这部分的内容也是容器实现中的一大特性，有能力在同一宿主机上运行不同发行版的 Linux。





## 问题记录

go get 无法下载问题解决，镜像安装：https://www.sunzhongwei.com/problem-of-domestic-go-get-unable-to-download 注意查看 help 的帮助信息，其中指明了要指明版本。

 终于解决了这个报错，反思记录：

- 认真观察报错提示，并且多去尝试搜索几个关键词而不是 抓着一个不放；
- 这种问题明显是由欠缺 go 工程管理的学习所致，生成的 go.mod 文件就是很好的例子，之前见过很多次但都不知道是什么意思。

- 如果你是使用go mod 管理依赖，首先检查：项目根目录有没有go.mod文件，前为引用，何为 go 中的 mod？
- [慢慢来理解吧](https://blog.csdn.net/fly910905/article/details/104297446#Go%20modules%E4%BD%BF%E7%94%A8%E6%AD%A5%E9%AA%A4%EF%BC%9A)



为什么需要执行`go mod init`？

为什么`GOPATH`变量为空时，`go mod init`命令会执行失败？

为什么`.profile`文件中的命令在开机时不会自动执行？它与`.bashrc`的区别又是什么？

`.profile`中会存一些不跟明确 shell 有关的配置，好比像是环境变量，但如果测试发现你的 zsh 不会自动读取该文件，可将`source $/HOME/.profile`行加入你的`$HOME/.zshrc`中。