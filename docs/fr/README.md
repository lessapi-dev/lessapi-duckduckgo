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

(Ce document est généré par une traduction IA ; en cas de doute, veuillez vous référer aux documents en chinois
simplifié ou en anglais.)

---

## Introduction

LessAPI-DuckDuckGo est un service d'API de moteur de recherche.
Basé sur playwright et le moteur de recherche DuckDuckGo, il est encapsulé pour fournir une interface API simple.
Simple, léger, fiable, déployable via Docker, facile à maintenir.

## État

> En développement expérimental, pas recommandé pour une utilisation en production.

## Utilisation

### Documentation OpenAPI (Swagger 3.0)

[Documentation OpenAPI (Swagger 3.0)](../../resource/openapi.json)

### Recherche texte `GET /search/text`

**Paramètres de la requête:**

- keyword: Mot-clé de recherche (obligatoire)
- region: Région (facultatif) en-US, fr-FR, zh-CN, ru-RU, etc. La valeur par défaut est en-US
- maxCount: Nombre maximal de résultats retournés (facultatif) La valeur par défaut est 20
- viaProxyUrl: Adresse du proxy utilisé par le navigateur (facultatif) Exemple : http://proxy.server:3000 La valeur par
  défaut est vide

**Exemple de requête:**

```shell
curl 'http://127.0.0.1:8080/search/text?keyword=hello&maxCount=2'
```

**Exemple de réponse:**

```json
{
  "code": "success",
  "data": {
    "results": [
      {
        "order": 1,
        "title": "Adele - Hello (Vidéo musicale officielle) - YouTube",
        "url": "https://www.youtube.com/watch?v=YQHsXMglC9A",
        "description": "Écoutez \"Easy On Me\" ici: http://Adele.lnk.to/EOMPre-commandez le nouvel album d'Adele \"30\" avant sa sortie le 19 novembre: https://www.adele.comBoutique \"Adele..."
      },
      {
        "order": 2,
        "title": "Définition et signification de Hello - Merriam-Webster",
        "url": "https://www.merriam-webster.com/dictionary/hello",
        "description": "Apprenez l'origine, l'utilisation et les synonymes du mot hello, une expression ou un geste de salutation. Voir des exemples d'utilisation de hello dans des phrases et des mots associés du dictionnaire."
      }
    ]
  }
}
```

## Licence

[MIT](./../../LICENSE)
