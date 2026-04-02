# Challenge Backend 7

## Conectando no MySQL do Docker

Este projeto sobe dois containers:
- app: aplicação Go com cliente MySQL instalado
- db: servidor MySQL 8

### 1. Subir os containers

```bash
docker compose up -d --build
```

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

No `docker-compose.yml`, as configurações atuais são:

- MYSQL_HOST=db
- MYSQL_PORT=3306
- MYSQL_USER=root
- MYSQL_PASSWORD=root
- MYSQL_DATABASE=app

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
