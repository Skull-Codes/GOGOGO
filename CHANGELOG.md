## 1.2.5 (17-01-2025)

* Se ha cambiado el nombre de los contenedores asignando el prefijo gogogo-


## 1.2.4

* Se ha añadido una variable de entorno para definir el tiempo que tarda en expirar el token.


## 1.2.3

* Se ha añadido un pipeline para Github Actions


## 1.2.2

* Eliminada la ruta estática a los archivos subidos y ha sido reemplazada por la ruta media/ con un controlador que sirve los archivos y la opción de usar token.
* Actualizada la documentación de swagger con la ruta media/ y los tests para esta función.


## 1.2.1

* Se ha añadido la ruta estática hacia los archivos en uploads.


## 1.2.0

* Se han eliminado prints innecesarios en user-controller.go
* Se ha añadido el método GetUserFromToken para extraer el nombre de usuario a partir del token.
* Se ha creado un directorio uploads para la subida de archivos y su volumen lógico en docker para que no se eliminen los archivos al reconstruir los contenedores.


## 1.1.0

* Se ha cambiado el tipo de petición en las siguientes rutas:
    - change_password/ pasa de POST a PATCH.
    - delete_user/ pasa de POST a DELETE.


## 1.0.1

* Se corrigió un error ortográfico en el README en la palabra esqueleto, puse exqueleto porque se me vino a la cabeza exoesqueleto y ya comencé a escribirlo así en todos lados.
