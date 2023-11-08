Regra de negócio:
1. Ordem de Serviço: recebe o preço e a taxa e calcula o preço final. 


- Temos um webserver.
- Temos um graphQL.
- Temos um gRPC.
- Ao criar a ordem de serviço, gera um evento que envia uma msg ao RabbitMQ.



STEPS:
1. Pasta INTERNAL - Pasta interna do projeto. Coração. 
   1. Pasta entity - Coração do app. Possui a regra de negócio. Cria uma ordem e calcula o preço final. 
   2. Pasta usecase - Cria a intenção do usuário. Recebe os dados, cria a ordem e dispara o evento. 
      1. create-order.go - Componentes:
         1. Temos o inputDTO e o outputDTO para entrada e saida de dados. 
         2. CreateOrderUseCase - Tem um repo para acessar o BD. Tem um evento para a ordem criada. E o EventDispacher que dispacha um evento. São todos interfaces. Pois se quisermos mudar qq um deles, a interface é a mesma. 
         3. Execute - O coração do sistema. Pega os dados do input, cria a ordem, calcula o valor final, gera o evento, dispara o evento e retorna os dados de outputDTO. 
   3. Pasta infra - Comunicação com o mundo externo. (gateway, bd, graph, grpc, web, ...)
      1. database - como na interface do entity temos um método save, aqui no database também temos que ter a definição dele.
      2. web - pasta do webserver
         1. order_handler.go - handler que roda quando for chamada a rota de order. Ele trabalha entre a requisição externa, cria o inputDTO, tem o OrderRepository, EventDispacher, EventOrder e cria o UseCase. Pega a resposta e retorna ao mundo externo. 
      3. grpc - pasta do grpc
         1. pb - gerado pelo protoc -> raiz> protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
         2. protofiles - aqui temos o createOrder, que recebe um CreateOrderRequest e retorna um CreateOrderResponse
         3. service - service do gRPC - usa os arquivos gerados pelo protoc, da pasta pb. Recebe um UseCase. 
            1. Pelo UseCase, pega o inputDTO, executa o service do gRPC e retorna o response do service.
      4. graph - Pasta do GraphQL - geramos os arquivos, criamos o resolver CreateOrder e acrecentamos no resolver.go.
         1. schema.graphqls - arquivo criado para gerar os arquivos do graphql.  
            1. Model/models.gen_go - arquivo gerado pelo gen. service com order input output. 
         2. generated.go, resolver.go, schema.resolvers.go - arquivos todos gerados. 
         3. resolver.go - arquivo editado para linkar com o usecase. 
         4. schema.resolvers.go - arquivo que vamos usar. Gera o inputDTO, executa a order e retorna. Após gerar pelo sqlgen editamos ele para completar as funções.
   4. Pasta event - Eventos do sistema, essa pasta poderia estar dentro de infra. 
   5. Pasta cmd - pasta com arquivos .go de execução. 
      1. .env - arquivo de configuração
      2. main.go - arquivo principal de execução .go. 
      3. wire.go - arquivo para o dependency injection do wire
      4. wire_gen.go - arquivo gerado pelo wire a partir do wire.go. 
   6. Pasta config
      1. config.go - arquivo que carrega as configurações. Starter. 
