# fullcycle pacotes privados
- criando repositorio privado

```shell
go mod init github.com/bianavic/GOPRIVATE 
go: creating new go.mod: module github.com/bianavic/GOPRIVATE
go: to add module requirements and sums:
        go mod tidy
```
- configurando GOPRIVATE

```shell
go mod tidy
go: finding module for package github.com/bianavic/utils_secret/pkg/events
go: found github.com/bianavic/utils_secret/pkg/events in github.com/bianavic/utils_secret v0.0.0-20250202212629-1ff6398089fa
```

- configurando vendor

```shell
go mod vendor
```