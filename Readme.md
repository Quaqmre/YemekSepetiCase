### YemekSepetiCase

This case have one store package and one api package to write persitance in-memory api to store data,

3 API available to access, you can find postman collection thoose;

- /get
- /create 
- /flush

Curl request like:

__/get__

curl --location --request GET 'localhost:8080/get/1'

* You can get value with given key in the URL path , but if it doesnt exist 404 will return

__/create__

curl --location --request POST 'localhost:8080/create' \
--header 'Content-Type: text/plain' \
--data-raw 'denemetest3'

* You can put the value with default key , and api return object key value.Then u can access value with given key

__/flush__

curl --location --request PUT 'localhost:8080/flush' 

* Flush is clear the hole data in the storage



###### By the way every one minute in-memory db write themself into the file for persistance mod.


## Build and Run Stage

1- docker build -t yemekSepetiCase .
2- docker run -p 8080:8080 yemekSepetiCase 

Then u can access api with postman collenction or curl request easly.



