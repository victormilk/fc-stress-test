# Load Testing CLI in Go

Este projeto é uma ferramenta de linha de comando (CLI) para realizar testes de carga em um serviço web. O usuário pode fornecer a URL do serviço, o número total de requisições e a quantidade de chamadas simultâneas.

## Funcionalidades

- Realiza requisições HTTP para a URL especificada.
- Distribui as requisições de acordo com o nível de concorrência definido.
- Gera um relatório com informações específicas após a execução dos testes.

## Parâmetros de Entrada via CLI

- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requisições.
- `--concurrency`: Número de chamadas simultâneas.

## Execução do Teste

Para executar o teste, use o comando:

```sh
go run main.go --url=http://example.com --requests=100 --concurrency=10
```

### Executando via Docker

Você também pode executar o teste usando Docker. Primeiro, construa a imagem Docker:

```sh
docker build -t fc-stress-test .
```

Em seguida, execute o contêiner com os parâmetros desejados:

```sh
docker run --rm fc-stress-test --url=http://example.com --requests=100 --concurrency=10
```

### Executando via Imagem (DockerHub)

```sh
docker run victormilk/fc-stress-test --url=http://example.com --requests=100 --concurrency=10