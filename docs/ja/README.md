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

(この言語のドキュメントはAIによって翻訳されています。不明な点があれば、簡体字中国語または英語のドキュメントを参照してください)

---

## はじめに

LessAPI-DuckDuckGo は、搜索引擎APIサービスです。
playwright と DuckDuckGo 検索エンジンを基に、単純なAPIインターフェースを実装しています。
シンプルで軽量、信頼性が高く、Dockerによるデプロイメントと保守が容易です。

## ステータス

> 実験的に開発中であり、本番環境での使用は推奨されていません。

## 使い方

### OpenAPI標準ドキュメント (Swagger 3.0)

[OpenAPI標準ドキュメント (Swagger 3.0)](../../resource/openapi.json)

### テキスト検索 `GET /search/text`

**リクエストパラメータ:**

- keyword: 検索キーワード(必須)
- region: 地域(任意)  wt-wt, us-en, uk-en, ru-ru, など。デフォルト値は wt-wt
- maxCount: 最大返信数(任意)  デフォルト値は 20
- viaProxyUrl: ブラウザが使用するプロキシのアドレス(任意)  例: http://proxy.server:3000  デフォルト値は空

**リクエスト例:**

```shell
curl 'http://127.0.0.1:8080/search/text?keyword=hello&maxCount=2'
```

**レスポンス例:**

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

## ライセンス

[MIT](./../../LICENSE)
