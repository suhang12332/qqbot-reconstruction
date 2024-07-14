#!/bin/bash

# 获取当前脚本的绝对路径
current_script=$(readlink -f "$0")

# 获取当前行号
current_line=$LINENO

# 定义颜色常量
CYAN='\033[0;32m'
NC='\033[0m' # No Color

# 使用颜色输出函数
color_echo() {
    echo -en "${CYAN}$1${NC}"
}
color_echon() {
    echo -e "${CYAN}$1${NC}"
}
# 获取当前时间
current_time=$(date +"%Y年%m月%d日%H时%M分%S秒")

info="[$current_time][INFO]:"
go run generate.go
color_echo "${info} "
echo -n "${current_script}:${current_line}"
color_echon " 插件注册文件生成成功!👍"
color_echo "${info} "
echo -n "${current_script}:${current_line}"
color_echon " 开始启动qqbot🚀"
go run ../server/server.go