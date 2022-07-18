# Prueba técnica Mercadolibre

Identificar si un ADN (array de caracteres) pertenece o no a un mutante. 
Se sabe que es mutante si se encuentra más de una secuencia de cuatro caracteres A, T, C, G.

## Solución

### Validaciones iniciales:
* Se valida que el body que llega de la peticion, tenga informacion
* Se valida que los caracteres correspondan a las letras A, T, C, G para la version 1
* Se valida que cada cadena tenga la misma cantidad de caracteres
* Se valida que sea una matriz NXN y sea una matriz mayor o igual de 4 caracteres
     NOTA: Si alguna validación falla, finaliza y retorna estado de error.

Se inicia validando la información del body de la petición, agregando un contador cada vez que se consuma el servicio.
Posteriormente se valida y se arma una matriz con el array que contiene el adn de entrada, y se comienza a iterar las columnas de una fila para encontrar si existe un mutante con una posición horizontal, posterior se itera las filas para una columna determinada y encontrar en posición vertical, y por última iteramos con diagonal derecha y diagonal izquierda con una posición oblicua.
Cada método retorna como tipo un boolean, porque el objetivo de cada función es preguntar si existe o no un mutante.
Cada vez que la función encuentra un mutante, lo agregamos en un documento, para guardar un registro de los mutantes.

### Ejemplo:
//imagen

## Tecnologías utilizadas
* go1.18.3
* dynamodb
* aws

## Servicios Rest

*	/mutant<br>
     Servicio POST que recibe un Json con el siguiente formato

{
"adn": ["TGCTA", "TCGAT", "TTAGA", "TAAAA"]
}

*	/stats<br>
     Servicio GET que retorna un Json con el siguiente formato
     {
     "count_mutant_dna":40,
     "count_human_dna":100,
     "ratio":0.4
     }

### Códigos de respuesta

* 200 - Ok
* 403 - Forbidden
*	400 - Bad request

## Instrucciones de ejecución


## Pruebas unitarias
//imagen

