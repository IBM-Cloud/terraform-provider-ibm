rm -rf ../riaas/*
function install_swagger {
    # As the go get install swagger command will result in different swagger behavior due to the build dependencies
    # let's doownload the swagger binary directly
    # version: v0.19.0
    # commit: 312366608bbf17dd219190b66ab63bdc8b4d0db
    distribution=$(uname | tr '[:upper:]' '[:lower:]')
    download_url="https://github.com/go-swagger/go-swagger/releases/download/v0.19.0/swagger_${distribution}_amd64"
    curl -o /usr/local/bin/swagger -L'#' "$download_url"
    chmod +x /usr/local/bin/swagger
}
swagger version 2>&1 > /dev/null || install_swagger
swagger generate client -t ../riaas -f swagger.yaml  -A ../riaas

