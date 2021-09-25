# OhMyAppleStore
[![Go](https://github.com/Bpazy/OhMyAppleStore/actions/workflows/test.yml/badge.svg)](https://github.com/Bpazy/OhMyAppleStore/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Bpazy/OhMyAppleStore)](https://goreportcard.com/report/github.com/Bpazy/OhMyAppleStore)

我要抢 iPhone 自提

## 预览
当检测到有 **自提库存** 时，发送钉钉消息:
![image](https://user-images.githubusercontent.com/9838749/134766917-0c2b7658-e706-42d0-8ef0-ac56745adf76.png)

## 使用方法
1. 创建配置文件 `$HOME/.oh-my-apple-store.yaml`，内容为:
```yaml
# 钉钉机器人配置，需要群机器人使用加签模式
dingtalk:
  token: 
  secret: 

# 要检测自提的手机型号
phones:
  - name: 远峰蓝色
    code: MLHC3CH/A
  - name: 石墨色
    code: MLH83CH/A
```

2. 下载 `.oh-my-apple-store` 并在该程序目录新建 `cookies.json` 文件，利用 `EditThsiCookie` 插件导出 cookies，形如:
```json
[{
    "domain": ".apple.com.cn",
    "expirationDate": 1663954303.104427,
    "hostOnly": false,
    "httpOnly": true,
    "name": "acn01",
    "path": "/",
    "sameSite": "unspecified",
    "secure": true,
    "session": false,
    "storeId": "0",
    "value": "",
    "id": 1
}]
```

3. 运行: `.oh-my-apple-store loop`


