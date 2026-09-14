# API de Usuarios (Go) - Docker + CI/CD

API do repositorio do curso, coloquei pra rodar em Docker, Compose com Postgres + PGAdmin, e fiz um CI/CD basico.

Repo original: https://github.com/profdiegocbcastro/gerencia-de-config

A api era em memoria, troquei pra usar Postgres de verdade (user/postgres_repository.go e db.go) porque a atividade pedia acesso a banco. Tambem corrigi a porta que tava errada no codigo original (3001 pra 3000).

## Rodar

docker compose up --build

Sobe api (3000), postgres (5432) e pgadmin (8080, login admin@admin.com / admin123).

Nao entrei com SonarQube, Trivy, Render e ZAP do desafio completo, foquei no essencial funcionando.
