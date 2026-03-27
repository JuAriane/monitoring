# monitoring

Aplicação simples de monitoramento de sites em Go.

## Visão geral

Este projeto é um exemplo didático de um monitor de sites escrito em Go. O programa apresenta uma pequena interface de console que exibe uma introdução, um menu com comandos e executa ações básicas conforme a opção selecionada.

## Funcionalidades

- Exibe introdução com nome do autor e versão.
- Mostra menu interativo no terminal.
- Recebe comandos do usuário e executa ações placeholder (monitoramento, exibição de logs, saída).

## Estrutura do projeto

- `main.go` — ponto de entrada da aplicação; inicializa a interface e encaminha comandos.
- `internal/app/` — lógica de leitura e execução de comandos (`command.go`, `monitor.go`).
- `internal/ui/` — funções de interface de console (`introduction.go`, `menu.go`).
- `internal/domain/` — (placeholder) modelos de domínio (`site.go`).
- `internal/infra/` — (placeholder) infraestruturas como cliente HTTP e logging (`http_client.go`, `log.go`).

## Requisitos

- Go 1.20+ (qualquer versão moderna do Go deve funcionar)

## Instalação e execução

No diretório raiz do projeto, execute:

```bash
go build -o monitoring
./monitoring
```

Ou rode diretamente com `go run`:

```bash
go run main.go
```

## Uso

Ao executar, o programa exibirá uma introdução e o menu:

- `1` — Iniciar Monitoramento (ainda implementado como placeholder)
- `2` — Exibir logs (placeholder)
- `0` — Sair do Programa

Exemplo de sessão:

```
Olá, Jurina
Este programa está na versão 1
1- Iniciar Monitoramento
2- Exibir logs
0- Sair do Programa
> 1
Monitorando...
```

## Como estender

- Implemente a lógica real de monitoramento em `internal/app/monitor.go` ou mova a lógica de rede para `internal/infra/http_client.go`.
- Use `internal/domain` para definir estruturas e regras de domínio (por exemplo, modelo `Site`).

## Contribuição

Pull requests são bem-vindos. Para mudanças maiores, abra uma issue descrevendo a proposta antes de implementá-la.

## Licença

Este repositório não possui uma licença explícita definida. Adicione um arquivo `LICENSE` se quiser liberar o código sob uma licença específica.

---
