THAI Version : 
# คู่มือการเรียนรู้ Microservices ด้วยภาษา Go

ยินดีต้อนรับสู่โปรเจกต์การเรียนรู้ Microservices ด้วยภาษา Go! คู่มือนี้ถูกออกแบบมาเพื่อช่วยให้คุณเข้าใจแนวคิดและรายละเอียดการพัฒนา Microservices โดยใช้ภาษา Go และ Framework **Fiber**  

---

## โครงสร้างโปรเจกต์

การเข้าใจโครงสร้างของโปรเจกต์เป็นสิ่งสำคัญสำหรับการนำทางและทำความเข้าใจโค้ดเบส ต่อไปนี้คือรายละเอียด:

```
.
├── cmd/ # จุดเริ่มต้น (Entry Point) ของแอปพลิเคชัน
│   ├── gateway/ # บริการ API Gateway
│   ├── service1/ # บริการ Hello
│   ├── service2/ # บริการ World
│   └── calculator-service/ # บริการคำนวณ
└── internal/ # โค้ดหลักของแอปพลิเคชัน
    ├── core/ # โลจิกทางธุรกิจ (Business Logic)
    │   ├── ports/ # การกำหนด Interfaces
    │   └── services/ # การทำงานของ Service
    └── handlers/ # HTTP Handlers
```

---

## แนวคิดสำคัญ

### Microservices
- **ความหมาย**: บริการย่อยที่เป็นอิสระจากกัน ทำงานร่วมกันเพื่อสร้างแอปพลิเคชันแบบสมบูรณ์
- **ข้อดี**: สามารถขยายได้ง่าย ยืดหยุ่น และสะดวกในการนำไปใช้งาน

### API Gateway
- **หน้าที่**: เป็นจุดเริ่มต้นที่คอยจัดการการเรียกใช้งานและส่งคำขอไปยังบริการที่เหมาะสม
- **ข้อดี**: ช่วยลดความซับซ้อนของการเชื่อมต่อ และสามารถจัดการเรื่อง Authentication หรือ Logging ได้

### Clean Architecture
- **หลักการ**: แยกความรับผิดชอบ (Separation of Concerns) และจัดการการพึ่งพิง (Dependency Management)
- **เป้าหมาย**: ทำให้ Business Logic เป็นอิสระจากการทำงานส่วนอื่น เช่น การรับส่งข้อมูล

---

## วัตถุประสงค์การเรียนรู้

1. **การสร้าง Microservices ด้วย Go**
   - เข้าใจวิธีการจัดโครงสร้างโปรเจกต์สำหรับ Microservices
   - เรียนรู้การใช้ Framework **Fiber** สำหรับจัดการ HTTP Requests

2. **การประยุกต์ใช้ Clean Architecture**
   - สำรวจวิธีแยก Business Logic ออกจากส่วนอื่น
   - ใช้ Interfaces เพื่อกำหนดสัญญาการทำงาน (Service Contracts)

3. **การสื่อสารระหว่าง Services ด้วย HTTP**
   - เรียนรู้วิธีที่บริการต่างๆ สื่อสารกันผ่าน HTTP Requests

---

## ขั้นตอนการเรียนรู้

### ขั้นตอนที่ 1: เข้าใจโครงสร้างโปรเจกต์
- ศึกษาโครงสร้างโฟลเดอร์และหน้าที่ของแต่ละโฟลเดอร์

### ขั้นตอนที่ 2: ส่วนประกอบหลัก
- **Ports (Interfaces)**: กำหนดสัญญาที่ Services ต้องปฏิบัติตาม
  ```go
  // ตัวอย่าง: internal/core/ports/service.go
  ```

- **Services (Implementation)**: ตำแหน่งที่ Business Logic ถูกนำมาประยุกต์ใช้งาน
  ```go
  // ตัวอย่าง: internal/core/services/hello.go
  ```

### ขั้นตอนที่ 3: HTTP Handlers
- เชื่อมโยง Services กับ HTTP Endpoints
  ```go
  // ตัวอย่าง: internal/handlers/hello.go
  ```

### ขั้นตอนที่ 4: รันบริการต่างๆ
- รันบริการแต่ละตัวในหน้าต่าง Terminal ที่แยกกัน:
  ```bash
  # Terminal 1
  go run cmd/service1/main.go

  # Terminal 2
  go run cmd/service2/main.go

  # Terminal 3
  go run cmd/gateway/main.go
  ```

---

## แบบฝึกหัดปฏิบัติ

1. **เพิ่มฟีเจอร์ใหม่**
   - สร้างบริการใหม่ที่ส่งข้อความ "Goodbye"
   - เพิ่ม Interface, Service และ HTTP Handler
   - เขียน Unit Test สำหรับบริการใหม่

2. **ปรับปรุง Gateway**
   - เพิ่มการจัดการข้อผิดพลาดและ Logging
   - เพิ่มการทำงาน Authentication แบบพื้นฐาน

3. **การสื่อสารระหว่าง Services**
   - ปรับบริการให้สามารถเรียกใช้งานกันเอง
   - สร้าง Endpoint ที่รวบรวมผลลัพธ์จากหลายบริการ

---

## กลยุทธ์การทดสอบ

1. **Unit Tests**
   - ทดสอบส่วนประกอบต่างๆ แยกกัน
   ```go
   // ตัวอย่าง: internal/core/services/hello_test.go
   ```

2. **Integration Tests**
   - ทดสอบการทำงานร่วมกันของส่วนประกอบ
   ```go
   // ตัวอย่าง: internal/handlers/hello_test.go
   ```

---

## ขั้นตอนถัดไป

1. **ฟีเจอร์ขั้นสูง**
   - เพิ่มการเชื่อมต่อฐานข้อมูลและการทำ Caching
   - เพิ่มระบบ Authentication/Authorization

2. **การเตรียมใช้งานใน Production**
   - ใช้ Docker สำหรับ Containerization
   - ตั้งค่า Kubernetes สำหรับการ Deploy
   - ติดตั้งระบบ Monitoring และ Logging

---

### ขอให้เรียนรู้อย่างสนุก! 😊

