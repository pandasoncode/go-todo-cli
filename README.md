# gotodo - CLI para gestionar tareas (TODO)

**gotodo** es una herramienta de línea de comandos (CLI) escrita en Go que permite gestionar tareas de manera sencilla. Con **gotodo** puedes crear, visualizar y listar tareas con una interfaz amigable desde tu terminal.

## Instalación

Para instalar **gotodo**, primero necesitas tener instalado [Go](https://golang.org/dl/). Luego, puedes clonar este repositorio y compilar el binario:

```bash
git clone https://github.com/pandasoncode/go-todo-cli
cd gotodo
go build -o gotodo
```

## Uso

Una vez instalado, puedes utilizar el comando `gotodo` con los siguientes subcomandos:

### 1. Crear una tarea

Utiliza el subcomando `create` para crear una nueva tarea. Debes proporcionar un título, descripción y categoría.

```bash
gotodo create --title "Comprar comida" --description "Comprar verduras y frutas" --category "Personal"
```

### 2. Mostrar una tarea

Utiliza el subcomando `show` para mostrar la información completa de una tarea que coincida con el título proporcionado.

```bash
gotodo show --title "Comprar comida"
```

### 3. Listar tareas por categoría

Utiliza el subcomando `list` para listar todas las tareas que coincidan con una categoría.

```bash
gotodo list --category "Personal"
```

Si no se incluye la opción `--category`, el comando listará todas las tareas:

```bash
gotodo list
```

## Ejemplos

- Crear una tarea en la categoría "Trabajo":

  ```bash
  gotodo create --title "Revisar reporte" --description "Revisar y enviar el reporte semanal" --category "Trabajo"
  ```

- Mostrar la tarea "Revisar reporte":

  ```bash
  gotodo show --title "Revisar reporte"
  ```

- Listar todas las tareas en la categoría "Trabajo":

  ```bash
  gotodo list --category "Trabajo"
  ```
