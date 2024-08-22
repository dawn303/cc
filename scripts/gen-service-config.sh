#!/usr/bin/env bash

CC_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source "${CC_ROOT}/scripts/common.sh"

if [ $# -ne 4 ];then
    cc::log::error "Usage: gen-service-config.sh SERVICE_NAME ENV_FILE TEMPLATE_FILE OUTPUT_DIR"
    exit 1
fi

export SERVICE_NAME=$1
ENV_FILE=$2
TEMPLATE_FILE=$3
OUTPUT_DIR=$4

if [ ! -d ${OUTPUT_DIR} ];then
  mkdir -p ${OUTPUT_DIR}
fi

if [[ " ${CC_ALL_COMPONENTS[*]} " != *" ${SERVICE_NAME} "* ]]; then
  exit 0
fi

source ${ENV_FILE}

suffix=$(echo $TEMPLATE_FILE | awk -F'.' '{print $NF}')
${CC_ROOT}/scripts/gen-config.sh ${ENV_FILE} ${TEMPLATE_FILE} > ${OUTPUT_DIR}/${SERVICE_NAME}.${suffix}

# 在Mac本地开发时，将配置复制到用户主目录用于测试
if [[ "$OSTYPE" = darwin* ]]; then
    cp -r ${OUTPUT_DIR}/${SERVICE_NAME}.${suffix} ${HOME}/.cc/
fi
