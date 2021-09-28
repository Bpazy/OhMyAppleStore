# OhMyAppleStore
[![Go](https://github.com/Bpazy/OhMyAppleStore/actions/workflows/test.yml/badge.svg)](https://github.com/Bpazy/OhMyAppleStore/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Bpazy/OhMyAppleStore)](https://goreportcard.com/report/github.com/Bpazy/OhMyAppleStore)

我要抢 iPhone 自提，**本项目仅支持自提库存检测**。

## 预览
当检测到有 **自提库存** 时，发送钉钉消息:
![image](https://user-images.githubusercontent.com/9838749/134766917-0c2b7658-e706-42d0-8ef0-ac56745adf76.png)

## 使用方法
1. 创建配置文件 `$HOME/.oh-my-apple-store.yaml`，内容为:
```yaml
# 钉钉机器人配置，需要群机器人使用加签模式
dingtalk:
  token: 填入群机器人 access_token
  secret: 填入群机器人选择“加签”选项后展示的密钥

# 要检测自提的手机型号
phones:
  # 这是远峰蓝 256G
  - name: 远峰蓝色
    code: MLHC3CH/A

  # 这是石墨色 256G
  - name: 石墨色
    code: MLH83CH/A
```
型号可以在网站上选择型号后通过开发者工具获取 ![image](https://user-images.githubusercontent.com/9838749/135055720-9348531e-d09d-4715-9519-9e079b174b93.png)

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


