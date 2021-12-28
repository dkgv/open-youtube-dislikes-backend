if [ ! -d "scripts" ]; then
  cd ..
fi

rm -rf ./generated/restapi
mkdir -p ./generated/restapi

alias swagger='docker run --rm -it  --user $(id -u):$(id -g) -e GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger'
swagger generate server -t generated/restapi -f ./api/swagger.yaml --exclude-main