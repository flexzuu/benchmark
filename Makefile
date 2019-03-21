base = micro-service/
project = github.com/flexzuu/benchmark/
grpc = grpc/
rest = rest/
hal = hal/
graphql = graphql/
mixed = mixed/
grpcbase = .$(base)$(grpc)
gosrc = $(GOPATH)/src/
grpcbase = $(gosrc)$(project)$(base)$(grpc)
restbase = $(gosrc)$(project)$(base)$(rest)
halbase = $(gosrc)$(project)$(base)$(hal)
graphqlbase = $(gosrc)$(project)$(base)$(graphql)
mixedbase = $(gosrc)$(project)$(base)$(mixed)
user = user/user
post = post/post
rating = rating/rating

facade = facade/facade

stats = stats

docc-grpc = $(grpcbase)docker-compose.yaml
docc-rest = $(restbase)docker-compose.yaml
docc-hal = $(halbase)docker-compose.yaml
docc-graphql = $(graphqlbase)docker-compose.yaml
docc-mixed = $(mixedbase)docker-compose.yaml

generate:
	protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(user) $(grpcbase)$(user)/*.proto
	protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(post) $(grpcbase)$(post)/*.proto
	protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(rating) $(grpcbase)$(rating)/*.proto
	protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(facade) $(grpcbase)$(facade)/*.proto
	protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(stats) $(grpcbase)$(stats)/*.proto
up-grpc:
	docker-compose -f $(docc-grpc) up --build --scale client=0 --scale client-facade=0 -d
logs-grpc:
	docker-compose -f $(docc-grpc) logs -f
down-grpc: 
	docker-compose -f $(docc-grpc) down
benchmark-grpc: benchmark-client-grpc benchmark-client-facade-grpc
benchmark-client-grpc:
	docker-compose -f $(docc-grpc) up --no-deps --build client
benchmark-client-facade-grpc:
	docker-compose -f $(docc-grpc) up --no-deps --build client-facade

up-rest:
	docker-compose -f $(docc-rest) up --build --scale client=0 --scale client-facade=0 -d
logs-rest:
	docker-compose -f $(docc-rest) logs -f
down-rest: 
	docker-compose -f $(docc-rest) down
benchmark-rest: benchmark-client-rest benchmark-client-facade-rest
benchmark-client-rest:
	docker-compose -f $(docc-rest) up --no-deps --build client
benchmark-client-facade-rest:
	docker-compose -f $(docc-rest) up --no-deps --build client-facade

up-hal:
	docker-compose -f $(docc-hal) up --build --scale client=0 --scale client-facade=0 -d
logs-hal:
	docker-compose -f $(docc-hal) logs -f
down-hal: 
	docker-compose -f $(docc-hal) down
benchmark-hal: benchmark-client-hal benchmark-client-facade-hal
benchmark-client-hal:
	docker-compose -f $(docc-hal) up --no-deps --build client
# benchmark-client-facade-hal:
# 	docker-compose -f $(docc-hal) up --no-deps --build client-facade

up-graphql:
	docker-compose -f $(docc-graphql) up --build --scale client=0 --scale client-facade=0 -d
logs-graphql:
	docker-compose -f $(docc-graphql) logs -f
down-graphql: 
	docker-compose -f $(docc-graphql) down
benchmark-graphql: benchmark-client-graphql benchmark-client-facade-graphql
benchmark-client-graphql:
	docker-compose -f $(docc-graphql) up --no-deps --build client
benchmark-client-facade-graphql:
	docker-compose -f $(docc-graphql) up --no-deps --build client-facade

up-mixed:
	docker-compose -f $(docc-mixed) up --build --scale client=0 --scale client-facade=0 -d
logs-mixed:
	docker-compose -f $(docc-mixed) logs -f
down-mixed: 
	docker-compose -f $(docc-mixed) down
benchmark-mixed: benchmark-client-mixed benchmark-client-facade-mixed
benchmark-client-mixed:
	docker-compose -f $(docc-mixed) up --no-deps --build client
benchmark-client-facade-mixed:
	docker-compose -f $(docc-mixed) up --no-deps --build client-facade

gui: gui-1 gui-2 gui-3 gui-4
gui-1: 
	grpcui -port 50160 -plaintext localhost:50060
gui-2:
	grpcui -port 50151 -plaintext localhost:50051
gui-3:
	grpcui -port 50152 -plaintext localhost:50052
gui-4:
	grpcui -port 50153 -plaintext localhost:50053 
generate-openapi-user:
	docker run --rm -v $(restbase)user/openapi:/local openapitools/openapi-generator-cli generate --model-name-suffix=Model -i /local/user.yaml -g go-gin-server -o /local/out/go
generate-openapi-user-client:
	# client
	docker run --rm -v $(restbase)user/openapi:/local openapitools/openapi-generator-cli generate --model-name-suffix=Model -i /local/user.yaml -g go -DpackageName=client -o /local/client
generate-openapi-post:
	# server
	docker run --rm -v $(restbase)post/openapi:/local openapitools/openapi-generator-cli generate --model-name-suffix=Model -i /local/post.yaml -g go-gin-server -o /local/out/go
generate-openapi-post-client:
	# client
	docker run --rm -v $(restbase)post/openapi:/local openapitools/openapi-generator-cli generate --model-name-suffix=Model -i /local/post.yaml -g go -DpackageName=client -o /local/client
generate-openapi-rating:
	# server
	docker run --rm -v $(restbase)rating/openapi:/local openapitools/openapi-generator-cli generate --model-name-suffix=Model -i /local/rating.yaml -g go-gin-server -o /local/out/go
generate-openapi-rating-client:
	# client
	docker run --rm -v $(restbase)rating/openapi:/local openapitools/openapi-generator-cli generate --model-name-suffix=Model -i /local/rating.yaml -g go -DpackageName=client -o /local/client
generate-openapi-facade:
	# server
	docker run --rm -v $(restbase)facade/openapi:/local openapitools/openapi-generator-cli generate --model-name-suffix=Model -i /local/facade.yaml -g go-gin-server -o /local/out/go
generate-openapi-facade-client:
	# client
	docker run --rm -v $(restbase)facade/openapi:/local openapitools/openapi-generator-cli generate --model-name-suffix=Model -i /local/facade.yaml -g go -DpackageName=client -o /local/client