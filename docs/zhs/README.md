# LessAPI-DuckDuckGo

[![GitHub](https://img.shields.io/github/license/lessapi-dev/lessapi-duckduckgo?style=for-the-badge)](https://github.com/lessapi-dev/lessapi-duckduckgo)
[![Docker](https://img.shields.io/docker/pulls/lessapi/lessapi-duckduckgo?style=for-the-badge)](https://hub.docker.com/r/lessapi/lessapi-duckduckgo)

[English](./../../README.md) |
[简体中文](./../zhs/README.md) |
[繁體中文](./../zht/README.md) |
[日本語](./../ja/README.md) |
[한국어](./../ko/README.md) |
[Русский](./../ru/README.md) |
[Deutsch](./../de/README.md) |
[Français](./../fr/README.md) |
[Español](./../es/README.md) |
[Italiano](./../it/README.md) |
[Português](./../pt/README.md)


---

## 简介

LessAPI-DuckDuckGo 是一个搜索引擎API服务。

基于 playwright 和 DuckDuckGo 搜索引擎，封装后实现简单的API接口。

简单、轻量、可靠、Docker部署、易于维护。

大型语言模型(LLM)友好。支持纯文本响应。

## 状态

> 实验性地开发中，不建议在生产环境中使用。

## 部署

使用Docker只需要一个命令即可将服务部署到8080端口。

```shell
docker run -d -p 8080:8080 --restart=unless-stopped --name lessapi-duckduckgo lessapi/lessapi-duckduckgo:v0.0.2
```

## 使用

### OpenAPI标准文档 (Swagger 3.0)

[OpenAPI标准文档 (Swagger 3.0)](../../resource/openapi.json)

### 文本搜索 `GET /search/text`

**请求参数:**

- keyword: 搜索关键字(必填)
- region: 地区(选填)  en-US, fr-FR, zh-CN, ru-RU, 等 默认值 en-US
- maxCount: 最大返回数量(选填)  默认值 20
- viaProxyUrl: 浏览器使用代理的地址(选填) 如 http://proxy.server:3000  默认值 空
- llmStyle: 是否使用大型语言模型(LLM)友好风格响应(选填)  1, 0 默认值 0

**请求示例:**

```shell
curl 'http://127.0.0.1:8080/search/text?keyword=lessapi&maxCount=10'
```

```shell
curl 'http://127.0.0.1:8080/search/text?keyword=lessapi&maxCount=99&viaProxyUrl=http://proxy.server:3000'
```

**响应示例:**

```json
{
  "code": "success",
  "data": {
    "results": [
      {
        "order": 1,
        "title": "Adele - Hello (Official Music Video) - YouTube",
        "url": "https://www.youtube.com/watch?v=YQHsXMglC9A",
        "description": "Listen to \"Easy On Me\" here: http://Adele.lnk.to/EOMPre-order Adele's new album \"30\" before its release on November 19: https://www.adele.comShop the \"Adele..."
      },
      {
        "order": 2,
        "title": "Hello Definition & Meaning - Merriam-Webster",
        "url": "https://www.merriam-webster.com/dictionary/hello",
        "description": "Learn the origin, usage, and synonyms of the word hello, an expression or gesture of greeting. See examples of hello in sentences and related words from the dictionary."
      }
    ]
  }
}

```

## 安全

[![Security Status](https://www.murphysec.com/platform3/v31/badge/1779906127272730624.svg)](https://www.murphysec.com/console/report/1778449242088529920/1779906127272730624)

## 许可

[MIT](./../../LICENSE)
