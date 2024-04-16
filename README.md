# LessAPI-DuckDuckGo

[![GitHub](https://img.shields.io/github/license/lessapi-dev/lessapi-duckduckgo?style=for-the-badge)](https://github.com/lessapi-dev/lessapi-duckduckgo)
[![Docker](https://img.shields.io/docker/pulls/lessapi/lessapi-duckduckgo?style=for-the-badge)](https://hub.docker.com/r/lessapi-dev/lessapi-duckduckgo)

[English](./README.md) |
[简体中文](./docs/zhs/README.md) |
[繁體中文](./docs/zht/README.md) |
[日本語](./docs/ja/README.md) |
[한국어](./docs/ko/README.md) |
[Русский](./docs/ru/README.md) |
[Deutsch](./docs/de/README.md) |
[Français](./docs/fr/README.md) |
[Español](./docs/es/README.md) |
[Italiano](./docs/it/README.md) |
[Português](./docs/pt/README.md)

## Introduction

LessAPI-DuckDuckGo is an API service for a search engine.

Based on Playwright and DuckDuckGo's search engine, it encapsulates to provide simple API interfaces.

Simple, lightweight, reliable, Docker deployable, easy to maintain.

## Status

> Experimentally under development and not recommended for use in production environments.

## Usage

### OpenAPI Standard Documentation (Swagger 3.0)

[OpenAPI Standard Documentation (Swagger 3.0)](./lessapi-duckduckgo.openapi.json)

### Text Search `GET /search/text`

**Request Parameters:**

- keyword: Search keyword (required)
- region: Region (optional) such as wt-wt, us-en, uk-en, ru-ru, etc. Default is wt-wt
- maxCount: Maximum number of results returned (optional) Default is 20
- viaProxyUrl: The address of the proxy used by the browser (optional) e.g., http://proxy.server:3000 Default is empty

**Request Example:**

```shell
curl 'http://127.0.0.1:8080/search/text?keyword=hello&maxCount=2'
```

**Response Example:**

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

## Security

[![Security Status](https://www.murphysec.com/platform3/v31/badge/1779906127272730624.svg)](https://www.murphysec.com/console/report/1778449242088529920/1779906127272730624)

## About

This project is developed and maintained by [Xiamen Jingdu Network Technology Co., Ltd.](https://gentletld.cn).

## License

[MIT](./LICENSE)
