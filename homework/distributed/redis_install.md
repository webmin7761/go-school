# redis安装设置

## 编译

1. 升级gcc

    ```sh
    yum -y install centos-release-scl
    yum -y install devtoolset-9-gcc devtoolset-9-gcc-c++ devtoolset-9-binutils

    scl enable devtoolset-9 bash

    #scl命令启用只是临时的，推出xshell或者重启就会恢复到原来的gcc版本。如果要长期生效的话
    echo "source /opt/rh/devtoolset-9/enable" >>/etc/profile
    ```

2. 下载redis源码

    ```sh
    wget http://download.redis.io/releases/redis-6.0.6.tar.gz
    tar xzf redis-6.0.6.tar.gz
    ```

3. 编译jemalloc等依赖

    ```sh
    cd deps
    make hiredis lua jemalloc linenoise
    ```

4. 编译不同内存分配器的redis

    - jemalloc - `make distclean && make MALLOC=jemalloc`
    - malloc - `make distclean && make MALLOC=libc`

## redis.conf

1. 关闭RBD&AOF

    - 关闭RBD - 注释掉所有save xx
    - 关闭AOF - appendonly no

2. 不同内存分配器使用不同的端口

    - jemalloc 6379
    - libc 6380

- `./redis-server redis.conf`

## 调整linux内核参数

```sh
vi /etc/sysctl.conf
vm.overcommit_memory=1
net.core.somaxconn=1024
sysctl -p    

#检查是否启动THP
cat /sys/kernel/mm/transparent_hugepage/enabled
[always] madvise never

#临时调整，重启失效
echo never > /sys/kernel/mm/transparent_hugepage/enabled

#永久调整
vi /etc/default/grub
#在GRUB_CMDLINE_LINUX中添加transparent_hugepage=never
GRUB_TIMEOUT=5
GRUB_DEFAULT=saved
GRUB_DISABLE_SUBMENU=true
GRUB_TERMINAL_OUTPUT="console"
GRUB_CMDLINE_LINUX="nomodeset crashkernel=auto rd.lvm.lv=vg_os/lv_root rd.lvm.lv=vg_os/lv_swap rhgb quiet transparent_hugepage=never"
GRUB_DISABLE_RECOVERY="true"

#不同的BIOS需要生成在不同的地方，可以通过检查目录是否存在判断BIOS的类型
#On BIOS-based machines
grub2-mkconfig -o /boot/grub2/grub.cfg
```
