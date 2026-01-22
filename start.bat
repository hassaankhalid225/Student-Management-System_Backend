@echo off
echo Starting Student Management System Backend...
echo.
echo Using SQLite database (development mode)
echo Database file: sms.db
echo.
go run cmd/server/main.go
