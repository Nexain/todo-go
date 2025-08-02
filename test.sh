#!/bin/bash

# Warna
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # Reset

# Fungsi print
print_pass() { echo -e "${GREEN}✅ $1${NC}"; }
print_fail() { echo -e "${RED}❌ $1${NC}"; }
print_info() { echo -e "${BLUE}ℹ️  $1${NC}"; }
print_warn() { echo -e "${YELLOW}⚠️  $1${NC}"; }

set -e

print_info "Membersihkan data lama..."
rm -f todos.json

echo
print_info "=== TEST CASE: SUKSES ==="

print_info "1. ADD Todo"
go run main.go add "Belajar Golang" && print_pass "Todo 1 berhasil ditambahkan"
go run main.go add "Main ke rumah ayank" && print_pass "Todo 2 berhasil ditambahkan"
go run main.go add "Ngoding side project" && print_pass "Todo 3 berhasil ditambahkan"

echo
print_info "2. LIST Todo"
go run main.go list && print_pass "List berhasil ditampilkan"

echo
print_info "3. COMPLETE Todo ID 2"
go run main.go complete 2 && print_pass "Todo 2 berhasil diselesaikan"

echo
print_info "4. UPDATE Todo ID 3"
go run main.go update 3 "Ngoding REST API side project" && print_pass "Todo 3 berhasil diupdate"

echo
print_info "5. DELETE Todo ID 1"
go run main.go delete 1 && print_pass "Todo 1 berhasil dihapus"

echo
print_info "=== TEST CASE: GAGAL ==="

print_info "1. ADD dengan task kosong (harus error)"
go run main.go add "" && print_fail "Seharusnya gagal" || print_pass "Error terdeteksi"

echo
print_info "2. COMPLETE ID tidak valid (string)"
go run main.go complete abc && print_fail "Seharusnya gagal" || print_pass "Error terdeteksi"

echo
print_info "3. COMPLETE Todo yang sudah selesai"
go run main.go complete 2 && print_fail "Seharusnya gagal" || print_pass "Error terdeteksi"

echo
print_info "4. UPDATE ID tidak ada"
go run main.go update 99 "Testing error" && print_fail "Seharusnya gagal" || print_pass "Error terdeteksi"

echo
print_info "5. DELETE ID tidak ada"
go run main.go delete 99 && print_fail "Seharusnya gagal" || print_pass "Error terdeteksi"

echo
print_info "=== 🎯 TEST SELESAI ==="
