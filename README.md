## Get Temperature By CEP LAB Challenge

Este projeto Microservices construído em Go fornece informações meteorológicas com base em um CEP fornecido. Ele utiliza serviços de APIs externas para obter dados de localização e clima e implementa um sofisticado sistema de rastreamento distribuído e monitoramento de métricas.

### Principais Tecnologias Utilizadas

- **Golang**: Linguagem de programação usada para construir os microservices.
- **Docker/Docker Compose**: Usado para criar imagens dos serviços e gerenciar containers.
- **OpenTelemetry**: Utilizado para capturar e exportar traces, métricas e logs.
- **Prometheus e Grafana**: Usados para monitorar métricas e visualizar dados através de dashboards.
- **Zipkin e Jaeger**: Ferramentas de visualização para traces distribuídos.
- **Redis**: Banco de dados em memória utilizado pelo microservice de Rate Limiter.

### Estrutura do Projeto

- **Microservice-input**: Ponto inicial do fluxo, responsável por receber e validar o CEP.
- **Microservice-orchestration**: Com o CEP validado, busca a localidade na API externa ViaCEP e temperatura na Weather API.
- **Microservice-ratelimiter**: Limita a quantidade de requisições por segundo nos serviços input e orchestration, baseado em IP e Token personalizados.

### Executando Localmente com Docker Compose

#### Pré-requisitos

- Docker: [Instruções de instalação](https://docs.docker.com/get-docker/)
- Docker Compose: [Instruções de instalação](https://docs.docker.com/compose/install/)

#### Instruções de Execução

1. **Clone o repositório**:
   ```bash
   git clone https://github.com/NayronFerreira/otel_temperature_challenge_lab.git


2. **Navegue até a pasta raiz do projeto**:
   ```bash
   cd otel_temperature_challenge_lab
   ````

3. **Construa e inicie os containers:**:
   ```bash
   docker-compose up --build
   ````

4. **Acesse a aplicação**:
   A aplicação está configurada para receber requisições POST no seguinte endpoint:
   ```bash
   curl --location --request POST 'http://localhost:8181/' \
   --header 'Content-Type: application/json' \
   --data-raw '{
       "cep": "06448190"
   }'
   ```

Substitua "06448190" pelo CEP desejado. Exemplos de CEPs válidos:

São Paulo: 01001-000
Rio de Janeiro: 20000-000
Brasília: 70000-000

#### Visualizando Traces e Métricas

- **Zipkin**: Acessível em http://localhost:9411/, visualiza traces distribuídos coletados pelo OpenTelemetry.
- **Jaeger**: Disponível em http://localhost:16686/, oferece uma interface rica para visualização e análise de traces.
- **Grafana**: Configure a conexão com Prometheus e visualize dashboards em http://localhost:3000/. Dashboards podem ser configurados para mostrar métricas detalhadas dos serviços.
- **Prometheus**: Interface de métricas brutas disponível em http://localhost:9090/, onde consultas podem ser feitas para visualizar métricas coletadas de microservices.


####  Contribuindo

Se você deseja contribuir para o projeto, por favor faça um fork do repositório e submeta um pull request. Sua ajuda é bem-vinda!

