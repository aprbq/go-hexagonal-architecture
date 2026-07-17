APP_NAME := bank
DB_NAME  := banking

.PHONY: help run build tidy test fmt vet db-create db-migrate db-reset clean

help: ## แสดงรายการคำสั่งทั้งหมด
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'

run: ## รันเซิร์ฟเวอร์ (go run main.go)
	go run main.go

build: ## build binary ออกมาเป็น bin/$(APP_NAME)
	go build -o bin/$(APP_NAME) .

tidy: ## จัดการ dependencies (go mod tidy)
	go mod tidy

test: ## รันเทสทั้งหมด
	go test ./...

fmt: ## จัดรูปแบบโค้ด
	go fmt ./...

vet: ## ตรวจโค้ดด้วย go vet
	go vet ./...

db-create: ## สร้าง database $(DB_NAME)
	createdb $(DB_NAME)

db-migrate: ## รัน migration (migrations/init.sql)
	psql -d $(DB_NAME) -f migrations/init.sql

db-reset: ## ลบแล้วสร้าง database ใหม่พร้อมรัน migration
	dropdb --if-exists $(DB_NAME)
	createdb $(DB_NAME)
	psql -d $(DB_NAME) -f migrations/init.sql

clean: ## ลบไฟล์ build
	rm -rf bin
