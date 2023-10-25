Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:


NOT WORKING :(
https://cdn.apicep.com/file/apicep/" + cep + ".json

https://brasilapi.com.br/api/cep/v1/{cep}

http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.


EX. JSON brasilapi
URL: https://brasilapi.com.br/api/cep/v1/{cep}
{
"cep": "89010025",
"state": "SC",
"city": "Blumenau",
"neighborhood": "Centro",
"street": "Rua Doutor Luiz de Freitas Melro",
"service": "viacep"
}

EX. JSON viacep
 http://viacep.com.br/ws/" + cep + "/json/
 {
  "cep": "01246-030",
  "logradouro": "Rua Itaquera",
  "complemento": "",
  "bairro": "Pacaembu",
  "localidade": "São Paulo",
  "uf": "SP",
  "ibge": "3550308",
  "gia": "1004",
  "ddd": "11",
  "siafi": "7107"
}

JSON DE DADOS DE SAIDA
saida
cep
estado
cidade
rua
service