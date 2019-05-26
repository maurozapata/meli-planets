# Weather predictor
Los Vulcanos requieren tener una base de datos con las condiciones meteorológicas de todos los días y brindar una API REST de consulta sobre las condiciones de un día en particular.

## Comenzando

Para correr el programa se necesita RedisDB
```
docker run --name <name> --restart always -p 6379:6379 -d redis --requirepass <password>
```

Es necesario cargar datos de configuracion en un archivo .env
```
REDIS_ADDRESS="localhost:6379"
REDIS_PASSWORD="password"
REDIS_DATABASE=0
PORT=7080
API_AUTH_TOKEN="localtoken"
```

Correr la tool que calculara el clima de los proximos 10 años
```
go run cmd/tool/main.go
```

Levantar el servidor con la API
```
go run cmd/server/main.go
```

## Consideraciones previas
* Si bien cada año tiene una duración distinta para cada planeta (ferengi 360, betasoide 120 y vulcano 72 días) y el que nos pidió el servicio es la civilización de los vulcanos, se toma como medida de año la cantidad de 365 días.

* Cada 360 días el ciclo de climas vuelve a comenzar, esto solo se debe a que la velocidad de los planetas es entera, si la velocidad no fuera un número entero el ciclo no se repetiría cada 360 días, por lo que se calcula el clima de todos los días, permitiendo así poder cambiar las velocidades si es necesario.

* La posición del sol es (x, y) = (0, 0), el sistema inicia con todos los planetas alineados en el ángulo 0 usando coordenadas polares (ángulo y distancia al sol), el movimiento de los planetas es discreto y los datos sistema comienzan a partir del día uno.

* Como se realizan cálculos matemáticos, a fin de simplificar las cuentas, la precisión en los cálculos matemáticos será de 2 decimales.

* En caso de que el sol se encuentre en el perímetro del triángulo no se lo considerara dentro de él.
