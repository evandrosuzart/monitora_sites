# Monitorador de Site

Este é um projeto básico em Go que permite monitorar o status de sites e exibir logs das consultas anteriores de status.

## Funcionalidades

1. **Monitoramento de Sites**:
   - O programa verifica periodicamente o status de sites fornecidos.
   - Se um site estiver online (status 200), ele é considerado "carregado com sucesso".
   - Se um site estiver offline ou retornar um status diferente de 200, ele é considerado com "problemas".

2. **Logs de Consultas Anteriores**:
   - O programa mantém um registro dos resultados das consultas anteriores.
   - Os logs incluem a data e hora da consulta, o site verificado e se estava online ou com problemas.

3. **Exemplos de Leitura e Escrita de Arquivos**:
   - O programa lê os sites a serem monitorados de um arquivo.
   - Ele também registra os resultados das consultas em um arquivo de log.

## Como Usar

1. Clone o repositório:

   ```bash
   git clone https://github.com/seu-usuario/monitorador-de-site.git
   ```

2. Execute o programa:

   ```bash
   cd monitor_sites
   go run main.go
   ```

3. Adicione os sites que deseja monitorar no arquivo `sites.txt`.

4. Verifique os resultados no arquivo `log.txt`.
