# Orbi Techtest

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
