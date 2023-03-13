
# URL Shortener

A URL shortener service creates an alias or a short URL for a long URL. Users are redirected to the original url when they visit thest short links

For example, the following long URL can be changed to a shorter URL

+ [Long URL](https://karanpratapsingh.com/courses/system-design/url-shortener)
+ [Short URL](https://bit.ly/3I71d3o)

Why do we need a URL shortener? URL shortener saves spaces when we are sharing URLs. Users're also less likely to mistype shorter URLs. Moreover, we can also optimize links across devices, this allows us to track individual links

## Requirements

Our URL shortenering system should meet the following requirements

+ Requirements{.mindmap}
    + Functional requirements
        + Given a URL, our service should generate a shorter and unique alias for it
        + Users should be redirected to the original URL when they visit the short link
        + Links should expire after a default timespan
    + Non-functional requirements
        + High availability with minimal latency
        + The system should be scalable and efficient
    + Extended requirements
        + Prevent abuse of services
        + Record analytics and metrics for redirections

## Estimation and Constraints

> *Make sure to check any scale or traffic related assumptions with your interviewer*

### Traffic

This will be a read-heavy system, so let's assume a `100:1` read/write ratio with 100 million links generated per month

For reads per month:
$$100 * 100 \text{ million} = 10 \text{ billion/month}$$
Similarity for writes:
$$1 * 100 \text{ million} = 100 \text{ million/month}$$

100 million requests per month translate into 40 requests per second
$$\frac{100 \text{ million}}{(30 \text{ days} * 24 \text { hours} * 3600 \text{ seconds})} \approx 40 \text{ URLs/second}$$

And with a `100:1` read/write ratio, the number of redirections will be:
$$100 * 40 \text{ URLs/second} = 4000 \text{ requests/second}$$

### Bandwidth

If we assume each request is of size 500 bytes then the total incoming data for write requests would be
$$40 * 500 \text{ bytes} \approx 20 \text{ KB/second}$$
Similarity, for the read requests, since we expect about 4K redirections, the total outgoing data would be:
$$4000 \text{ URLs/second} * 500 bytes \approx 2 MB/second$$

### Storage

For storage, we'll assume we store each link or record in our databases for 10 years. Since we expect around 100M new requests every month, the total number of records we will need to store would be:
$$100 \text{ million} * 10 \text{ years} * 12 \text{ months} = 12 \text{ billion}$$
Like earlier, if we assume each stored record will be approximately 500 bytes, we will need around 6TB of storage:
$$12 \text{ billion}*500\text{ bytes} \approx 6TB$$

### Cache

For caching, we will follow the classic 80/20 rule: **80% of the requests are for 20% of the data**. Since we get around 4K read or redirection requests each second, this translates into 350M requests per day
$$4000 \text{ URLs/second} * 24 \text { hours} * 3600 \text{ seconds} \approx 350 \text{ million requests/day}$$
Hence, we will need around 35GB of memory per day
$$20\% * 350 \text{ million} * 500 \text{ bytes} \approx 35 \text{ GB/day}$$

Here's our high-level estimates

| Type | Estimate |
| -- | -- |
| Writes(New URLs) | 40/s |
| Reads(Redirection) | 4k/s |
| Bandwidth(Incoming) | 20 KB/s |
| Bandwidth(Outgoing) | 2 MB/s |
| Storage(10 years) | 6TB |
| Memory(Caching) | ~35 GB/day |

## Data model design

![data model](./FILES/chapter-5.md/d63fd9e9.png)

## API design

:::: group URL Shortening API design
::: group-item Create URL
This API should create a new short URL in our system given an original URL

```go
createURL(apiKey: string, originalURL: string, expiration: Date): string
```

| Parameters/Return | Type | Description |
| :--: | :--: | :--: |
| API Key | string | API key provided by the user |
| Original URL | string | Original URL to be shortened |
| Expiration | Date | Expiration date of the new URL(optional) |
| Short URL | string | New shortened URL |
:::
::: group-item Get URL
This API should retrieve the original URL from a given short URL

```go
getURL(apiKey: string, shortURL: string): string
```

| Parameters/Return | Type | Description |
| :--: | :--: | :--: |
| API Key | string | API key provided by the user |
| Short URL | string | Short URL mapped to the original URL |
| Original URL | string | Original URL to be retrieved |
:::
::: group-item Delete URL
This API should delete a given shortURL from our system

```go
deleteURL(apiKey: string, shortURL: string): bool
```

| Parameters/Return | Type | Description |
| :--: | :--: | :--: |
| API Key | string | API key provided by the user |
| Short URL | string | Short URL to be deleted |
| Result | bool | Represents whether the operation was successful or not |
:::
::::

::: tip Why do we need an API key?
As you must've noticed, we're using an API key to prevent abuse of our services. Using the API key we can limit the users to a certain number of requests per second or minute. This's quite a standard practice for developer APIs and should cover our extended requirements
:::

## High-level design

### Caching

We can use Redis or Memcached servers alongside our API server

For more details, refer to [caching](#Cache)

### Desgin

![system-design-draft](./FILES/chapter-5.md/0de42359.png)

:::: group How it works
::: group-item Creating a new URL

1. When a user creates a new URL, our API server requests a new unique key from the Key Generation Service(KGS)
1. Key Generation Service provides a unique key to the API server and marks the key as used
1. API server writes the new URL entry to the database and cache
1. Our service returns an HTTP 201(Created) response to the user
:::
::: group-item Accessing a URL
1. When a client navigates to a certain short URL, the request is sent to the API servers
1. The request first hits the cache, and if the entry is not found there then it's retrieved from the database and an HTTP 301(Redirect) is issued to the original URL
1. If the key is still not found in the database, an HTTP 404(Not Found) error is sent to the user
:::
::::


## Detail design

### Data Partitioning

### Database cleanup

If we do decide to remove expired entries, we can approach this in two different ways

:::: group Database cleanup
::: group-item Active cleanup
In active cleanup, we will run a separate cleanup service which will periodically remove expired links from our storage and cache. This will be a very lightweight service like a cronjob
:::
::: group-item Passive cleanup
For passive cleanup, we can remove the entry when a user tries to access an expired link. This can ensure a lazy cleanup of our database and cache
:::
::::

### Cache design

:::: group Cache
::: group-item Cache eviction policy
Least Recently Used(LRU) can be a good policy for our system. In this policy, we discard the least recently used key first
:::
::: group-item Cache miss
Whenever there's a cache miss, our servers can hit the database directly and update the cache with the new entires
:::
::::

### Metrics and Analytics

Recording analytics and metrics is one of our extended requirements. We can store and update metadata like visitor's country, platform, the number of views, etc alongside the URL in our database

### Security

For security, we can introduce private URLs and authorization. A separate table can be used to store user ids that have permission to access a specific URL. If a user does not have proper permissions, we can return an HTTP 401(Unauthorized) error

We can also use an API Gateway as they can support capabilities like authorization, rate limiting, and load balancing out of the box

## Identify and resolve bottlenecks

![Img](./FILES/chapter-5.md/6033aee7.png)


