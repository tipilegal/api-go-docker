# API de Usuarios (Go) - Docker + CI/CD

Peguei a API pronta do repositorio do curso (CRUD de usuarios em Go, guardando tudo em memoria) e coloquei ela pra rodar em Docker, Docker Compose com Postgres e PGAdmin, e fiz uma pipeline de CI/CD basica.

Repo original: https://github.com/profdiegocbcastro/gerencia-de-config

## O que precisei mudar na API

A api original guardava os usuarios so em memoria (perdia tudo quando reiniciava), sem nenhuma ligacao com banco. A atividade pede pra api acessar banco de dados, entao troquei o repositorio em memoria por um repositorio Postgres (arquivo user/postgres_repository.go), usando o driver lib/pq. Criei tambem um db.go que abre a conexao lendo as variaveis de ambiente e tenta de novo por alguns segundos caso o banco ainda nao esteja pronto (isso acontece bastante quando sobe tudo junto no compose).

Tambem corrigi uma inconsistencia que ja vinha no codigo original: o log dizia que a api rodava na porta 3000 mas o servidor de fato subia na 3001. Ajustei pra rodar na 3000 mesmo, que e o que a atividade pede.

Fora isso nao mexi em nada da regra de negocio (os endpoints de usuario continuam os mesmos, so troquei de onde os dados vem).

## Rodando sem docker

go mod tidy
export DB_HOST=localhost DB_PORT=5432 DB_USER=apiuser DB_PASSWORD=apipassword DB_NAME=apidb
go run .

## Atividade 1 - Docker

docker build -t api-go .
docker run -p 3000:3000 api-go


O Dockerfile usa multi-stage: um estagio com a imagem do Go pra compilar o binario, e outro com alpine, so pra rodar o binario final. Fica bem mais leve que carregar toda a toolchain do Go na imagem final.

## Atividade 2 - Compose

docker compose up --build

Sobe 3 containers: a api na porta 3000, o postgres na 5432 e o pgadmin na 8080.

Pra entrar no pgadmin: localhost:8080, login admin@admin.com senha admin123. Servidor: host "db", user apiuser, senha apipassword, banco apidb.

Rotas disponiveis: GET /users, POST /users, GET /users/{id}, PUT /users/{id}, DELETE /users/{id}.

## Desafio - CI/CD

O workflow em .github/workflows/ci-cd.yml builda a aplicacao Go, builda a imagem docker e sobe o compose pra ver se nao quebra.

Nao coloquei SonarQube, Semgrep, Trivy, deploy no Render nem OWASP ZAP porque ia precisar de conta em varios servicos externos e configurar secret no github, preferi entregar o essencial funcionando.
