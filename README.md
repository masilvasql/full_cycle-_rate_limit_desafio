### Como utilizar o projeto.

## 1- Definir as variáveis de ambiente no arquivo .env

#### Arquivo .env 

* IS_LIMITED_BY_IP: true --> Define que o middleware aceitara limitação por IP. Caso o valor for false, o middleware aceitara limitação apenas por token.
* IS_LIMITED_BY_TOKEN: true --> Define que o middleware aceitara limitação por token. Caso o valor for false, o middleware aceitara limitação apenas por IP.
    * Caso os dois valores sejam false, o middleware não aceitara requisições e retornará o status 415 (Não autorizado).
* SERVER_PORT: 8080 --> Define a porta que o servidor web irá responder.
* DRIVER: redis --> Define o driver que será utilizado para armazenar as informações de limitação. outras implementações podem ser feitas. 

* REDIS_HOST: localhost --> Define o host do Redis.
* REDIS_PORT: 6379 --> Define a porta do Redis.
* REDIS_PASSWORD: "" --> Define a senha do Redis.
* REDIS_DB: 0 --> Define o banco de dados do Redis.

## 2- Cadastrando regras de limitação por IP

***Para deixar mais dinâmico o projeto, foi utilizado o Redis para armazenar as informações de limitação. 
Existem rotas administrativas para configurar o rate limiter.*** 

## 2.1 Configurar limitação por IP

Deverá ser utilizada a rota */admin/ip/ip-rule* para configurar a limitação por IP.

<b>Poderá ser utilizado o arquivo .http que se encontra em:</b> *http/admin/ip.http* 

Exemplo de requisição:

### CREATE IP RULE

POST http://localhost:8080/admin/ip/ip-rule

Content-Type: application/json

```json
{
"ip": "192.168.0.213", // IP que será limitado
"max_request": 3, // Quantidade máxima de requisições por segundo
"expires_in": "10s" // Tempo de bloqueio do IP
}
```

#### Outras Rotas Administrativas de IP
____
GET http://localhost:8080/admin/ip/ip-rule/all

Content-Type: application/json

Accept: application/json

____
### GET BY IP
GET http://localhost:8080/admin/ip/ip-rule/123.456.789.013

Content-Type: application/json

Accept: application/json

____
### UPDATE BY ID
PUT http://localhost:8080/admin/ip/ip-rule/98effa60-25f1-440d-a32a-52c26565a131

Content-Type: application/json

Accept: application/json

```json
{
"ip": "192.168.0.213", // IP que será limitado
"max_request": 3,// Quantidade máxima de requisições por segundo
"expires_in": "5s"// Tempo de bloqueio do IP
}
```

____
### DELETE BY ID
DELETE http://localhost:8080/admin/ip/ip-rule/9d1f4814-7de3-4114-95be-289e4b48d6b2

Content-Type: application/json

Accept: application/json

____________________________________________________
## 3- Configuração por Token
Um token independente pode ser gerado através da rota /admin/token/token-rule.

<b>Poderá ser utilizado o arquivo .http que se encontra em:</b> *http/admin/token.http*

Exemplo de requisição:

### CREATE token RULE
POST http://localhost:8080/admin/token/token-rule

Content-Type: application/json

***INPUT***
```json
{
"max_request": 20, // Quantidade máxima de requisições por segundo
"expires_in": "5s" // Tempo de bloqueio do token
}
```

***OUTPUT***

```json
{
  "ID": "1393eb6b-d8f1-44ac-b374-8b2239fb1a87",
  "Token": "6bf64ec5-5eee-4d14-994f-365c929b37fb",
  "MaxRequest": 20,
  "ExpiresIn": "5s",
  "CreatedAt": "2025-03-17T08:03:07.8967545-03:00"
}
```

#### Outras Rotas Administrativas de Token

### GET ALL

GET http://localhost:8080/admin/ip/ip-rule/all

Content-Type: application/json

Accept: application/json

### GET BY IP
GET http://localhost:8080/admin/ip/ip-rule/123.456.789.013

Content-Type: application/json

Accept: application/json

### UPDATE BY ID
PUT http://localhost:8080/admin/ip/ip-rule/98effa60-25f1-440d-a32a-52c26565a131

Content-Type: application/json

Accept: application/json

```json
{
"ip": "192.168.0.213",
"max_request": 3,
"expires_in": "5s"
}
```

### DELETE BY ID
DELETE http://localhost:8080/admin/ip/ip-rule/9d1f4814-7de3-4114-95be-289e4b48d6b2
Content-Type: application/json
Accept: application/json


### 3- Testando o Rate Limiter

Após as configurações, existem rotas para teste que passam pelo middleware de limitação.
são elas:
/app/hello 
/app/bye

Pode ser utilizado o arquivo .http que se encontra em: *http/app/test.http*

### Pode ser utilizado o apache benchmark para testar a limitação.

AB para teste com IP
```shell
ab -n 3 -c 3 -v 2 http://192.111.0.111:8080/app/hello | grep "HTTP/"
```

AB para teste com Token
```shell
ab -n 3 -c 3 -H "API_KEY: 5df77e0d-64fd-4297-87f6-fa4f88c73718" -v 2 http://192.111.0.111:8080/app/hello | grep "HTTP/"
```