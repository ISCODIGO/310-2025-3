# Golang

- Desarrollado por Google.
- Es un lenguaje: compilado, tipado es estático y fuerte.

## Tipos de datos
### Booleano
- bool: solo tiene `true` y `false`

### Enteros
- int8 (byte)
- int16 (short)
- int32 (int) -> Este es `int`
- int64 (long)

El literal por defecto es int32

En el caso de los enteros tenemos **sin signo** (positivos):
- uint8
- uint16
- uint32
- uint64

Por ejemplo: int8 -> X X X X X X X X (en cada X puede estar un 1 o un 0)
2^8 posibles valores.
256 posibles valores.
Debo repartir entre negativos, el cero y los positivos.
Rango [-128, 127]

En un uint8 siempre son 256 posibilidades pero solamente positivas
[0, 255]

### Flotante
- float32 (float en java)
- float64 (double en java)

### Caracteres
- rune (char en java)

### Cadenas
- string

El literal por defecto es float64


--|--x---x--x---x---x--x--|---
  0                       1

## Funciones

Para identificar los elementos de la funcion seria de la siguiente forma:
```func nombre-de-funcion (entrada1 tipo1, entrada2 tipo2) (tipo-salida1, tipo-salida2)```

En este caso:
```Atoi(s string) (int, error)```

Entradas:
- s string

Salidas
- int
- error