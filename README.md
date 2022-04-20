## Decorator em Golang

Esta é uma implementação do padrão de projeto decorador na linguagem Golang (Go), tendo como base o ["kata cupcake"](https://codingdojo.org/kata/cupcake/).

Definição do padrão de projeto ["decorator"](https://refactoring.guru/pt-br/design-patterns/decorator).

#### Build Docker Image:

```sh
docker build -t decorator .
```

#### Run Docker Image:

```sh
docker run decorator
```

#### Run Project:
```sh
go run ./main.go
```

#### Output:

```shell
Preparing some cupcakes 🧁...
Cupcake with: chocolate, costing $1.10
Cupcake with: nuts, costing $1.20
Cupcake with: sugar, costing $1.00
Cupcake with: chocolate, nuts, costing $1.30       
Cupcake with: chocolate, sugar, costing $1.10      
Cupcake with: chocolate, nuts, sugar, costing $1.30
```

#### Run Tests:

```shell
go test ./tests -v
```

#### Output Tests:

```shell
=== RUN   TestChocolate_AddIngredient        
--- PASS: TestChocolate_AddIngredient (0.00s)
=== RUN   TestCupcake_AddIngredient
--- PASS: TestCupcake_AddIngredient (0.00s)  
=== RUN   TestCupcake_FullStack
--- PASS: TestCupcake_FullStack (0.00s)      
=== RUN   TestNuts_AddIngredient
--- PASS: TestNuts_AddIngredient (0.00s)     
=== RUN   TestSugar_AddIngredient
--- PASS: TestSugar_AddIngredient (0.00s)    
PASS
```