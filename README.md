# Megasena

Busca o resultado da MegaSena e envia mensagem no Telegram.

A busca é feita às terças, quintas e sábados no horário definido em `horarios` no `megasena.json` (padrão 21:30).

É feita a verificação dos números sorteados. O resultado é informado na mensagem enviada.

## Apostas

No `megasena.json`, o campo `aposta` aceita um ou mais cartões (dezenas separadas por espaço):

```json
"aposta": [
  "05 12 29 32 38 42",
  "01 07 15 23 44 60"
]
```

Uma string única (`"aposta": "05 12 29 32 38 42"`) também é aceita.


Antiga URL da Caixa Econômica Federal:

```
https://servicebus2.caixa.gov.br/portaldeloterias/api/megasena
```