# Lions Forum

Lions Forum is a web forum application made using Golang, HTML and CSS.
For the data storage, it uses SQLite.
The forum is a web portal where users can read, create, react and comment posts about any subject related to literature.

### Features:
	
    User Authentication: Secure user login and registration.
    Creation of Posts (with possibility to select up to three categories for each post).
    React to Posts and Comments (Like/Dislike).
    Comment Posts.
    Non-LoggedIn users have limited features but are still able to navigate through posts.
    Search feature allows users to search within the content and title of posts.
    Filters by category, by user, by trending and by untrending.
    Selectable avatar image.
    Pagination system for all pages.
    Simple breadcrumb (Example: Home / Profile)
    Error/Succesful informative messages for several actions
    Informative Logs on terminal

### Setup and Installation


### Run the application:

    go run main.go

The server will start running on http:/localhost:8080.



Alternatively, you can use Docker to set up and run the application:


Build the Docker image:
`docker build -t <image name> .`

Run the Docker container:
`docker run -p 8080:8080 [--name <optional container name>] -it <image name>`

After closing the application, remember to stop the container:
`docker container stop <container name>`

If you didn't specify a name for the container, you can check its name by:
`docker container ls`

To remove (delete) the container:
`docker container rm <container name>`

### Summarized Description:

This project is built using Go for the server side and SQLite for the database. The Front-End is made using only HTML and CSS. In order to render dynamid data in the HTML, it uses Go's built-in templating.

When the server receives a request, this one is sent to a Middleware where we check for cookies containing a valid open session for that user. After determining if user is logged in or not, the request is sent to its specific handler.

Then, within each handler, specific functions are called to gather the data needed from the data base and pack it in a final struct that will be send to the client. Along the process, informative log messages are printed in the terminal as well as errors.

Errors and anauthorized access to endpoints have been handled to inform the user, either by Redirecting them to personalized 404 and 500 pages or by redirecting them to other accessible pages of the web.
