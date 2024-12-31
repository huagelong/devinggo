<p align="center">
  <img src="web/system/public/logo.svg" width="120" />
</p>
<p align="center">
 <img src="https://svg.hamm.cn/badge.svg?key=License&value=Apache-2.0" />
  <img src="https://svg.hamm.cn/badge.svg?key=Version&value=1.0.0%20LTS" />
	<img src="https://img.shields.io/badge/goframe-2.7-green" alt="goframe">
	<img src="https://img.shields.io/badge/vue.js-vue3.4-green" alt="vue">
	<img src="https://img.shields.io/badge/vite-%3E5.1.4-yellow" alt="vite">
</p>

## 简介

`devinggo` 是一个基于 GoFrame v2 + Vue3 + Arco Design 开发的全栈前后端分离的后台管理系统。采用最新的前后端技术栈，专注于提供一个简洁、高效、安全的开发框架。

## 技术栈

- **后端**：GoFrame v2、MySQL、Redis
- **前端**：Vue3、Vite5、Arco Design、Pinia

## 核心功能

- [x] 用户管理：用户是系统操作者，该功能主要完成系统用户配置
- [x] 部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限
- [x] 岗位管理：配置系统用户所属担任职务
- [x] 菜单管理：配置系统菜单，操作权限，按钮权限标识等
- [x] 角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分
- [x] 字典管理：对系统中经常使用的一些较为固定的数据进行维护
- [x] 参数管理：对系统动态配置常用参数
- [x] 通知公告：系统通知公告信息发布维护
- [x] 操作日志：系统正常操作日志记录和查询
- [x] 登录日志：系统登录日志记录和查询
- [x] 代码生成：前后端代码的生成；支持CRUD下载
- [x] 系统监控：服务器、Redis、在线用户监控
- [x] 模块管理：管理系统当前所有模块
- [x] 定时任务：在线（添加、修改、删除)任务调度包含执行结果日志
- [x] 附件管理：管理系统当前所有附件
- [x] 支持多语言：支持多语言切换
- [x] 支持多租户：支持多租户管理
- [x] 支持多主题：支持多主题切换
- [x] 支持队列：支持基于redis的队列
- [x] 支持websocket：支持websocket通信,websocket支持分布式服务器
- [x] 支持docker部署

### 环境要求

- Go >= 1.18
- Node >= 16
- MySQL >= 5.7
- Redis >= 6.0

### 快速开始

> 以windows10为例

1. 安装make
```
#make 编译环境安装
@powershell -NoProfile -ExecutionPolicy Bypass -Command "iex ((new-object net.webclient).DownloadString('https://chocolatey.org/install.ps1'))" && SET PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin

choco install make
choco install sed
```
2. 安装golang
```
详见官网对应
```
3. 修改配置文件（mysql,redis配置）
- make config-init # 初始化配置文件
- `manifest/config/config.yaml` # 修改redis,mysql配置
- `hack/config.yaml` # 修改mysql配置,如果需要开发
```
4. 初始化
```
make build
```
5. 导入数据库
```
make install
```
6. 启动
```
make run
```
7. 访问服务端
```
http://localhost:8070
```