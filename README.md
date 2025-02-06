# Positivity Pulse

## Descrição

O **Positivity Pulse** é um projeto desenvolvido com fins educacionais, visando aprender e aplicar diversas tecnologias modernas no desenvolvimento de aplicações backend e frontend. O objetivo principal do projeto é simular um serviço que consome e publica mensagens, processando-as e enviando notificações em tempo real aos usuários.

### Tecnologias utilizadas

- **Go (Gin, GORM)**: Linguagem de programação e framework para construir a API.
- **RabbitMQ**: Sistema de filas para gerenciamento de mensagens entre os componentes do sistema.
- **PostgreSQL**: Banco de dados relacional para armazenar as mensagens processadas.
- **Vue.js 3**: Framework JavaScript para construção da interface do usuário (frontend).
- **WebSockets**: Comunicação bidirecional em tempo real entre o backend e o frontend.
- **Goroutines**: Para paralelização e processamento assíncrono no backend.

### Arquitetura

Este projeto segue a **arquitetura Ports and Adapters**, também conhecida como **Arquitetura Hexagonal**, para separar claramente as responsabilidades de cada parte do sistema.

- **Adaptadores**: A camada de adaptadores (como RabbitMQ, PostgreSQL e WebSocket) comunica as diferentes partes do sistema com a lógica de negócios, mantendo o core do sistema isolado e testável.
- **Ports**: São as interfaces que definem como os adaptadores interagem com o sistema.
- **Core (Domínio)**: Lógica de negócios que define o que a aplicação faz, como a criação e publicação de mensagens de conselhos.

### Funcionalidades

- **Consumo de mensagens**: O sistema consome mensagens de uma fila do RabbitMQ e as processa, salvando-as no banco de dados PostgreSQL.
- **Notificação em tempo real**: Ao processar uma nova mensagem, o sistema envia uma notificação em tempo real para todos os clientes conectados via WebSocket.
- **API REST**: Oferece rotas para verificar o status da aplicação e disparar a publicação de novas mensagens.

### Rotas

- **`GET /alive`**: Retorna o status da aplicação.
- **`GET /fetch-advice`**: Dispara a publicação de uma nova mensagem de "conselho" na fila RabbitMQ. Resposta: `"Message sent to the queue"`.
- **`GET /ws`**: Estabelece uma conexão WebSocket para receber notificações em tempo real.

### Como rodar

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/seu-usuario/positivity-pulse.git
   cd positivity-pulse
2. **Tendo o docker instalado em sua máquina apenas execute:**

   `docker-compose up -d`
   - Certifique-se de que o `RabbitMQ` e o `PostgreSQL` estão em execução na sua máquina local ou em um ambiente configurado.
   - Configure as credenciais do RabbitMQ e PostgreSQL no código conforme necessário.
3. **Rodando o projeto:**
   Para rodar o backend (Go API):
   
   ```bash
   go run cmd/main.go
