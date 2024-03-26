## Get Temperature By CEP LAB Challenge

Esta é uma aplicação construída em Go que fornece informações meteorológicas com base em um CEP fornecido. Ela utiliza serviços de API externos para obter dados de localização e clima.

## Estrutura do Projeto

A aplicação é dividida em alguns pacotes, organizados da seguinte forma:

- **main.go**: Ponto de entrada da aplicação. Configura e inicia o servidor web.

- **config**: Carrega as configurações da aplicação a partir de variáveis de ambiente.

- **infra/**: Contém infraestruturas auxiliares, como o servidor web (web). Caso houvesse conexao com banco de dados, você poderia adicionar o pacote database.

- **web**: Inclui o servidor HTTP, middlewares e rotas.

- **web/api**: Inclui as chamadas para APIs externas (ViaCEP e OpenWeatherMapMapper).

- **web/model**: Inclui modelos de response das API externas + entity retornaod pelo servidor web da aplicação .

- **service**: Contém a logica de conversao das temepraturas.

## Executando a Aplicação na Nuvem

A aplicacao foi deployada no Cloud Run e voce pode acessá-la em:

[https://temperature-challenge-lab-3w3tcjirhq-uc.a.run.app/{CEP}]

Substitua {CEP} pelo CEP desejado.

Exemplos de CEP:

- São Paulo: 01001-000
- Rio de Janeiro: 20000-000
- Brasília: 70000-000

## Executando Localmente com Docker Compose

Para rodar a aplicação localmente utilizando o Docker Compose, siga os passos abaixo:

Certifique-se de ter o Docker e o Docker Compose instalados em sua máquina. Se você ainda não os tem instalados, você pode baixá-los a partir dos seguintes links:

- Docker: [https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/)
- Docker Compose: [https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/)

Depois de instalar o Docker e o Docker Compose, siga os passos abaixo para executar a aplicação:

1. Clone o repositório da aplicação para o seu ambiente local.

2. Navegue até a pasta raiz do projeto.

3. Execute o comando abaixo para construir e iniciar os contêineres da aplicação.

```bash
docker-compose up --build
```
4. A aplicação estará acessível localemnte em http://localhost:8181{CEP}.

Basta substituir {CEP} pelo CEP desejado.
