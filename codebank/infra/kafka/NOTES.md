## Anotações

- Open-source
- Escalável
- Latência extremamente baixa
- Alta disponibilidade
- Armazenamento
- Se conecta com qualquer linguagem

### Dinâmica de Funcionamento

O sistema que gera as informações envia para o Kafka, que funciona em formato de cluster, contendo diversas máquinas. Cada máquina é chamada de Broker, e cada uma possui o seu próprio banco de dados para armazenar as mensagens.
Após isso, outros sistemas que tem interesse nestes dados, farão as consultas no Kafka
 - O Kafka não envia mensagens, ele recebe as informações e permite que sejam lidas.
 - A forma de envio das mensagens, e leitura das mensagens, usam um canal de comunicação que é chamado de `Tópico`.
   - Eles funcionam como "logs", armazenando os dados em formato de fila.
   - Lembrando que as mensagens são persistidas em disco.

### Anatomia de um Registro
1. Header
2. Key
3. Value (payload)
4. Timestamp

### Partições

Cada tópico pode ter uma ou mais partições para garantir a distribuição dos dados. A divisão das partições em brokers diferentes garante um certo nível de resiliência. <br><br>
Existe um fator importante que é o `Replication Factor`. Por exemplo, se o `Tópico` estiver utilizando suas partições separadas em 3 `Brokers`, e o fator de replicação for 2, significa que cada partição é duplicada. Ou seja, é feita uma cópia de uma partição em outro broker. Desta forma, se um dos brokers cair, as informações não se perdem, pois haverá uma cópia em algum outro broker.

### Consumer

É o sistema que lê as mensagens armazenadas no `Kafka`. Quando o `Producer` envia uma mensagem para o tópico, ela será armazenada em uma partição, e o `Consumer` terá que ler as partições do tópico para achar as mensages. Essa abordagem abre uma brecha para gerar um gargalo de vazão das mensagens.

#### Consumer Group

Para corrigir o problema citado, é possível criar um grupo de consumidores, que serão duas (ou mais) máquinas rodando o mesmo software de consumo dos dados. Com isso, é possível definir que cada `Consumer` lerá de uma `Partição`, deixando completamente otimizado.
- Porém, se houverem mais `Consumers` do que `Partições`, os adicionais ficarão parados. Pois existe uma limitação de 1 para 1 deste relacionamento.