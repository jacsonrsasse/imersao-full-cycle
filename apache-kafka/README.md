Esse docker-compose é especifico para criar as imagens do Kafka.

### Zookeeper
 - Ajuda a identificar os Brokers e fazer os balanceamentos

### Kafka
 - Imagem do Kafka mesmo

### Kafka-Topics-Generator
 - Cria os tópicos, neste caso, irá criar um chamado "payments" que é para onde serão enviadas as informações

### Control-Center
 - Painel de controle do Kafka

<hr>

#### Lembretes
 - No caso do Kafka, é melhor limpar sempre o armazenamento do container anterior, rodando o comando `docker-compose down` antes de subir os containers com o comando `docker-compose up`