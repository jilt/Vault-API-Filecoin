# Mintbase API

Use swagger to specify the api.  
Generate API documentation from the `swagger.yml` file.  
Host the api documentation.

Refer to the [Reference](#references) section, it contains links to documentations of the resources that were used.

## Generate the documentation

### Install tools

Download the [swagger-codegen-cli-2.4.27.jar][9] file.  
You can also follow the installation guide from the swagger-codegen [documentation][10]

```sh
# Download current stable 2.x.x branch (Swagger and OpenAPI version 2)
wget https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/2.4.27/swagger-codegen-cli-2.4.27.jar -O swagger-codegen-cli.jar
```

The swagger-codegen binary is a java tar file. Install java.  
Use [sdkman][11] to easily install java sdk.

```sh
sdk install java
```

Now use the `swagger-codegen` tool

```sh
java -jar swagger-codegen-cli.jar --help
```

### Generate

`$HOME/programs/swagger-codegen-cli-2.4.27.jar` refers to where you have installed swagger-codegen.  
Make sure to **replace** this with the actual path to where you have saved swagger-codegen after downloading it.

Generate a swagger.json file type from the swagger.yml file.  
The swagger-codegen tool uses the swagger.json file type.

```sh
# generate a swagger.json file type from the swagger.yml file and save to the folder api-docs
java \
 -jar $HOME/programs/swagger-codegen-cli-2.4.27.jar generate \
 -i swagger.yml \
 -l swagger \
 -o api-docs
```

```sh
# generate an html documentation of your api
java \
    -jar $HOME/programs/swagger-codegen-cli-2.4.27.jar generate \
    -i api-docs/swagger.json \
    -l html2 \
    -o api-docs
```

#### Host the documentation

The folder api-docs can be deployed. as a documentation for the API.  
In this example I will host the documentation on firebase, follow the documentation at [firebase-hosting][12]

Here is the hosted documentation [https://mintbase.web.app/](https://mintbase.web.app/)

## References

- OpenAPI Version 2.0 [specification][1]
- Swagger Editor [Swagger-editor-github][2]
- Link to the Swagger editor hosted online for use [Swagger-editor][3]
- OpenAPI vscode editor [Open-API-marketplace-link][4]
- Swagger Codegen [Swagger-codegen-github][5]
- Golang Gin Web Framework [documentation][6]
- Go Text template [documentation][7]
- HTTP status codes [HTTP-Status-Codes-Website][8]
- Use SDK Manager to install the development tools, used by Swagger tools [sdkman-website][11]

[1]: https://swagger.io/specification/v2/
[2]: https://github.com/swagger-api/swagger-editor
[3]: https://editor.swagger.io
[4]: https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi
[5]: https://github.com/swagger-api/swagger-codegen
[6]: https://github.com/gin-gonic/gin
[7]: https://pkg.go.dev/text/template
[8]: https://restfulapi.net/http-status-codes
[9]: https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/2.4.27/swagger-codegen-cli-2.4.27.jar
[10]: https://github.com/swagger-api/swagger-codegen
[11]: https://sdkman.io/sdks
[12]: https://firebase.google.com/docs/hosting
