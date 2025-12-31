# Seguridad TLS usando SSL Labs + Go
Este proyecto es una herramineta en Go que utiliza la API publica de **SSL LABS** para analizar la condición de seguridad de un dominio.

## Cómo ejecutar el proyecto

1. Clona el repositorio y asegúrate de tener **Go 1.20 o superior** instalado.
2. Ejecuta el siguiente comando desde la raíz del proyecto:

```bash
go run SSL_VALIDATION.go
```

## TLS Grades (SSL Labs)

Las calificaciones TLS reflejan el nivel de seguridad de la configuración SSL/TLS de un servidor.

### 🟢 Calificaciones Seguras
- 🟢 **A+** - Excelente seguridad (configuración moderna + HSTS)
- 🟢 **A** - Muy segura, recomendada para producción

### 🟡 Calificación Aceptable
- 🟡 **B** - Segura, pero con configuraciones antiguas

### 🟠 Calificaciones Débiles
- 🟠 **C** - Configuración débil, cifrados obsoletos
- 🔴 **D** - Muy débil, alto riesgo de seguridad

### 🔴 Calificaciones Críticas
- 🔴 **F** - Insegura o vulnerable
- ⚪ **T** - Problemas de confianza en el certificado
- ⚪ **M** - Contenido mixto (HTTPS + HTTP)


# Objetivo:
- Evaluar la seguridad TLS de un dominio
- Consumir una API REST externa (SSL labs)
- Manejar polling y estados asíncronos
- Parsear respuestas JSON

# Alcance:
- Análisis TLS mediante SSL Labs
- Obtención de calificación (A, B, C, etc.)
- Manejo de múltiples endpoints/IPs
- Ejecución vía CLI

# Requisitos del Sistema
**Requisitos Funcionales**
- El usuario debe poder introducir un dominio
- El sistema debe iniciar un análisis TLS
- El sistema debe consultar el estado del análisis
- El sistema debe mostrar la calificación TLS final

**Requisitos No Funcionales**
- Lenguaje: Go 1.20+
- API externa: SSL Labs v3
