#!/bin/bash

# Pastikan script berhenti jika ada error
set -e

# Hapus file todos.json agar mulai fresh
rm -f todos.json

echo "🔹 Testing ADD..."
go run main.go add "Belajar Golang"
go run main.go add "Main ke rumah ayank"
go run main.go add "Ngoding side project"
echo

echo "🔹 Testing LIST..."
go run main.go list
echo

echo "🔹 Testing COMPLETE..."
go run main.go complete 2
go run main.go list
echo

echo "🔹 Testing UPDATE..."
go run main.go update 3 "Ngoding REST API side project"
go run main.go list
echo

echo "🔹 Testing DELETE..."
go run main.go delete 1
go run main.go list
echo

echo "✅ Semua fitur berhasil diuji!"
