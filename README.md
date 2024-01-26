# FamPay Backend Assignment

## Problem Statement

<details>
  <summary>Click to expand!</summary>

### Basic Requirements:

- Server should call the YouTube API continuously in background (async) with some interval (say 10 seconds) for fetching the latest videos for a predefined search query and should store the data of videos (specifically these fields - Video title, description, publishing datetime, thumbnails URLs and any other fields you require) in a database with proper indexes.
- A GET API which returns the stored video data in a paginated response sorted in descending order of published datetime.
- It should be scalable and optimised.

### Bonus Points:

- [x] Add support for supplying multiple API keys so that if quota is exhausted on one, it automatically uses the next available key.

- [x] Make a dashboard to view the stored videos with filters and sorting options (optional)

### Instructions:
- ou are free to choose any search query, for example: official, cricket, football etc. (choose something that has high frequency of video uploads)
- Try and keep your commit messages clean, and leave comments explaining what you are doing wherever it makes sense.
- Also try and use meaningful variable/function names, and maintain indentation and code style.
- Submission should have a `README` file containing instructions to run the server and test the API.
- Accepted language & Framework
    1. Python (DRF, Django, Flask, etc)
    2. GoLang
- Send your submission (Git repository) link at hiring@fampay.in

### Reference:
* [YouTube data v3 API](https://developers.google.com/youtube/v3/getting-started)
* [Search API reference](https://developers.google.com/youtube/v3/docs/search/list)
* To fetch the latest videos you need to specify these: ```type=video, order=date, publishedAfter=<SOME_DATE_TIME>```
Without publishedAfter, it will give you cached results which will be too old
</details>

## How it Works
- When a client requests data from Echo, using our REST API, we read the DB for latest data and send back a paginated JSON response.

- In the background, in Echo we have scheduled a goroutines periodic task which make a request to Youtube API and get the latest data.

- The response received is stored in the DB(SQL).

### API Endpoints
<details>
  <summary>Click to expand!</summary>

## Endpoints
  1. **Get Value**
     - Endpoint: `/v1/getVideos`
     - Method: `GET`
     - Parameters: `None`
     - Example: `curl -X GET "http://localhost:1323/v1/getVideos"`
     - Response: 
       ```json
        [{
            "title": "free fire dedicated to God sree RAM 🙏🙏🙏🙏🙏🙏🙏#SHORS # GAMMING #newsong",
            "description": "",
            "thumbnails": "https://i.ytimg.com/vi/r1Kfqj2mp2o/default.jpg",
            "published_at": "2024-01-25T03:33:02Z"
        },
        {
            "title": "Who else is playing in the fncs!! #fncs #tournament #fortnite #pc #gamming",
            "description": "",
            "thumbnails": "https://i.ytimg.com/vi/LGeDV7X2A50/default.jpg",
            "published_at": "2024-01-26T00:37:02Z"
        }]
       ```

  2. **Delete all data:**
     - Endpoint: `/v1/deleteVideos`
     - Method: `GET`
     - Parameters: `None`
     - Example: `curl -X GET "http://localhost:1323/v1/deleteVideos`
     - Response:
       ```json
       "Deleted"
       ```
</details>


## How to Run This Project

### Using Source Code (Recommended)

1. Clone the GitHub repository:
    ```
    $ git clone https://github.com/GoutamVerma/Fampay-Backend
    $ cd Fampay-Backend
    ```

2. Before starting the server, provide the necessary information in the `backend-server.yaml` file. This includes the credentials for the database (AWS RDS) and YouTube API keys. Example:

    ```
    username: "admin"
    password: "USERabc"
    hostname: "videos.cz154656546svs.ap-north-10.rds.amazonaws.com"
    port: 3306
    databaseName: "youtubeAPI"

    FetchInterval : 10
    MaxResult: 5
    PageSize: 10

    TotalAPIKeys: 1
    youtubeAPI1: "AIzaSyAiL3oNc4527523752452T8T_8r2Y2FaCSI"
    ```

3. Once you have provided the desired information, you can start the server.

4. Run the following command to start the Echo server:
    ```
    $ make
    $ make run
    ```

    OR

    ```
    $ go run cmd/delivery/main.go
    ```
5. You can access the API on port `localhost:1323`.

### Using the Dockerfile

To use the Dockerfile in this project, follow these steps:

1. Make sure you have the Docker installed on your system.

2. Open the `backend-server.yaml` file located in the root directory of the project. Provide the necessary information, such as the credits for AWS RDS or YouTube API key, in the `backend-server.yaml` file.

3. Once you have provided the required information, you can build the Docker image using the Dockerfile. Run the following command in the terminal:

    ```bash
    docker build -t fampay-server:latest .
    ```

    This command will build the Docker image with the tag `fampay-server:latest`.

4. After the Docker image is built, you can run it using the following command:

    ```bash
    docker run -p 1323:1323 fampay-server:latest
    ```

    This command will run the Docker container and map port 1323 of the container to port 1323 of the host machine.

Now you should be able to access your application running inside the Docker container at `http://localhost:1323`.
