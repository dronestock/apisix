FROM storezhang/alpine:3.20.0


LABEL author="storezhang<华寅>" \
email="storezhang@gmail.com" \
qq="160290688" \
wechat="storezhang" \
description="Drone持续集成系统Apisix网关插件，支持的功能有：1、Protobuf文件上传"


# 复制文件
COPY apisix /bin


RUN set -ex \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/apisix \
    \
    \
    \
    && rm -rf /var/cache/apk/*


# 执行命令
ENTRYPOINT /bin/apisix
