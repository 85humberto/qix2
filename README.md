# qix2

Para gerar o código do cliente e servidor:

Baixar bibliotecas
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

Instalar para conseguir executar o "protoc"
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

Cria os códigos na pasta pb (definida na opção go_package do arquivo .proto)
protoc --go_out=. --go-grpc_out=. proto/transacao.proto