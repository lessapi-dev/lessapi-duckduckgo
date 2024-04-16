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

(Este documento actualmente está siendo traducido por AI, por favor, consulta el documento en chino simplificado o
inglés en caso de duda)


---

## Introducción

LessAPI-DuckDuckGo es un servicio de API de motor de búsqueda.
Basado en playwright y DuckDuckGo, se ha encapsulado para implementar una interfaz de API simple.
Es simple, ligero, confiable, se puede desplegar con Docker y fácil de mantener.

## Estado

> Actualmente en desarrollo experimental, no se recomienda usar en entornos de producción.

## Uso

### Documento estándar OpenAPI (Swagger 3.0)

[Documento estándar OpenAPI (Swagger 3.0)](./../../lessapi-duckduckgo.openapi.json)

### Búsqueda de texto `GET /search/text`

**Parámetros de solicitud:**

- keyword: Palabra clave de búsqueda (obligatorio)
- region: Región (opcional) como wt-wt, us-en, uk-en, ru-ru, etc. El valor predeterminado es wt-wt
- maxCount: Número máximo de resultados devueltos (opcional) El valor predeterminado es 20
- viaProxyUrl: Dirección del proxy utilizado por el navegador (opcional) como http://proxy.server:3000 El valor
  predeterminado es vacío

**Ejemplo de solicitud:**

```shell
curl 'http://127.0.0.1:8080/search/text?keyword=hello&maxCount=2'
```

**Ejemplo de respuesta:**

```json
{
  "code": "success",
  "data": {
    "results": [
      {
        "order": 1,
        "title": "Adele - Hello (Official Music Video) - YouTube",
        "url": "https://www.youtube.com/watch?v=YQHsXMglC9A",
        "description": "Escucha \"Easy On Me\" aquí: http://Adele.lnk.to/EOMPre-order Adele's new album \"30\" before its release on November 19: https://www.adele.comShop the \"Adele..."
      },
      {
        "order": 2,
        "title": "Hello Definition & Meaning - Merriam-Webster",
        "url": "https://www.merriam-webster.com/dictionary/hello",
        "description": "Aprende el origen, uso y sinónimos de la palabra hello, una expresión o gesto de saludo. Ver ejemplos de hello en frases y palabras relacionadas del diccionario."
      }
    ]
  }
}
```

## Licencia

[MIT](./../../LICENSE)
