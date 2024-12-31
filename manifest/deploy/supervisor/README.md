# supervisor

## 安装 supervisor

基于 CentOS7 系统演示.

```sh
yum update -y
yum install epel-release -y
yum install supervisor -y

systemctl enable supervisord
systemctl start supervisord
systemctl status supervisord

echo_supervisord_conf > /etc/supervisord.conf
cat >> /etc/supervisord.conf <<EOF
[include]
files = supervisord.d/*.ini
EOF

mkdir -p /etc/supervisord.d
```

## 在发布机执行部署脚本

```sh
# 部署测试环境
$ manifest/deploy/supervisor/deploy.sh test

# 部署生产环境
$ manifest/deploy/supervisor/deploy.sh prod
```

## supervisorctl 常用命令

```sh
# 启动
$ sudo supervisorctl start devinggo

# 关闭(SIGTERM信号)
$ sudo supervisorctl stop devinggo

# 重启: 先关闭(SIGTERM信号), 再启动.
# NOTE: /etc/supervisord.*相关配置有变动, 重启具体某服务并不会生效
$ sudo supervisorctl restart devinggo

# 重启 supervisor 控制所有服务.
# NOTE: 当 /etc/supervisord.*相关配置有变动, 必须执行此命令才能加载生效
$ sudo supervisorctl reload
```
