# GoLink Checker

Comprobador concurrente de URLs, desarrollado con Go.

Consulta varias direcciones mediante goroutines y muestra estado HTTP, duración y destino final de redirecciones.

## Funciones actuales

- URLs desde argumentos, archivo de texto o ambos.
- Eliminación de argumentos duplicados exactamente iguales.
- Timeout configurable y señalización de respuestas lentas.
- Resumen con conteos, promedio y tiempos mínimo/máximo de respuestas saludables.
- Código de salida para integrarlo en scripts.

## Requisitos

Go 1.22 o superior y conexión a los destinos que se quieran comprobar.

## Instalación y uso

```bash
git clone https://github.com/DonMilo-22/GoLink-Checker.git
cd GoLink-Checker
go run . -timeout 5 -slow 1000 https://github.com https://example.com
```

Desde archivo:

```bash
go run . -file urls.example.txt
```

El archivo acepta una URL por línea; ignora líneas vacías y comentarios que comienzan con `#`.

## Opciones

Coloca las opciones antes de las URLs.

| Opción | Predeterminado | Uso |
|---|---|---|
| `-file <ruta>` | Sin archivo | Agrega URLs de un archivo de texto. |
| `-timeout <segundos>` | `5` | Tiempo máximo por solicitud; debe ser positivo. |
| `-slow <milisegundos>` | `1000` | Marca `SLOW` si la duración supera el umbral; admite cero. |

## Ejemplo de salida

Salida ilustrativa; tiempos y orden varían porque las solicitudes son concurrentes:

```text
✓ https://example.com                 200  120ms

Summary: 1 healthy | 0 failed | 0 slow | avg 120ms | fastest 120ms | slowest 120ms
```

## Estados y códigos de salida

Se considera saludable una respuesta HTTP entre 200 y 399. El símbolo `✓` indica que hubo una respuesta HTTP; consulta el número de estado y el resumen para saber si fue saludable. Se siguen redirecciones con el cliente HTTP de Go y se muestra el destino final cuando cambia.

| Código del ejecutable | Significado |
|---|---|
| `0` | Todas las comprobaciones fueron saludables; también se usa al mostrar uso sin URLs. |
| `1` | Alguna solicitud falló, hubo estado HTTP fuera de 200–399 o no se pudo leer el archivo. |
| `2` | Timeout/umbral inválido o error de análisis de opciones. |

Para conservar estos códigos en CI, compila y ejecuta el binario: `go run` puede presentar el código del programa a través de su propio error de ejecución.

```bash
go build -o golink-checker .
./golink-checker -file urls.example.txt
```

## Almacenamiento y límites

No guarda resultados. Agrega `https://` a entradas sin esquema. La deduplicación ocurre antes de esa normalización. La duración mide la solicitud hasta recibir la respuesta, no la descarga completa del contenido. El promedio incluye intentos fallidos; mínimo y máximo incluyen únicamente respuestas saludables. No verifica enlaces internos ni contenido de páginas.

## Comprobaciones

`go build .` comprueba la compilación. Para comprobar la conectividad, ejecuta el ejemplo contra destinos conocidos; sus resultados dependen de la red.

## Autor y licencia

Desarrollado por [DonMilo-22](https://github.com/DonMilo-22).

Licencia: MIT.

## Actualización 2026-10-12

- The final summary now includes how many checked URLs redirected to a different destination.

## Actualización adicional 2026-10-12

- URL results now include the HTTP status text beside the numeric status.

## Historial

Consulta [CHANGELOG.md](CHANGELOG.md) para ver las actualizaciones.
