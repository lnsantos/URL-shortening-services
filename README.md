# URL-shortening

this project is to practice me create services in golang, none done to production.

### How run

1. Clone project
2. Create on root folder file .env, add variables

```xml
   MONGO_URI="mongodb+srv://"
        ENV="dev"
```

3. Download dependencies
4. Run server

### API documentation

| Method | Endpoint          | Description        |
|--------|-------------------|--------------------|
| GET    | /api/url/{short}  | info short url     |
| POST   | /api/register/url | register new short |

### GET:

> Request
`http://127.0.0.1:8080/api/url/aHR0cHM6Ly9wa2cuZ28uZGV2L2VuY29kaW5nL2pzb24=`

> Response
>```json
> {
>  "original": "https://pkg.go.dev/encoding/json",
>  "short": "aHR0cHM6Ly9wa2cuZ28uZGV2L2VuY29kaW5nL2pzb24="
>} 
>```

### POST:
> Request
`http://127.0.0.1:8080/api/register/url`

> Body
> ```json 
> { "original": "https://pkg.go.dev/encoding/json" }
>```

> Response
>```json
> {
>  "short": "aHR0cHM6Ly9wa2cuZ28uZGV2L2VuY29kaW5nL2pzb24="
>} 
>```