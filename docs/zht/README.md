# LessAPI-DuckDuckGo

[![GitHub](https://img.shields.io/github/license/lessapi-dev/lessapi-duckduckgo?style=for-the-badge)](https://github.com/lessapi-dev/lessapi-duckduckgo)
[![Docker](https://img.shields.io/docker/pulls/lessapi-dev/lessapi-duckduckgo?style=for-the-badge)](https://hub.docker.com/r/lessapi-dev/lessapi-duckduckgo)

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

(當前語言文檔由AI翻譯生成，如有疑問，請以簡體中文或英文文檔為準)

---

## 簡介

LessAPI-DuckDuckGo 是一個搜索引擎API服務。

基於 playwright 和 DuckDuckGo 搜索引擎，封裝後實現簡單的API接口。

簡單、輕量、可靠、Docker部署、易於維護。

## 狀態

> 專案正在實驗性地開發中，不建議在生產環境中使用。

## 使用

### OpenAPI標準文檔 (Swagger 3.0)

[OpenAPI標準文檔 (Swagger 3.0)](./../../lessapi-duckduckgo.openapi.json)

### 文本搜索 `GET /search/text`

**請求參數:**

- keyword: 搜索關鍵字(必填)
- region: 地區(選填)  wt-wt, us-en, uk-en, ru-ru, 等 預設值 wt-wt
- maxCount: 最大返回數量(選填)  預設值 20
- viaProxyUrl: 瀏覽器使用代理的地址(選填) 如 http://proxy.server:3000  預設值 空
  **請求示例:**

```shell
curl 'http://127.0.0.1:8080/search/text?keyword=hello&maxCount=2'
```

**響應示例:**

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

## 關於

本專案由 [廈門靜篤網絡科技有限公司](https://gentletld.cn) 開發和維護。

## 許可

[Apache-2.0](./../../LICENSE)
