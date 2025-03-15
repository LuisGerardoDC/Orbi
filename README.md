# Orbi Techtest

## Descripción del flujo de comunicación entre servicios

## Resumen
Comportamiento de la interaccion entre  **UserService** y **NotificationService**. Se detallan las acciones que se realizan al llamar a los microservicios, prerequisitos, instalacion, ejecucion 

## Flujo de proceso

### 1. Crear Usuario - `POST /user-service/user/`
Cuando se realiza una solicitud **POST** al endpoint `/user-service/user/` con un cuerpo en formato JSON para crear un nuevo usuario, se ejecutan las siguientes acciones:

- El usuario se **añade a la base de datos** del servicio UserService.
- Se envía una **notificación por medio de gRPC** al servicio **NotificationService**.
- El servicio NotificationService, al recibir la notificación, consulta el endpoint **GET `/user-service/user/:id`** para obtener los datos del usuario recién creado.
- Una vez obtenidos los datos, NotificationService simula el envío de un correo electrónico:
  - Si la carpeta `Notification_Service/notification/carpetanueva` no existe, **se crea automáticamente**.
  - Se genera un archivo **HTML** con el contenido del "email" y se guarda dentro de la carpeta mencionada.

### 2. Actualización o Eliminación de Usuario - `DELETE /user-service/user/:id` o `PUT /user-service/user/`
Cuando se realiza una solicitud **DELETE** `/user-service/user/:id` o **PUT** a `/user-service/user/`, se lleva a cabo el siguiente proceso:

- Se actualiza o marca como eliminado en la base de datos
- Se **envía un mensaje de actualización** (ya sea eliminación o modificación) a **RabbitMQ**.
- El servicio **NotificationService** consume del mensaje enviado en RabbitMQ.
- Al recibir el mensaje, NotificationService realiza un **log del mensaje recibido**, procesando la actualización o eliminación del usuario según el caso.


## Servicios
- **User Service**
- **Notification Service**
- **MySQL**
- **RabbitMQ**

## Requisitos Previos
Antes de proceder con la instalación, asegúrese de tener instalados los siguientes requisitos:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Postman](https://www.postman.com/)

## Método de Instalación
Siga los siguientes pasos para desplegar el proyecto:

1. Clonar el repositorio:
   ```sh
   git clone git@github.com:LuisGerardoDC/Orbi.git
   cd Orbi
   ```
2. En la raíz del proyecto, ejecutar el siguiente comando:
   ```sh
   docker-compose up --build -d
   ```

## Endpoints de User Service

### **Crear Usuario**
**POST** `/user/`

#### **Solicitud (JSON)**
```json
{
    "name": "John Doe",
    "email": "correo@correo.com"
}
```

#### **Respuesta (JSON)**
```json
{
    "success": true,
    "message": "Usuario creado exitosamente",
    "user": {
        "id": 1,
        "name": "John Doe",
        "email": "correo@correo.com",
    }
}
```

---

### **Obtener Usuario por ID**
**GET** `/user/:id`

#### **Respuesta (JSON)**
```json
{
    "success": true,
    "message": "Usuario encontrado",
    "user": {
        "id": 1,
        "name": "John Doe",
        "email": "correo@correo.com",
        "deletedAt": null
    }
}
```

---

### **Actualizar Usuario**
**PUT** `/user/`

#### **Solicitud (JSON)**
```json
{
    "name": "John Doe",
    "email": "nuevo_correo@correo.com"
}
```

#### **Respuesta (JSON)**
```json
{
    "success": true,
    "message": "Usuario actualizado exitosamente",
    "user": {
        "id": 1,
        "name": "John Doe",
        "email": "nuevo_correo@correo.com",
    }
}
```
---

### **Borrar Usuario**
**DELEYE** `/user/:id`

#### **Respuesta (JSON)**
```json
{
    "success": true,
    "message": "Usuario actualizado exitosamente",
    "user": {
        "id": 1,
    }
}
```
