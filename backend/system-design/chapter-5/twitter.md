# Twitter

Twitter is a social media service where users can read or post short message (up to 280 characters) called tweets. It's available on the web and mobile platforms such as Android and IOS

## Requirements

Our system should meet the following requirements:

+ Requirements{.mindmap}
    + Functional requirements
        + Should be able to post new tweets(can be text, image, video, etc)
        + Should be able to follow other users
        + Should have a newsfeed feature consisting of tweets from the people the user is following
        + Should be able to search tweets
    + Non-Functional requirements
        + High availability with minimal latency
        + The system should be scalable and efficient
    + Extended requirements
        + Metrics and analytics
        + Retweet functionality
        + Favorite tweets

## Estimation and Constraints

## API design

:::: group URL Tweets API design
::: group-item Post a tweet
This API will allow the user to post a tweet on the platform

```go
postTweet(userID: UUID, content: string, mediaURL?: string): boolean
```

| Parameters/Return | Type | Description |
| :--: | :--: | :--: |
| User ID | UUID | ID of the user |
| Content | string | Contents of the tweet |
| Media URL | string | URL of the attached media(optional) |
| Result | string | New shortened URL |
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