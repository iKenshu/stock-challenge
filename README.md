# Stock Challenge

Este proyecto es una aplicación que recupera información de acciones desde una API y la muestra en una interfaz de usuario amigable. La aplicación está dividida en dos partes principales: un backend escrito en Go y un frontend escrito en Vue.js.

## Requisitos

- Go 1.20 o superior
- Node.js 18 o superior
- Docker
- Docker Compose

## Instalación y ejecución

1. Clonar el repositorio: `git clone https://github.com/ikenshu/stock-challenge.git`

2. Ir a la carpeta del proyecto: `cd stock-challenge`
3. Ir a la carpeta del frontend: `cd stock-ui`
4. Instalar las dependencias: `npm install`

El backend ya esta en lambdas de AWS así que puedes saltarte este paso, el front tiene configurada la URL de las lambdas asi que puedes solamente ejecutar el front.

Esto lo puedes hacer simplemente con el siguiente comando: `npm run serve`

Ahora deberías tener el front corriendo en http://localhost:8081/


## Licencia

Este proyecto se encuentra bajo la licencia MIT.
