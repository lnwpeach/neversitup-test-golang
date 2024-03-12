### Project structure

    .
    ├── configs
    │   └── config.json
    │
    ├── cores
    │   ├── database.go
    │   └── routes.go
    │
    ├── controllers
    │   ├── <name>_controller.go
    │   └── <name>_controller_test.go
    │
    ├── services
    │   ├── <name>_service.go
    │   └── <name>_service_test.go 
    │
    ├── repositories
    │   ├── <name>_repository.go
    │   └── <name>_repository_test.go
    │
    ├── models
    │   └── <name>.go
    │
    ├── mock_object
    │   └── ...
    │
    ├── util
    │   └── util.go
    │
    ├── main.go
    │
    └── ... (Dockerfile, docker-compose, ...)


### configs
เก็บไฟล์ config ต่างๆ เช่น host, port, service name

### database.go
ส่วน Connect database

### routes.go
Setting api route ที่เปิดให้ใช้งาน

### controllers
เก็บ source code ส่วน controller ทำหน้าที่รับข้อมูลจากภายนอกแล้วแปลงเป็น Struct เพื่อใช้งาน รวมถึงการ Validation ต่างๆ แล้วส่งให้ระดับ service

### services
เก็บ source code ส่วน service นำข้อมูลจาก controller มาผ่าน Business Logic ต่างๆ ของระบบ ถ้าหากต้องการต่อกับ Database จะส่งข้อมูลต่อให้ระดับ repository

### repositories
เก็บ source code ส่วน repository ทำหน้าที่เชื่อมต่อกับ Database โดยรับข้อมูลจากระดับ Service มาใช้งาน โดยมี Database model เป็นตัวกำหนด Table structure

### models
เก็บ Database model, request struct, response struct เพื่อให้ส่วนอื่นใช้งาน

### mock_object
เก็บ Mock object ของ controllers, services, repositories

### util
เก็บ function ที่ใช้งานบ่อยๆ และเป็น function ทั่วไปที่ใช้ได้กับทุก Business logic
