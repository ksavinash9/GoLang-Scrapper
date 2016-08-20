A Simple Amazon Scrapper
===========================================

## Tasks
- [x] realtime scrapping
- [x] unit test for APIs

## Usage

### Installation
-----------
```sh
$ go install
```

### Running Server
-----------
```sh
$ go build
$ ./binary
```

### Testing
-----------
```sh
$ go test test/
```


### EndPoints
-------------
#### Valid ID
GET - http://localhost:8080/movie/amazon/B00K19SD8Q
```json
{
   "title":"Um Jeden Preis",
   "poster":"https://images-eu.ssl-images-amazon.com/images/I/51VELYHd4TL._SX200_QL80_.jpg",
   "release_year":2013,
   "actors":[
      "Dennis Quaid",
      " Zac Efron",
      " Kim Dickens"
   ],
   "similar_ids":[
      "B00HDZMP94",
      "B00JM0JXYI",
      "B00HUZXE8S",
      "B00L9KET84",
      "B00O4G7QQ2",
      "B00ILR6N8M",
      "B00ILNAY5O",
      "B00I4DWI6O",
      "B00HXL8THU",
      "B00KR5NIZC",
      "B00IKEQKU2",
      "B00ILNV05W",
      "B00ILIGIU4",
      "B00G0NPW1I",
      "B00IK6HLUI",
      "B00JLCZN8M",
      "B00OGSSXLK",
      "B00IU74CQS",
      "B00IYSFS6Q",
      "B00PWXF8JI"
   ]
}
```
#### Invalid ID
GET - http://localhost:8080/movie/amazon/ABCDEFGHIJ
```json
{
   "code":404,
   "text":"Item Not Found"
}
```
#### Null Product ID
GET - http://localhost:8080/movie/amazon/
```json
{
   "code":200,
   "text":"Enter Amazon Product Id"
}
```
#### Network Issue
GET - http://localhost:8080/movie/amazon/B00K19SD8Q
```json
{
   "code":500,
   "text":"Site Unreachable"
}
```
