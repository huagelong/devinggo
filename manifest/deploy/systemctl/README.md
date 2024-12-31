# systemctl

## 根据实际情况修改 deploy.sh 部署相关变量

```sh
project_name="devinggo"
app_name=${project_name}
systemd_service="${app_name}.service"

# 不要使用个人账号部署和运行服务, 建议建立独立的公共账号部署和运行应用.
deploy_user="vagrant"
# NOTE: 1. 请提前配置发布机到目标服务器之间的ssh key信任:
#    1) ssh-keygen -t rsa -b 4096 -C "vagrant@$(hostname)"
#    2) ssh-copy-id -i ~/.ssh/id_rsa.pub vagrant@$(deploy_server)
# 2. 目标机器上解决sudo执行命令提示输入密码的问题:
#    执行visudo, 末尾添加一行: vagrant ALL=(ALL:ALL) NOPASSWD:ALL
deploy_server="devinggo.sre.im"
deploy_dir=/opt/$project_name
```

## 在发布机执行部署脚本 deploy.sh

```sh
# 部署测试环境
$ manifest/deploy/systemctl/deploy.sh test

# 部署生产环境
$ manifest/deploy/systemctl/deploy.sh prod
```

## service 文件配置

```ini
[Unit]
Description=devinggo production environment
Documentation=
After=network.target
Wants=

[Service]
Type=simple
# 运行应用的账号.
User=vagrant
# 程序启动用到的环境变量.
Environment=GF_GERROR_BRIEF=true GF_GCFG_FILE=config.yaml
# 程序所在目录.
WorkingDirectory=/opt/devinggo
# 使用 sh -c "" 这种方式启动, 避免将stdout和stderr输出到/var/log/message中,
# 统一输出到指定的文件中.
ExecStart=/bin/sh -c '/opt/devinggo/devinggo >> devinggo.log 2>&1'
# 发送SIGTERM信号给所有控制的进程.
KillMode=ctrl-group
# 只能通过 systemctl stop 来关闭服务, 直接kill进程服务还会自动重启.
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
```

## systemctl 常用操作命令

```sh
# devinggo.service 配置有变动的时候, 需要重新加载使生效
$ sudo systemctl daemon-reload

# 启动
$ sudo systemctl start devinggo

# 关闭: 发送 SIGTERM 信号给主(sh)和子进程(devinggo),
# devinggo程序可通过捕捉SIGTERM信号来实现优雅关闭.
$ sudo systemctl stop devinggo

# 重启: 先关闭(SIGTERM), 再启动
$ sudo systemctl restart devinggo
```
