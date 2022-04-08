module ShoppingBasket

go 1.16

require (
	entgo.io/ent v0.10.1
	github.com/asim/go-micro/plugins/config/encoder/yaml/v4 v4.0.0-20220304145315-a2f6fac85247
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/google/wire v0.5.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/kevinburke/ssh_config v1.1.0 // indirect
	github.com/miekg/dns v1.1.46 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/stretchr/testify v1.7.1-0.20210427113832-6241f9ab9942
	go-micro.dev/v4 v4.6.0
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
	gotest.tools v2.2.0+incompatible
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.27.0
