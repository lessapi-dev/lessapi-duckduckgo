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

(questo documento è stato tradotto da un traduttore AI; in caso di dubbi, fare riferimento ai documenti in cinese
semplificato o in inglese)

---

## Introduzione

LessAPI-DuckDuckGo è un servizio API di ricerca basato sulla libreria Playwright e sull'indicizzatore DuckDuckGo, che
offre un'interfaccia API semplice, leggera e affidabile, ed è facilmente mantenibile grazie al supporto Docker.

## Stato

> In fase sperimentale di sviluppo, non consigliato per l'uso in produzione.

## Utilizzo

### Documento OpenAPI standard (Swagger 3.0)

[Documento OpenAPI standard (Swagger 3.0)](../../resource/openapi.json)

### Ricerca testuale `GET /search/text`

**Parametri della richiesta:**

- keyword: parola chiave da cercare (obbligatorio)
- region: regione (facoltativo) es. wt-wt, us-en, uk-en, ru-ru, ecc. Il valore predefinito è wt-wt
- maxCount: numero massimo di risultati restituiti (facoltativo) Il valore predefinito è 20
- viaProxyUrl: indirizzo del proxy da utilizzare con il browser (facoltativo) ad esempio http://proxy.server:3000 Il
  valore predefinito è vuoto

**Esempio di richiesta:**

```shell
curl 'http://127.0.0.1:8080/search/text?keyword=hello&maxCount=2'
```

**Esempio di risposta:**

```json
{
  "code": "success",
  "data": {
    "results": [
      {
        "order": 1,
        "title": "Adele - Hello (Official Music Video) - YouTube",
        "url": "https://www.youtube.com/watch?v=YQHsXMglC9A",
        "description": "Ascolta \"Easy On Me\" qui: http://Adele.lnk.to/EOMPre-order Adele's new album \"30\" before its release on November 19: https://www.adele.comShop the \"Adele..."
      },
      {
        "order": 2,
        "title": "Hello Definition & Meaning - Merriam-Webster",
        "url": "https://www.merriam-webster.com/dictionary/hello",
        "description": "Scopri l'origine, l'uso e i sinonimi della parola hello, un'espressione o gesto di saluto. Guarda gli esempi di hello in frase e parole collegate dal dizionario."
      }
    ]
  }
}
```

## Licenza

[MIT](./../../LICENSE)
