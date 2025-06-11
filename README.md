# Go RESTful MVC Project

Project RESTful API sử dụng Go với kiến trúc MVC, tích hợp Kafka cho xử lý bất đồng bộ và gửi email.

## Cài đặt

1. Clone repository:

```bash
git clone <repository-url>
cd go_restful_mvc
```

2. Cài đặt dependencies:

```bash
go mod download
```

## Cấu hình

1. Tạo file `.env` trong thư mục gốc của project:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=your_database

# Kafka Configuration
KAFKA_BROKERS=localhost:9092

# SMTP Configuration
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your-email@gmail.com
SMTP_PASSWORD=your-app-password
SMTP_FROM=your-email@gmail.com

# Application Configuration
APP_PORT=8080
```

## Chạy Kafka với Docker

1. Di chuyển vào thư mục kafka-docker:

```bash
cd kafka-docker
```

2. Khởi động Kafka và Zookeeper:

```bash
docker-compose up -d
```

3. Kiểm tra logs để đảm bảo Kafka đang chạy:

```bash
docker-compose logs -f
```

4. Để dừng Kafka:

```bash
docker-compose down
```

## Chạy ứng dụng

```bash
go run main.go
```

Server sẽ chạy tại `http://localhost:8080`

## API Endpoints

### Authentication

- Đăng ký: `POST /auth/register`
- Đăng nhập: `POST /auth/login`
- Cập nhật thông tin: `PUT /auth/user/:id`

### Products

- Tạo sản phẩm: `POST /products`
- Lấy danh sách: `GET /products`
- Lấy chi tiết: `GET /products/:id`
- Cập nhật: `PUT /products/:id`
- Xóa: `DELETE /products/:id`

## Cấu trúc thư mục

```
.
├── config/         # Cấu hình database và ứng dụng
├── controllers/    # Xử lý request/response
├── models/         # Định nghĩa model
├── repositories/   # Tương tác với database
├── routes/         # Định nghĩa routes
├── services/       # Business logic
├── kafka-docker/   # Cấu hình Docker cho Kafka
└── main.go         # Entry point
```
