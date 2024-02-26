# CRM Application

This is a scenario that mimics a Customer Relationship Management application and supports querying various endpoints to retrieve data from a mock database and manipulating the data within it.

# Instructions

After unzipping the file, kindly change directory into the CRM_Backend_Project with the following command: ```cd CRM_Backend_Project``` and running the command: ```"go run main.go"```. This will start the server up on localhost:3000, this project also contains two additional modules: "gorilla mux" and "google uuid" for routing and ensuring that each customer ID is unique. After the app is running, it will be possible to send requests to different endpoints in order to obtain the data neccessary, the static HTML that is present on the endpoint "/" will also provide more details into the data that can be obtained from each of the endpoints.