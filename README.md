# LessAPI-DuckDuckGo

[![GitHub](https://img.shields.io/github/license/lessapi-dev/lessapi-duckduckgo?style=for-the-badge)](https://github.com/lessapi-dev/lessapi-duckduckgo)
[![Docker](https://img.shields.io/docker/pulls/lessapi/lessapi-duckduckgo?style=for-the-badge)](https://hub.docker.com/r/lessapi/lessapi-duckduckgo)

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

Large Language Model (LLM) Friendly. Support Plain Text Response.

## Status

> Experimentally under development and not recommended for use in production environments.

## Deployment

One command is all it takes to deploy the service to port 8080 using Docker.

```shell
docker run -d -p 8080:8080 --restart=unless-stopped --name lessapi-duckduckgo lessapi/lessapi-duckduckgo:v0.0.2
```

## Usage

### OpenAPI Standard Documentation (Swagger 3.0)

[OpenAPI Standard Documentation (Swagger 3.0)](resource/openapi.json)

### Text Search `GET /search/text`

**Request Parameters:**

- keyword: Search keyword (required)
- region: Region (optional) such as en-US, fr-FR, zh-CN, ru-RU, etc. Default is en-US
- maxCount: Maximum number of results returned. (optional) Default is 20
- viaProxyUrl: The address of the proxy used by the browser. e.g., http://proxy.server:3000 (optional) Default is empty
- llmStyle: Whether to use Large Language Model (LLM) Friendly style response. e.g. 1, 0 (optional) Default is 0

**Request Example:**

```shell
curl 'http://127.0.0.1:8080/search/text?keyword=lessapi&maxCount=10'
```

```shell
curl 'http://127.0.0.1:8080/search/text?keyword=lessapi&maxCount=99&viaProxyUrl=http://proxy.server:3000'
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

## Advanced

### Use Environment Variables

- **LESSAPI_DEFAULT_LANGUAGE**: (optional) Default language, such as en-US, fr-FR, zh-CN, ru-RU, etc. Default is en-US
- **LESSAPI_DEFAULT_VIA_PROXY_URL**: (optional) The address of the proxy used by the browser,
  e.g., http://proxy.server:3000 Default is empty

## Security

[![Security Status](https://www.murphysec.com/platform3/v31/badge/1779906127272730624.svg)](https://www.murphysec.com/console/report/1778449242088529920/1779906127272730624)

## License

[MIT](./LICENSE)
