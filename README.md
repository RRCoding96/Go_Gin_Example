# Go Gin 게시판 API

간단한 게시판 API 서버입니다. Go 언어와 Gin 웹 프레임워크를 사용하여 구현되었습니다.

## 기술 스택

- Go 1.21
- Gin Web Framework
- GORM
- MySQL

## 프로젝트 구조

```
.
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handler/
│   │   │   ├── health_handler.go
│   │   │   └── post_handler.go
│   │   └── router/
│   │       └── router.go
│   ├── config/
│   │   ├── config.go
│   │   └── database.go
│   ├── domain/
│   │   └── post.go
│   ├── repository/
│   │   └── post_repository.go
│   └── service/
│       └── post_service.go
├── go.mod
└── go.sum
```

## 주요 기능

- 게시글 CRUD (생성, 조회, 수정, 삭제)
- Health Check 엔드포인트
- MySQL 데이터베이스 연동

## API 엔드포인트

| Method | Endpoint      | Description           |
|--------|---------------|-----------------------|
| GET    | /health       | 서버 상태 확인        |
| GET    | /posts        | 모든 게시글 조회      |
| GET    | /posts/:id    | 특정 게시글 조회      |
| POST   | /posts        | 새 게시글 생성        |
| PUT    | /posts/:id    | 게시글 수정          |
| DELETE | /posts/:id    | 게시글 삭제          |

## 시작하기

### 필수 조건

- Go 1.21 이상
- MySQL 8.0 이상

### 설치

1. 저장소 클론
```bash
git clone https://github.com/RRCoding96/Go_Gin_Example.git
```

2. 의존성 설치
```bash
go mod download
```

3. 환경 설정
- MySQL 데이터베이스 생성
- `internal/config/config.go` 파일에서 데이터베이스 설정 수정

4. 서버 실행
```bash
go run cmd/api/main.go
```

서버는 기본적으로 8080 포트에서 실행됩니다.

## 데이터베이스 스키마

### Posts 테이블

| Column    | Type      | Description           |
|-----------|-----------|-----------------------|
| id        | uint      | Primary Key          |
| title     | string    | 게시글 제목          |
| content   | text      | 게시글 내용          |
| author    | string    | 작성자              |
| created_at| timestamp | 생성 시간           |
| updated_at| timestamp | 수정 시간           |
