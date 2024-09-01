# Api клиент на Go для веб-сервиса Авто-Евро

Установка из npm:
```shell
//
```

Пример использования:
```go
service := service2.NewAutoeuroService(client.ApiClientConfig{
BaseURL: "https://api.autoeuro.ru/api/v2/json",
Timeout: 2 * time.Second,
Token:   "YOUR TOKEN",
})

res, err := service.GetBalance();

res, err := service.GetDeliveries();

res, err := service.GetWarehouses(requests_data.GetWarehousesRequestData{DeliveryKey: "YOUR DELIVERY KEY"});

res, err := service.GetPayers()

res, err := service.GetBrands()

res, err := service.SearchBrands(requests_data.SearchBrandsRequestData{Code: "lf16045"})

res, err := service.SearchItems(requests_data.SearchItemsRequestData{Code: "lf16045", Brand: "fleetguard", DeliveryKey: "YOUR DELIVERY KEY"})

//Данный метод я не тестировал
res, err := service.CreateOrder(requests_data.CreateOrderRequestData{})

res, err := service.GetOrders(requests_data.GetOrdersRequestData{})

res, err := service.GetStatuses()

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

