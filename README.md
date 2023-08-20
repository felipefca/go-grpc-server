# go-grpc-server
gRPC com gRPC Gateway: Recebe duas moedas e retorna a cotação do dia

[![LinkedIn][linkedin-shield]][linkedin-url]

gRPC 50051

![image](https://github.com/felipefca/go-grpc-server/assets/21323326/99293654-9f2c-4ff4-b605-5f6a7fb7ae2e)


gRPC Gateway 8080

![image](https://github.com/felipefca/go-grpc-server/assets/21323326/ab0dbbf9-9076-4b3c-a265-1e0e5ba71953)


<!-- SOBRE O PROJETO -->
## Sobre o Projeto

Aplicação que expoe duas portas:
- 50051: gRPC
- 8080: gRPC Gateway

### Utilizando

* [![Go][Go-badge]][Go-url]
* [![gRPC](https://img.shields.io/badge/gRPC-%20-blue)](https://github.com/grpc/grpc)
* [![Docker][Docker-badge]][Docker-url]
* [![VS Code][VSCode-badge]][VSCode-url]


<!-- GETTING STARTED -->
## Getting Started

Instruções para execução da aplicação

### Prerequisites
  ```sh
  go mod tidy
  ```

  ```sh
 cp .env.example .env
  ```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/your_username_/Project-Name.git
   ```
2. exec
   ```sh
   go run main.go
   ```


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/felipe-fernandes-fca/
[Go-url]: https://golang.org/
[Go-badge]: https://img.shields.io/badge/go-%2300ADD8.svg?style=flat&logo=go&logoColor=white
[MongoDB-badge]: https://img.shields.io/badge/mongodb-%234ea94b.svg?style=flat&logo=mongodb&logoColor=white
[MongoDB-url]: https://www.mongodb.com/
[RabbitMQ-badge]: https://img.shields.io/badge/rabbitmq-%23ff6600.svg?style=flat&logo=rabbitmq&logoColor=white
[RabbitMQ-url]: https://www.rabbitmq.com/
[Docker-badge]: https://img.shields.io/badge/docker-%230db7ed.svg?style=flat&logo=docker&logoColor=white
[Docker-url]: https://www.docker.com/
[VSCode-badge]: https://img.shields.io/badge/VS_Code-007ACC?style=flat&logo=visual-studio-code&logoColor=white
[VSCode-url]: https://code.visualstudio.com/
