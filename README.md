 Service-Consumer-Producer-Golang
 Desaﬁo: Simulação de emprés mo
Credor: Pessoa jurídica que concede emprés mos a Tomadores.
Tomador: Pessoa sica que solicita emprés mo ao Credor.

É uma aplicação no modelo SaaS:
o

Deve-se isolar os dados dos credores (mul -tenant).
4 Clientes (1 com alta volumetria):
o
Horário de pico de volumetria das 13h às 15h.
 Taxas personalizadas por credor via setup.
 Realizar o cálculo de juros simples. Fórmula: J = C x i x t - J é o juro, C é o capital, i é a
taxa de juro e t é o tempo.
 Contratos:
o
o
Contrato de criação (Corpo):
 ID do tomador;
 Valor do emprés mo (Moeda Real);
 Quan dade de parcelas (Mensal).
JWT payload:
 ID do credor;
 Escopo;
 Expiração.
 Simulação pode ser aceita ou negada pelo tomador.
 No ﬁcações que devem ser enviadas ao credor:o Simulações criadas;
o Simulações aceitas;
o Simulações negadas.
 Necessário manter as simulações por três meses, por questões de governança, após
isso as simulações devem ser excluídas.
 Deve-se validar questões de an fraude nas simulações.


o A empresa já possui uma aplicação de an fraude que opera via API REST.
o Contrato:
 Id tomador;
 Valor.
É necessário validar se o credor possui o serviço de simulação contratado.
o Através de escopo no JWT;
o Simulação deve validar o escopo recebido em um campo no payload do JWT.
A cobrança do SaaS para o credor será realizada por quan dade de requisições:
o Tempo de resposta menor do que um segundo: 100% do valor;
o Tempo de resposta de até dois segundos: 70% do valor;
o Tempo de resposta maior do que dois segundos: 50% do valor.