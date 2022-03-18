# URL Shortener

```
# Upload URL API
curl -X POST -H "Content-Type:application/json" http://localhost/api/v1/urls -d '{
    "url": "<original_url>",                # Need a valid URL (Eg: https://google.com/)
    "expireAt": "2021-02-08T09:20:41Z"      # RFC 3399 UTC format
}'

# Response
{
    "id": "<url_id>",                       # will generate a 5 character random string
    "shortUrl": "http://localhost/<url_id>" # will generate a short URL
}

# ------------------
# Redirect URL API
curl -L -X GET http://localhost/<url_id> => REDIRECT to original URL # HTTP 301
```