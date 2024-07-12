#!/bin/bash

# 定义颜色常量
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# 使用颜色输出函数
color_echo() {
    echo -e "${CYAN}$1${NC}"
}
# 获取当前时间
current_time=$(date +"%Y年%m月%d日%H时%M分%S秒")
info="[info] [$current_time]"
go run generate.go
color_echo "${info} 插件注册文件生成成功!👍"
color_echo "${info} 开始启动qqbot🚀"
go run ../server/server.go