# LessAPI-DuckDuckGo

[![GitHub](https://img.shields.io/github/license/lessapidev/lessapi-duckduckgo?style=for-the-badge)](https://github.com/username/lessapi-duckduckgo)
[![Docker](https://img.shields.io/docker/pulls/lessapidev/lessapi-duckduckgo?style=for-the-badge)](https://hub.docker.com/r/lessapidev/lessapi-duckduckgo)

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

(当前语言文档由AI翻译生成，如有疑问，请以简体中文或英文文档为准)

---

## 简介

LessAPI-DuckDuckGo 是一个搜索引擎API服务。

基于 playwright 和 DuckDuckGo 搜索引擎，封装后实现简单的API接口。

简单、轻量、可靠、Docker部署、易于维护。

## 状态

> 实验性地开发中，不建议在生产环境中使用。

## 使用

### OpenAPI标准文档 (Swagger 3.0)

[OpenAPI标准文档 (Swagger 3.0)](./../../lessapi-duckduckgo.openapi.json)

### 文本搜索 `GET /search/text`

**请求参数:**

- keyword: 搜索关键字(必填)
- region: 地区(选填)  wt-wt, us-en, uk-en, ru-ru, 等 默认值 wt-wt
- maxCount: 最大返回数量(选填)  默认值 20
- viaProxyUrl: 浏览器使用代理的地址(选填) 如 http://proxy.server:3000  默认值 空

**请求示例:**

```shell
curl 'http://127.0.0.1:8080/search/text?keyword=hello&maxCount=2'
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

## 关于

本项目由 [厦门静笃网络科技有限公司](https://gentletld.cn) 开发和维护。

## 许可

[Apache-2.0](./../../LICENSE)
