# Go Project

To store sales report and fetching the revenue details using GO Program

- Used Golang for REST API and file handling
- GORM to insert and fetch data from DB

## Tables

- Customers
- Products
- Orders
- OrderDetails

  Refer Scheme.Sql for table structure

- To debug the process log file will be create for every execution and error code is added to tract the error.

- Gorountine added to refresh the Sales data in the data base by reading the file in the given path
- can also explicitly call the api to refresh the data

## API

### 1. Upload Sales Data

- **Endpoint:** `/ReloadSalesData`
- **Method:** `GET`
- **Description:** Read the CSV file from the path and insert into the respective tables.

- **Response:**

```json
{
  "status": "S",
  "errMsg": ""
}
```

---

### 5. Get Total Revenue by Date

- **Endpoint:** `/GetRevenue`
- **Method:** `POST`
- **Description:** Returns the revenue between a given date range.

## Request 1

- **Body:**

```json
{
  "indicator": "Total",
  "fromDate": "2022-01-01",
  "toDate": "2024-01-01"
}
```

- **Response: 1**

```json
{
  "status": "S",
  "errMsg": "",
  "totalRevenue": "583933.66",
  "prodRevenue": null,
  "catRevenue": null,
  "regRevenue": null
}
```

---

## Request 2

- **Body:**

```json
{
  "indicator": "Prod",
  "fromDate": "2022-01-01",
  "toDate": "2024-01-01"
}
```

- **Response: 2**

```json
{
  "status": "S",
  "errMsg": "",
  "totalRevenue": null,
  "prodRevenue": [
    {
      "prodname": "Mouse",
      "total_revenue": 25462.25
    },
    {
      "prodname": "KeyBoard",
      "total_revenue": 6247985.23
    }
  ],
  "catRevenue": null,
  "regRevenue": null
}
```

---

## Request 3

- **Body:**

```json
{
  "indicator": "Prod",
  "fromDate": "2022-01-01",
  "toDate": "2024-01-01"
}
```

- **Response: 3**

````json
{
  "status": "S",
  "errMsg": "",
  "totalRevenue": null,
  "prodRevenue": null,
  "catRevenue": [
    {
      "category": "Electronics",
      "total_revenue": 6544852.32
    },
    {
      "category": "Household",
      "total_revenue": 46554.24
    }
  ],
  "regRevenue": null
}
---
---

## Request 4

- **Body:**

```json
{
  "indicator": "Reg",
  "fromDate": "2022-01-01",
  "toDate": "2024-01-01"
}
````

- **Response: 3**

```json
{
  "status": "S",
  "errMsg": "",
  "totalRevenue": null,
  "prodRevenue": null,
  "catRevenue": null,
  "regRevenue": [
    {
      "region": "Chennai",
      "total_revenue": 6544852.32
    },
    {
      "region": "Mumbai",
      "total_revenue": 46554445.24
    }
  ]
}
```
