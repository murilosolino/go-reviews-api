# Challenge Backend 7

API REST em Go desenvolvida para cumprir um desafio de backend da plataforma Alura e colocar em prática o aprendizado na linguagem Go.

O projeto expõe endpoints para gerenciamento de **destinos** e **avaliações (reviews)**, persistindo em **MySQL 8**. As migrations rodam automaticamente ao iniciar a aplicação.

## Funcionalidades

- Health check (`GET /`) para verificação rápida do servidor.
- Reviews
	- Listar todas (`GET /reviews`)
	- Buscar por id (`GET /reviews/{id}`)
	- Buscar 3 aleatórias para “home” (`GET /reviews-home`)
	- Criar (`POST /reviews`)
	- Atualizar parcial/total (`PATCH`/`PUT /reviews/{id}`)
	- Excluir (`DELETE /reviews/{id}`)
- Destinations
	- Criar (`POST /destinations`)
	- Listar todas (`GET /destinations`)
	- Buscar por nome via querystring (`GET /destinations?name=...`)
	- Atualizar (`PUT /destinations/{id}`)
	- Excluir (`DELETE /destinations/{id}`)
- Geração opcional de `descriptive_text` via OpenAI
	- Se o campo `descriptive_text` não for enviado no `POST /destinations`, o serviço tenta gerar um texto curto automaticamente.
	- Para isso, configure `OPENAI_API_KEY`. Sem a chave, o texto pode ficar vazio.

## Como rodar (Docker)

### 1. Subir os containers

```bash
docker compose up -d --build
```

- API: http://localhost:8080
- MySQL: localhost:3306

O serviço `app` já sobe com Air automaticamente (definido no `CMD` do Dockerfile).
Nao rode `go run` manualmente dentro do container, pois isso cria outro processo fora do controle do Air e o hot reload para de fazer sentido.

Para acompanhar o reload em tempo real:

```bash
docker compose logs -f app
```

### 2. Verificar se o MySQL está saudável

```bash
docker compose ps
```

Você deve ver o serviço `db` com status `healthy`.

## Variáveis de ambiente

Definidas no `docker-compose.yml` (valores padrão atuais):

- `MYSQL_HOST=db`
- `MYSQL_PORT=3306`
- `MYSQL_USER=root`
- `MYSQL_PASSWORD=root`
- `MYSQL_DATABASE=app`

Opcional:

- `OPENAI_API_KEY` (habilita geração automática de `descriptive_text` ao criar destinos)

## Banco de dados e migrations

- O projeto usa `golang-migrate` e executa `m.Up()` automaticamente na inicialização.
- As migrations ficam em `config/database/migrations`.

Esquema (resumo):

- `reviews`
	- `id` (auto increment)
	- `review` (texto)
	- `author_name` (varchar)
	- `url_photo` (varchar)
- `destinations`
	- `id` (auto increment)
	- `img` (varchar)
	- `name` (varchar, not null)
	- `price` (decimal)
	- `descriptive_text` (text)

## API (HTTP)

Base URL (Docker): `http://localhost:8080`

### Health check

- `GET /` → `200 OK`

### Reviews

Observação importante sobre nomes de campos:

- No banco, a coluna é `author_name`.
- Na resposta JSON, o campo sai como `author` (por causa da tag JSON do struct).
- Para evitar inconsistências, prefira enviar `author_name` nos requests de escrita.

#### Listar

```bash
curl http://localhost:8080/reviews
```

Resposta (exemplo):

```json
[
	{
		"id": 1,
		"review": "...",
		"author": "...",
		"url_photo": "..."
	}
]
```

#### Buscar por id

```bash
curl http://localhost:8080/reviews/1
```

#### Buscar 3 aleatórias (home)

```bash
curl http://localhost:8080/reviews-home
```

#### Criar

```bash
curl -X POST http://localhost:8080/reviews \
	-H "Content-Type: application/json" \
	-d '{"review":"Ótimo","author_name":"Murilo","url_photo":"http://localhost/photo.jpg"}'
```

Status: `201 Created`

#### Atualizar (PUT/PATCH)

```bash
curl -X PATCH http://localhost:8080/reviews/1 \
	-H "Content-Type: application/json" \
	-d '{"review":"Atualizado"}'
```

Status: `204 No Content`

#### Excluir

```bash
curl -X DELETE http://localhost:8080/reviews/1
```

Status: `204 No Content`

### Destinations

Observação sobre nomes de campos:

- No banco (e para requests de escrita), o campo de imagem é `img`.
- Nas respostas JSON, ele aparece como `image` (tag JSON do struct).

As rotas de destinos retornam um envelope JSON no formato:

```json
{
	"statusCode": 200,
	"message": "ok",
	"data": "..."
}
```

#### Criar

Se `descriptive_text` não for enviado, o serviço tenta gerar automaticamente.

```bash
curl -X POST http://localhost:8080/destinations \
	-H "Content-Type: application/json" \
	-d '{"name":"Paris","price":1000,"img":null}'
```

Status: `201 Created`

#### Listar

```bash
curl http://localhost:8080/destinations
```

#### Buscar por nome

```bash
curl "http://localhost:8080/destinations?name=Paris"
```

#### Atualizar

```bash
curl -X PUT http://localhost:8080/destinations/1 \
	-H "Content-Type: application/json" \
	-d '{"price":1200}'
```

Status: `204 No Content`

## CORS

O middleware de CORS está habilitado e atualmente permite origem `http://localhost:8080`.

#### Excluir

```bash
curl -X DELETE http://localhost:8080/destinations/1
```

Status: `204 No Content`

## Estrutura do projeto (alto nível)

- `main.go`: inicialização da conexão com o banco, montagem das dependências e start do servidor.
- `config/`
	- `database/`: conexão com MySQL e execução automática de migrations.
	- `router/`: declaração das rotas HTTP.
	- `dependencies.go`: wiring simples (controllers/services/models).
- `api/`
	- `controllers/`: handlers HTTP.
	- `services/`: regras de negócio (inclui integração opcional com OpenAI).
	- `model/`: acesso a dados e structs.
	- `middleware/`: CORS.
	- `helper/`: helper de resposta JSON (usado em destinations).

## Testes

Rodar na máquina host:

```bash
go test ./...
```

Rodar dentro do container:

```bash
docker compose exec app go test ./...
```

## Conectando no MySQL do Docker

Este projeto sobe dois containers:
- app: aplicação Go com cliente MySQL instalado
- db: servidor MySQL 8

## Opção A: conectar no MySQL pelo container da aplicação (app)

Essa opção usa o cliente MySQL que está instalado no container `app` e conecta no host `db` da rede interna do Docker.

```bash
docker compose exec app mysql -h db -uroot -proot --ssl=0 app
```

Se conectar com sucesso, você verá o prompt do MySQL.

### Comandos úteis no MySQL

```sql
SHOW DATABASES;
USE app;
SHOW TABLES;
SELECT NOW();
```

Para sair:

```sql
exit
```

## Opção B: entrar no container do MySQL e conectar localmente

Primeiro acesse o shell do container `db`:

```bash
docker compose exec db bash
```

Depois conecte no MySQL:

```bash
mysql -uroot -proot
```

## Opção C: conectar pela máquina host

Como a porta 3306 está publicada no compose, você pode usar qualquer cliente MySQL no seu computador:

- host: localhost
- port: 3306
- user: root
- password: root
- database: app

Exemplo com MySQL CLI local (adicione `--ssl-mode=DISABLED` ou `--ssl=false` dependendo do seu cliente se der erro de certificado):

```bash
mysql -h 127.0.0.1 -P 3306 -uroot -proot --ssl-mode=DISABLED app
```

## Credenciais e variáveis usadas no projeto

No `docker-compose.yml`, as configurações atuais são (veja também a seção **Variáveis de ambiente** acima):

- `MYSQL_HOST=db`
- `MYSQL_PORT=3306`
- `MYSQL_USER=root`
- `MYSQL_PASSWORD=root`
- `MYSQL_DATABASE=app`

Opcional:

- `OPENAI_API_KEY`

## Problemas comuns

### Access denied for user

- Confirme se está usando usuário `root` e senha `root`.
- Verifique se os containers estão rodando com `docker compose ps`.

### MySQL não fica healthy

- Veja logs com:

```bash
docker compose logs db --tail=100
```

### Porta 3306 em uso no Windows

- Pare o serviço local de MySQL/MariaDB da máquina host ou altere a porta publicada no `docker-compose.yml`.
