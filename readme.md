# Amazon product information downloader

### Installation

```bash
$ docker build . -t amazon_downloader:0.1
$ docker run -p 8083:8083 -e AMAZON_DOWNLOADER_PORT=8083 amazon_downloader:0.1
```


### Usage

Application API has 2 entrypoint:

POST \test\ - Add new task for execution and return task id 

Example response:

```json
{"value":{"id":"bc47fg5e4b3qcvootrc0","status":"done","urls":["https://www.amazon.co.uk/gp/product/1509836071","https://www.amazon.co.uk/gp/product/1509836071"],"result":[{"url":"https://www.amazon.co.uk/gp/product/1509836071","price":"£8.49","image":"","title":"The Fat-Loss Plan: 100 Quick and Easy Recipes with Workouts","is_sale":true},{"url":"https://www.amazon.co.uk/gp/product/1509836071","price":"£8.49","image":"","title":"The Fat-Loss Plan: 100 Quick and Easy Recipes with Workouts","is_sale":true}]},"error":"None","status":true}
```

GET \test\[task_id] - return status execution

Example response:

```json
{"value":{"id":"bc47q3le4b3sg4d0fu40","status":"done","urls":["https://www.amazon.co.uk/gp/product/B00IOY524S","https://www.amazon.co.uk/gp/product/B01BDNVY66"],"result":[{"url":"https://www.amazon.co.uk/gp/product/B00IOY524S","price":"£17.99","image":"https://images-na.ssl-images-amazon.com/images/I/61ta4nHKbVL._SL1000_.jpg","title":"Kindle Voyage E-reader, 6\" High-Resolution Display (300 ppi) with Adaptive Built-in Light, PagePress Sensors, Wi-Fi","is_sale":true},{"url":"https://www.amazon.co.uk/gp/product/B01BDNVY66","price":"None","image":"https://images-na.ssl-images-amazon.com/images/I/61RjeqZD5RL._SL1010_.jpg","title":"Bestdeal Wireless Bluetooth Game Controller Gamepad Joystick for Meizu Blue Charm \u0026 M1 note \u0026 M2 \u0026 M2 note \u0026 MX \u0026 MX2 \u0026 MX3 \u0026 MX4 \u0026 MX4 Pro \u0026 MX5 \u0026 MX5 Smartphone","is_sale":false}]},"error":"None","status":true}
```
