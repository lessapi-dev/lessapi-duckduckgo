# LessAPI-DuckDuckGo

[![GitHub](https://img.shields.io/github/license/lessapi-dev/lessapi-duckduckgo?style=for-the-badge)](https://github.com/lessapi-dev/lessapi-duckduckgo)
[![Docker](https://img.shields.io/docker/pulls/lessapi/lessapi-duckduckgo?style=for-the-badge)](https://hub.docker.com/r/lessapi-dev/lessapi-duckduckgo)

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

(Este documento em idioma é gerado por AI, caso haja dúvidas, por favor, consulte o documento em Chinês Simplificado ou
Inglês)

---

## Introdução

LessAPI-DuckDuckGo é um serviço de API de busca.
Baseado no playwright e no mecanismo de busca DuckDuckGo, é encapsulado para implementar uma interface de API simples.
Simples, leve, confiável, implantado com Docker, fácil de manter.

## Status

> Em desenvolvimento experimental, não recomendado para uso em ambientes de produção.

## Uso

### Documento padrão OpenAPI (Swagger 3.0)

[Documento padrão OpenAPI (Swagger 3.0)](./../../lessapi-duckduckgo.openapi.json)

### Pesquisa de texto `GET /search/text`

**Parâmetros de solicitação:**

- keyword: Palavra-chave de pesquisa (obrigatório)
- region: Região (opcional) como wt-wt, us-en, uk-en, ru-ru, etc. O valor padrão é wt-wt
- maxCount: Número máximo de resultados retornados (opcional) O valor padrão é 20
- viaProxyUrl: Endereço do proxy usado pelo navegador (opcional) como http://proxy.server:3000 O valor padrão é vazio

**Exemplo de solicitação:**

```shell
curl 'http://127.0.0.1:8080/search/text?keyword=hello&maxCount=2'
```

**Exemplo de resposta:**

```json
{
  "code": "success",
  "data": {
    "results": [
      {
        "order": 1,
        "title": "Adele - Hello (Official Music Video) - YouTube",
        "url": "https://www.youtube.com/watch?v=YQHsXMglC9A",
        "description": "Escute \"Easy On Me\" aqui: http://Adele.lnk.to/EOMPre-order Adele's new album \"30\" before its release on November 19: https://www.adele.comShop the \"Adele..."
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

## Sobre

Este projeto é desenvolvido e mantido pela [厦门静笃网络科技有限公司](https://gentletld.cn).

## Licença

[MIT](./../../LICENSE)
