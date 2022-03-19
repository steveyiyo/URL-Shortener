# URL Shortener

About technical details, please see [details](details.md).

## Usage

Upload URL API:

1. "url" require a valid URL. (Eg: https://google.com/)
2. "expireAt" require RFC 3399 UTC format.

```
curl -X POST -H "Content-Type:application/json" http://localhost/api/v1/urls -d '{
    "url": "https://steveyi.net/",
    "expireAt": "2021-02-08T09:20:41Z"
}'
```

Response:

1. "id" will generate a 5 character random string.
2. "shortUrl" will generate a short URL

```
{
    "id": "<url_id>",
    "shortUrl": "http://localhost/<url_id>"
}
```

Redirect URL API:

```
curl -L -X GET http://localhost/<url_id> => REDIRECT to original URL # HTTP 301
```