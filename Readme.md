# Orders Manager

### Описание:

**Orders Manager** — это система управления заказами, предназначенная для автоматизации и оптимизации процессов обработки, учета и контроля заказов в бизнесе. Проект предоставляет удобный интерфейс для менеджеров, сотрудников и администраторов, позволяя эффективно управлять заказами, отслеживать их статусы, анализировать данные и формировать отчеты.  

#### Основные возможности:  
- **Создание и редактирование заказов** – добавление новых заказов, изменение их параметров (клиент, товары, сроки исполнения и т. д.).  
- **Управление статусами** – гибкая система контроля этапов выполнения заказа (новый, в обработке, выполнен, отменен и др.).  
- **Поиск и фильтрация** – быстрый поиск заказов по различным критериям (ID, клиент, дата, статус).  
- **Уведомления и оповещения** – автоматические уведомления о смене статуса, просроченных заказах и других важных событиях.  
- **Отчетность и аналитика** – генерация отчетов по продажам, загруженности сотрудников, популярным товарам и другим метрикам.  
- **Интеграции** – возможность подключения к CRM, ERP, платежным системам и другим сервисам.  

#### Целевая аудитория:  
- Малый и средний бизнес (интернет-магазины, сервисные компании, рестораны и др.).  
- Менеджеры по продажам и операционные сотрудники.  
- Администраторы и аналитики, работающие с данными заказов.  

**Orders Manager** помогает сократить время на рутинные операции, минимизировать ошибки и повысить прозрачность бизнес-процессов, связанных с обработкой заказов.

---

### Техническая часть:

### Стек:
- **Golang/Gin**
- **Postgres**
- **MongoDB**
- **Docker-compose**

### API:
- **REST**
- **gRPC**

### Защита:
- **mTLS/TLS**

### Дополнительно:
- **JWT**
- **CORS**
- **Hashing**

---

## Переменные окружения:
- **DATABASE_PASSWORD** - пароль от базы данных
- **ADMIN_ACCESS_KEY** - ключ для доступа к роли администратора

*Конфигурационные файлы каждого сервиса лежат внутри директории каждого из сервисов в ввиде `config.yaml`*