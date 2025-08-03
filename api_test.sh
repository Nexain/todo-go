#!/bin/bash

# Warna
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # Reset

print_pass() { echo -e "${GREEN}✅ $1${NC}"; }
print_fail() { echo -e "${RED}❌ $1${NC}"; }
print_info() { echo -e "${BLUE}ℹ️  $1${NC}"; }

BASE_URL="http://localhost:3000"

echo "🧹 Menghapus data lama..."
rm -f todos.json

echo
print_info "=== TEST CASE: SUKSES ==="

print_info "1. Tambah todo"
curl -s -X POST "$BASE_URL/todos" \
     -H "Content-Type: application/json" \
     -d '{"task":"Belajar Fiber"}' && print_pass "Todo 1 ditambahkan"

curl -s -X POST "$BASE_URL/todos" \
     -H "Content-Type: application/json" \
     -d '{"task":"Main ke rumah ayank"}' && print_pass "Todo 2 ditambahkan"

curl -s -X POST "$BASE_URL/todos" \
     -H "Content-Type: application/json" \
     -d '{"task":"Ngoding side project"}' && print_pass "Todo 3 ditambahkan"

echo
print_info "2. List todo"
curl -s "$BASE_URL/todos" && echo && print_pass "List berhasil ditampilkan"

echo
print_info "3. Complete todo ID 2"
curl -s -X PATCH "$BASE_URL/todos/2/complete" && echo && print_pass "Todo 2 selesai"

echo
print_info "4. Update todo ID 3"
curl -s -X PUT "$BASE_URL/todos/3" \
     -H "Content-Type: application/json" \
     -d '{"task":"Ngoding REST API side project"}' && echo && print_pass "Todo 3 diupdate"

echo
print_info "5. Delete todo ID 1"
curl -s -X DELETE "$BASE_URL/todos/1" && echo && print_pass "Todo 1 dihapus"

echo
print_info "=== TEST CASE: GAGAL ==="

print_info "1. Tambah todo kosong"
curl -s -f -X POST "$BASE_URL/todos" \
     -H "Content-Type: application/json" \
     -d '{"task":""}' && echo && print_fail "Seharusnya gagal" || print_pass "Error terdeteksi"

echo
print_info "2. Complete ID tidak valid"
curl -s -f -X PATCH "$BASE_URL/todos/abc/complete" && echo && print_fail "Seharusnya gagal" || print_pass "Error terdeteksi"

echo
print_info "3. Update ID tidak ada"
curl -s -f -X PUT "$BASE_URL/todos/99" \
     -H "Content-Type: application/json" \
     -d '{"task":"Testing"}' && echo && print_fail "Seharusnya gagal" || print_pass "Error terdeteksi"

echo
print_info "4. Delete ID tidak ada"
curl -s -f -X DELETE "$BASE_URL/todos/99" && echo && print_fail "Seharusnya gagal" || print_pass "Error terdeteksi"

echo
print_info "=== 🎯 TEST SELESAI ==="
