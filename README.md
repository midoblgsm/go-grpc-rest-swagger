#### Golang grpc starter kit
TL;DR Bolerplate to get started with protobuf and grpc, bonus: REST gateway, swagger api and docker.
Clone the repo, build docker image and run:
```
git clone "github.com/midoblgsm/gogrpcrestswagger-boilerplate"

cd gogrpcrestswagger-boilerplate

docker build -t myniceboilerplate .
docker run -p 7788:7788 -p 7789:7789 myniceboilerplate
```

You can see the swagger api at http://localhost:7789/swaggerui

If you modify the `api/api.proto` you need to run `make` to generate the server, (client as a bonus) and swagger api.