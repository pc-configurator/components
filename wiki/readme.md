блабла
ещё и ещё
Примечания для себя:
Prices будет в отдельной таблице, будет ссылаться на Components
Загрузка файла новых товаров и актуальных цен

Возможные категории:

enum Category {
MemoryRAM,
MemorySSD,
MemoryHHD,
CPU,
Accessories
Motherboard
GraphicsCard
Power
Cool
Case
Display
}

Сущности и enum:

1. сущность Component {
   id string,
   name string,
   price int,
   category: string,
   description: string
   }

Эндпоинты:

1. GET /components/{id}

Тело ответа: сущность Component

2. POST /components
   Тело запроса: все поля сущности Component обязательны, кроме поля id (поля id вообще не должно присутствовать)
   Тело ответа: { id: string }

3. PATCH /components/{id}
   Тело запроса: все поля сущности Component опциональны, кроме поля id (поле id мы передаем в url)

4. DELETE /components/{id}

5. GET /components/all?category={string}
   Тело ответа: массив сущностей Component
