## Anotações

### O que é?

 - É um framework desenvolvido pela Google, com o objetivo de descongestionar a rede em momentos de muitas requisições.

### Quando utilizar?
 - Em aplicações baseadas em Microsserviços
 - Em aplicações Mobile, alguns casos são possíveis
 - Futuramente haverá a opção para usar no próprio browser

### RPC
 - Remote Procedure Call
   - O cliente executa uma função remota que está no servidor

### Protocol Buffers
 - É uma das caracteristicas do gRPC, basicamente o que difere do próprio RPC já existente.
 - É uma linguagem neutra desenvolvida pela Google
 - Facilmente extensível
 - A informação é muito menor do que se comparar com XML ou até mesmo o REST
 - Trabalha com a informação em formato binário, ao invés de plain-text
   - É performático ao serializar e deserializar as informações
   - Gasta menos recurso de rede, porque seus arquivos são ainda menores

### HTTP/2
 - Só é possível porque utiliza o protocolo HTTP/2
 - Seus dados são trafegados em binários, sendo mais rápido
 - Utiliza apenas uma conexão TCP para receber todos os dados necessários
   - Por exemplo, arquivos `.css` no HTTP/1.1, para cada arquivo é feito um download, já no HTTP/2 todos são baixados na mesma conexão.
 - Seus Headers são comprimidos
 - Gasta menos recurso de rede

<hr>

## Formatos de envio/recebimento gRPC

### Server Streaming
 - Significa que o server não precisa esperar todo o seu processamento terminar para retornar os dados em um único response. Conforme os dados vão ficando prontos, eles são enviados de volta para o cliente, com isso vários responses parciais são enviados até terminar.

 - É possível fazer inclusive, `Client Streaming`, enviando os dados aos poucos para o Server.

<hr>

## REST vs gRPC

### REST
 - Padrão de envio Text/JSON
 - Unidirecional
 - Alta latência
 - Sem contrato, podendo ocorrer erros por dados enviados em um formato inválido
 - Não suporta Streming
 - Design Pré-Definido, famoso CRUD (GET, PUT, DELETE, PATH, POST)

### gRPC
 - Utiliza Protocol Buffers
 - Bidirecional e Assíncrono
 - Baixa latência
 - Utiliza um contrato (arquivo .proto)