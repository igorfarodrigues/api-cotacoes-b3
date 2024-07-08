# Projeto de API de Cotação (api-cotacoes-b3)

Este projeto configura um ambiente Docker para uma API de cotação, incluindo um banco de dados PostgreSQL e arquivos de configuração necessários. Ele é desenvolvido como parte de um desafio para processar e expor dados de negociações da B3.

## Estrutura do Projeto

- **go.mod**: Arquivo de definição de dependências para um projeto Go.
- **create_db.sql**: Script SQL para a criação do banco de dados.
- **config.toml**: Arquivo de configuração da aplicação.
- **docker-compose.yml**: Arquivo de configuração do Docker Compose.
- **Dockerfile**: Arquivo de configuração para construção da imagem Docker.

## Desafio

### Objetivo

Criar uma solução para processar uma grande quantidade de dados de negociações da B3 e expor uma interface para consulta desses dados.

### Requisitos

1. **Baixar dados dos últimos 7 dias úteis do histórico de negociações da B3**:
    - [B3 Histórico de Cotações](https://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/cotacoes/cotacoes/)
2. **Persistir os seguintes campos no banco de dados**:
    - "HoraFechamento"
    - "DataNegocio"
    - "CodigoInstrumento" (ticker)
    - "PrecoNegocio"
    - "QuantidadeNegociada"
3. **Expor uma interface para retornar a seguinte visualização de dados**:
    - Filtro por "ticker" e "DataNegocio"
    - Campos de agregação:
        - `max_range_value`: Maior valor negociado para um dado ticker entre todos os dias.
        - `max_daily_volume`: Quantidade negociada máxima em um mesmo dia para um dado ticker.

### Avaliação

- Uso eficiente dos recursos da linguagem para carregar os arquivos.
- Modelagem de banco de dados para escrita e leitura performáticas.
- Performance na leitura de dados.

### Pontos Importantes

- Capacidade de desenvolver a solução, não apenas construir boilerplates.
- Aplicação com testes concisos cobrindo caminhos alternativos.
- Claro uso da aplicação via CLI se não houver implementação de uma API.
- Possibilidade de carregar arquivos manualmente no banco de dados.

## Pré-requisitos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Como Usar

1. Clone o repositório:
    ```sh
    git clone https://github.com/seu-usuario/seu-repositorio.git
    cd seu-repositorio
    ```

2. Construa e inicie os containers:
    ```bash
    docker compose build && docker compose up -d
    ```

3. Execute ´main.go´:
    ```bash
    go run ./loadFiles/main.go
    
    go run ./api/main.go
    ```

## Arquivos

### go.mod

Contém as dependências e informações do módulo Go utilizado no projeto.

### create_db.sql

Script SQL utilizado para criar o banco de dados e tabelas necessárias para a aplicação.

### config.toml

Arquivo de configuração da aplicação. Inclui parâmetros e ajustes necessários para o funcionamento da API.

### docker-compose.yml

Define os serviços necessários para a aplicação, incluindo:

- **db**: Serviço do banco de dados PostgreSQL
    - Porta: 5432
    - Volume: `create_db.sql` para inicializar o banco

### Dockerfile

Instruções para a construção da imagem Docker da aplicação.

## Comandos Úteis

- Para parar os containers:
    ```sh
    docker compose down
    ```

- Para reconstruir os containers:
    ```sh
    docker compose up --build
    ```

## Contribuições

Sinta-se à vontade para contribuir com o projeto enviando pull requests ou relatando issues.

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.
