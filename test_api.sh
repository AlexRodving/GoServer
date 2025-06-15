#!/bin/bash

# Цвета для вывода
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Базовый URL
BASE_URL="http://localhost:8080"

echo -e "${BLUE}Тестирование API задач${NC}\n"

# 1. Регистрация
echo -e "${GREEN}1. Регистрация нового пользователя${NC}"
REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/register" \
-H "Content-Type: application/json" \
-d '{
    "username": "testuser",
    "password": "testpass123"
}')

echo "Ответ: $REGISTER_RESPONSE"
echo "----------------------------------------"

# 2. Авторизация
echo -e "\n${GREEN}2. Авторизация и получение токена${NC}"
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/login" \
-H "Content-Type: application/json" \
-d '{
    "username": "testuser",
    "password": "testpass123"
}')

echo "Ответ: $LOGIN_RESPONSE"
echo "----------------------------------------"

# Извлекаем токен из ответа
TOKEN=$(echo $LOGIN_RESPONSE | grep -o '"token":"[^"]*' | cut -d'"' -f4)

if [ -z "$TOKEN" ]; then
    echo "Ошибка: Не удалось получить токен"
    exit 1
fi

# 3. Создание задачи
echo -e "\n${GREEN}3. Создание новой задачи${NC}"
CREATE_TASK_RESPONSE=$(curl -s -X POST "$BASE_URL/api/tasks" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $TOKEN" \
-d '{
    "title": "Тестовая задача",
    "start_date": "2024-03-20T10:00:00Z",
    "end_date": "2024-03-21T18:00:00Z"
}')

echo "Ответ: $CREATE_TASK_RESPONSE"
echo "----------------------------------------"

# 4. Получение списка задач
echo -e "\n${GREEN}4. Получение списка задач${NC}"
GET_TASKS_RESPONSE=$(curl -s -X GET "$BASE_URL/api/tasks" \
-H "Authorization: Bearer $TOKEN")

echo "Ответ: $GET_TASKS_RESPONSE"
echo "----------------------------------------" 