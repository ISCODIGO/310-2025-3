## Pasos

1. Definir `go mod init nombre-modulo`. Por lo general, el nombre-modulo es la carpeta raíz de todo el módulo.
2. Cada subcarpeta es un paquete y debe definirse al inicio de cada archivo con `package nombre-paquete`.
3. Cuando se requiere llamar a un paquete, por ejemplo: `modulo/paquete1/archivo.go`, sería de la siguiente forma: `import "modulo/paquete1"`.
4. Se puede usar un alias para los imports: `import p1 "modulo/paquete1"` para usar `p1.Funcion()` en lugar de `paquete1.Funcion()`.

## Estructura del proyecto

- `main.go` - Archivo principal del programa
- `paquete1/` - Contiene funciones aritméticas y de comparación
- `paquete2/` - Contiene la definición y funciones relacionadas con `Persona`