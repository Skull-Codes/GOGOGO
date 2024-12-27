# GOGOGO - Esqueleto básico en GIN-Gonic

He creado este esqueleto con el fin de simplificar un poco el tedioso proceso de empezar a crear un proyecto en blanco sin ninguna referencia base. 

Este proyecto se basa en un patrón MVC para API Rest con un sistema de autenticación JWT ya implementado, tests unitarios, linter y pipelines para CI_CD en Github y Gitlab. El proyecto está docketizado y adaptado a MongoDB

Las tecnologías que utiliza actualmente el proyecto son:

- Go
- Gin-Gonic
- MongoDB
- Docker


## Configurar el proyecto

1. Copia el **.env-example** con el nombre .env y modifica las constantes según tu necesidad. Esto se adaptará tanto al proyecto base como a la construcción de los contenedores.
2. Si cambias el nombre del directorio principal **gogogo** por uno nuevo recuerda modificar la primera línea del **go.mod** para que coincida el nombre del paquete y el del proyecto.
3. Para poner en marcha el proyecto ejecutamos ``docker compose up``
4. Si todo va bien tendremos acceso al servidor desde http://localhost:8080/docs/index.html donde encontraremos la documentación del proyecto y podremos hacer pruebas de endpoints.

## Testing

Los test se pueden ejecutar actualmente desde el proyecto con el comando:

- Local: ``go test tests/*.go`` (tendrás que configurar cosas)
- Docker (recomendado): ``docker exec -it api go test tests/*.go`` utiliza -v para ver mas detalles de los tests ejecutandose.

## Code cuality

Para validar la calidad del código tenemos que instalar el siguiente linter:

1. Descarga el lint: ``curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest``
2. Movemos en el sistema el lint: ``sudo mv ./bin/golangci-lint /usr/local/bin/``
3. Verificar instalación: ``golangci-lint --version``


Lo ejecutamos con el siguiente comando:

- `golangci-lint run`

## Documentación

1. Lo primero que haremos es lanzar el contenedor, de ese modo se instalara entre las depencencias swagger.
2. Para poder generar la documentación tenemos que añadir la siguiente variable a nuestro PATH: ``export PATH=$PATH:~/go/bin``
3. Recarga la configuración: ``source ~/.bashrc``
4. Y ahora cada vez que añadimos documentación solo ejecutamos este comando: ``swag init`` y relanzamos los contenedores para que se reflejen los cambios.

## Creando un CRUD de ejemplo (videos)

proximamente...