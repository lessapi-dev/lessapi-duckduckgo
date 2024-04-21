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

(Текущий язык документа был сгенерирован с помощью AI, в случае сомнений, пожалуйста, смотрите на документы на简体中文
или английском языке)

---

## Введение

LessAPI-DuckDuckGo является сервисом API поисковой системы.
Основан на playwright и DuckDuckGo поисковой системы, после упаковки реализует простой API интерфейс.
Простой, легкий, надежный, установка Docker, легко обслуживать.

## Состояние

> Экспериментально разрабатывается, не рекомендуется использовать в производственной среде.

## Использование

### OpenAPI стандартный документ (Swagger 3.0)

[OpenAPI стандартный документ (Swagger 3.0)](../../resource/openapi.json)

### Текстовый поиск `GET /search/text`

**Параметры запроса:**

- keyword: поисковые ключевые слова (обязательно)
- region: регион (опционально)  wt-wt, us-en, uk-en, ru-ru и т.д. Значение по умолчанию wt-wt
- maxCount: максимальное количество возвращаемых результатов (опционально)  Значение по умолчанию 20
- viaProxyUrl: адрес прокси для браузера (опционально)  Например, http://proxy.server:3000  Значение по умолчанию пустое

**Пример запроса:**

```shell
curl 'http://127.0.0.1:8080/search/text?keyword=hello&maxCount=2'
```

**Пример ответа:**

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

## Лицензия

[MIT](./../../LICENSE)
