# Тестовое задание для AvitoTech
## Микросервис для работы с балансом пользователей

Этот микросевис предоставляет методы для работы с балансом пользователей: пополнение баланса, перевод средств, резервирование средств, разрезервирование средств, получение списка транзакций пользователя с пагинацией и сортировкой по сумме и дате, предоставления отчета для бухгалтерии в формате .csv, признание выручки.

## Задание
    Основное задание :white_check_mark:
    Доп. задание 1 :white_check_mark:
    Доп. задание 2 :white_check_mark:
    Swagger File :white_check_mark:
    Сценарий разрезервирования денег:white_check_mark:

## Развертывание локально с помощью Docker

    docker build . -t api
    docker-compose build
    docker-compose up

## Описание методов
Swagger документация: http://localhost:8080/swagger/index.html

## USER BALANCE:
### Получение баланса:
По id возращает текущий баланс пользователя

     /api/ub/get_balance
  - POST http://localhost:8080/api/ub/get_balance
  - JSON:

        {
            "user_id": 0
        }

Ответ
  - JSON:

        {
            "balance": 0
        }

### Пополнение баланса:
Пополняет баланс пользователя
  - POST http://localhost:8080/api/ub/deposit_balance
  - JSON:

        {
            "user_id": 0, 
            "deposit": 0,
            "comment": "comment"
        }

Ответ:
  - JSON:

        {
            "err_code": 0,
            "err_msg": "string"
        } 

### Получение истории транзакций:
Получения списка транзакций пользователя с пагинацией и возможностью сортировки по цене и дате
  - GET http://localhost:8080/api/ub/history_transactions?user_id=0&page_no=1&page_size=1&sort=cost&ascending=true

        user_id = int > 0
        page_no = int > 0
        page_size = int > 0 (max 100)
        sort = string ("cost" or "date")
        ascending = bool (true = ASC, false = DESC) 

Ответ:
  - JSON:

        [
            {
                "transaction_id": 0,
                "type_transaction": "string"
                "cost": 0,
                "comment": "string",
                "date": "string",
            }
        ]

## SALE:

## Оплата услуги пользователем (резервирование средст):
Резервирует средства на отдельном счете до момента признания выручки

  - POST http://localhost:8080/api/sale/reserve
  - JSON:

        {
            "user_id": 0,
            "order_id": 0,
            "service_id": 0,
            "service_name": "string",
            "cost": 0
        }

Ответ:
  - JSON:

        {
            "err_code": 0,
            "err_msg": "string"
        } 

## Отмена услуги (возврат средст):
Переводит средства с резерного счета на счет пользователя
  - POST http://localhost:8080/api/sale/reserve_out
  - JSON:

        {
            "order_id": 0
        }

Ответ:
  - JSON:

        {
            "err_code": 0,
            "err_msg": "string"
        }
        

## Признание выручки:
Списывает деньги с резервного счета и сохраняет информация об успешной сделке для бухгалтерского отчета
  - POST http://localhost:8080api/sale/revenue
  - JSON:

        {
            "order_id": 0,
        }

Ответ:
  - JSON:

        {
            "err_code": 0,
            "err_msg": "string"
        }

## Получение отчета по сделкам:
Отправляет ссылку на скачивания отчета в формате .csv файла
  - POST http://localhost:8080/api/sale/sum_report
  - JSON:

        {
            "year": 2022,
            "month": 11
        }
Ответ:
  - JSON:

        {
            "url": "string"
        }

Формирует csv отчет о сделках за выбранный период и отправляет его
  - GET http://localhost:8080/api/sale/sum_report/download/{date}

        date = "11_2022"

Ответ:
  - CSV:

        service_id,service_name,cost


