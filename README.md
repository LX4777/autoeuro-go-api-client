# Api клиент на Go для веб-сервиса Авто-Евро

Установка:
```shell
go get github.com/LX4777/autoeuro-go-api-client@latest   
```

Пример использования:
```go
import (
    "encoding/json"
    "fmt"
    "github.com/LX4777/autoeuro-go-api-client/client"
    "github.com/LX4777/autoeuro-go-api-client/service"
    "time"
)

autoeuroService := service.NewAutoeuroService(client.ApiClientConfig{
    BaseURL: "https://api.autoeuro.ru/api/v2/json",
    Timeout: 2 * time.Second,
    Token:   "YOUR_TOKEN",
})

res, err := autoeuroService.GetBalance()

res, err := autoeuroService.GetDeliveries()

res, err := autoeuroService.GetWarehouses(requests.GetWarehousesRequestData{DeliveryKey: "YOUR DELIVERY KEY"})

res, err := autoeuroService.GetPayers()

res, err := autoeuroService.GetBrands()

res, err := autoeuroService.SearchBrands(requests.SearchBrandsRequestData{Code: "WDK962/12"})

res, err := autoeuroService.SearchItems(requests.SearchItemsRequestData{
    Code:        "lf16045",
    Brand:       "fleetguard",
    DeliveryKey: "YOUR DELIVERY KEY",
})

//Данный метод я не тестировал
res, err := autoeuroService.CreateOrder(requests.CreateOrderRequestData{})

res, err := autoeuroService.GetOrders(requests.GetOrdersRequestData{})

res, err := autoeuroService.GetStatuses()

if res != nil {
    str, _ := json.Marshal(res)
    fmt.Println(string(str))
} else {
    fmt.Println(err)
}
```

## Примечания
- Официальная документация [здесь](https://api.autoeuro.ru/doc/v2)
- Есть неточности в документации, которые я заметил:
    - В ответе `search_items` свойства `price`, `return`, `cross` по документации должны иметь тип `number`, а по факту `string`
    - В ответе `get_deliveries` свойство `time_shift_msk` по документации возвращает `number`, а по факту `string`
    - В ответе `get_statuses` свойство `status_id` по документации возвращает `number`, а по факту `string`

