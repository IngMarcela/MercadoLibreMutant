# Prueba técnica Mercadolibre

Identificar si un ADN (array de caracteres) pertenece o no a un mutante. 
Se sabe que es mutante si se encuentra una secuencia de cuatro caracteres iguales de manera en horizontal, 
vertical u oblicua. Los caracteres permitidos en un ADN son: A, T, C, G.

## Solución

### flujo 
Se inicia validando la existencia de información del body de la request, Posteriormente se valida el contenido del body y se arma una matriz con el array que contiene el ADN de entrada.
Con la información validada y transformada se comienza a iterar la matriz de horizontal, posterior de manera vertical y por último de oblicua. Cada una de las anteriores iteraciones retorna si encontró o no una secuencia de cuatro caracteres iguales, dado que estas iteraciones se realizan  para tratar de encontrar si existe un mutante.
Una vez se finaliza el proceso de validación se emite una notificación a un tópico para que distintos interesados puedan conocer y procesar esto datos.

### Validaciones iniciales:
* Se valida que el body que llega de la petición, contenga información
* Se valida que los caracteres correspondan a las letras A, T, C, G 
* Se valida que el ADN sea una matriz NXN y donde N es mayor o igual a 4

  NOTA: Si alguna validación falla, finaliza y retorna estado de error con código http 400.

### Arquitectura

![arquitecture](https://user-images.githubusercontent.com/25367590/179764025-c56ec3ee-efd8-4b6d-84fb-ae13a3bf5d0f.png)

## Tecnologías utilizadas
* go1.18.3
* Lambda Functions
* SNS
* Api Gateway
* CloudWatch

## Pruebas unitarias
Cobertura 92.3%
Uso de https://github.com/vektra/mockery en generación de mocks

![Captura de Pantalla 2022-07-19 a la(s) 8 42 39 a m](https://user-images.githubusercontent.com/25367590/179765173-9b5c9251-999e-4a26-9556-9b5d960a93c2.png)
