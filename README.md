[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)

<div style="text-align: center; color: dodgerblue"> <h3>LoadBalancerGo<h3> </div>

**Motivações**: Há um tempinho venho estudando Golang, e cofesso que estou gostando bastante, a fim  de melhor espandir meus conhecimentos em programação e melhor minhas skills em engenharia de software. Uma das coisas que tinha dúvida e curiosidade era de como saber como os Load Balancers "Balanceadores de Carga" funcionavam e resolvi estudar o assunto e desenvolver o meu com base no que achei na internet e tomando referência de outros projetinhos. No fim das contas uni o útil ao agradável e nasceu o projetinho em questão :)

--- 

**Como executar o projeto:**

Na raiz do projeto execute o arquivo `main.go`.

   `go run main.go`

Depois execute no terminal
  
  `curl localhost:3000`

E verá logs assim

  `Server name: SERVER (1)
`

De novo, 

`curl localhost:3000`

Veja que o servidor muda e vai até o Server (4) e volta para o primeiro.

`Server name: SERVER (2)`


*IMPORTANTE*:  Até agora eu não fiz halthy check para validar se um endpoint está ativo ou não e a única estratégia que fiz para balancear a carga é a *"Round Robin"*. Com passar dos commits, isso vai ser acrescentado com outras coisas mais.
