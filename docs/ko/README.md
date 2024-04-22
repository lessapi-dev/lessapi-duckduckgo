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

(LessAPI-DuckDuckGo 프로젝트의 한국어 문서는 AI가 번역한 것으로, 문의 사항이 있을 경우 중국어 간체나 영어 문서를 참고하십시오)

---

## 소개

LessAPI-DuckDuckGo는 DuckDuckGo 검색 엔진을 기반으로 하는 간단한 API 인터페이스를 제공하는 검색 엔진 API 서비스입니다. playwright를 사용하여 감싸져 있으며, 간단하고 가벼우며
신뢰度高며 Docker로 배포되어서 쉽게 유지 관리할 수 있습니다.

## 상태

> 실험적으로 개발 중입니다. 프로덕션 환경에서 사용하지 마십시오.

## 사용법

### OpenAPI 표준 문서 (Swagger 3.0)

[OpenAPI 표준 문서 (Swagger 3.0)](../../resource/openapi.json)

### 텍스트 검색 `GET /search/text`

**요청 매개 변수:**

- keyword: 검색 키워드(필수)
- region: 지역(선택)  en-US, fr-FR, zh-CN, ru-RU, 등 기본값 en-US
- maxCount: 최대 반환 수(선택)  기본값 20
- viaProxyUrl: 브라우저가 프록시를 사용하는 주소(선택)  예: http://proxy.server:3000  기본값: 빈 문자열

**요청 예시:**

```shell
curl 'http://127.0.0.1:8080/search/text?keyword=hello&maxCount=2'
```

**응답 예시:**

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

## 라이선스

[MIT](./../../LICENSE)
