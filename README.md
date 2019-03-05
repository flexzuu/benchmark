# benchmark

### create template for a service eg. user service
```
docker run --rm -v (PWD):/local openapitools/openapi-generator-cli generate -i /local/user.yaml -g go-server -o /local/out/go
```


# TODO
gRPC:
- [x] User
- [x] Post
- [x] Rating
- [x] Facade
- [x] Seed
- [x] Client
- [x] Client-Facade
- [x] Stats
- [x] docker-compose

REST:
- [x] User
- [x] Post
- [x] Rating
- [x] Facade
- [x] Seed
- [x] Pass URI via ENV-Vars for rest-services
- [x] Client
- [x] Client-Facade
- [ ] Stats
- [x] docker-compose

GraphQL:
- [x] User
- [x] Post
- [x] Rating
- [ ] Facade
- [ ] Seed
- [ ] Client
- [ ] Client-Facade
- [ ] Stats
- [ ] docker-compose