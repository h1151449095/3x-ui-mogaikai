# 3X-UI (h1151449095 Fork)

**3X-UI** 是一个基于网页的 Xray-core 控制面板，由 [h1151449095](https://github.com/h1151449095) 维护。

本仓库基于 [MHSanaei/3x-ui](https://github.com/MHSanaei/3x-ui) 二次开发，提供一键协议模板、中转、多服务器下发、批量管理等功能。

> 本项目仅供个人学习与通信使用，请勿用于任何非法用途。

---

## 快速开始

`ash
# 安装脚本（Linux VPS）
bash <(curl -Ls https://raw.githubusercontent.com/h1151449095/3x-ui-h1151449095/main/install.sh)
`

装完后输入 x-ui 打开管理菜单。

---

## 功能特性

- 一键协议模板（VLESS Reality、Trojan、VMess、Hysteria2）
- 中转（落地分流）
- 多服务器部署
- 批量删除与二维码导入
- SQLite / PostgreSQL 双数据库
- Docker 部署支持
- Telegram 机器人通知

---

## 数据库选项

- **SQLite**（默认）：单文件，零配置
- **PostgreSQL**：适合大规模部署

---

## 开源协议

本项目基于 [MHSanaei/3x-ui](https://github.com/MHSanaei/3x-ui)（GPL-3.0）二次开发，遵循 **GPL-3.0** 协议。
