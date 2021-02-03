# Gothon

So, this project is written using two languages -- python and golang(python version may be missing). Golang part of the project is already finished.

#How to run golang part?

First, you need to edit vars.env file. Here is a small template

```bash
  POSTGRES_USER=<postgres_user>
  POSTGRES_PASSWORD=<postgres_pass>
  POSTGRES_DB=<postgres_db>
  ADDR=0.0.0.0:<desired_port>
```

Second, you need to run docker-compose up. Is that it? Patience, there's one thing left to do. 

#Running query to create table in database
By the way, you also need to create database and name it gothon beforehand. The next command to run will be the following:
```bash
  sudo docker container exec <container_id> export PGPASSWORD=<postgres_password"; psql -h <your_lan_ip> -U postgres -W gothon -c "$(cat init_db.sql)"
```
Don't worry if an error unexpectedly appears out of nowhere -- that way the terminal will be blocked, so to enable you to type in postgres password
In a word, that particular command tells docker to log into the psql environment and immediately execute contents of init_db.sql, which happens to be 
a postgresql query to initialize new table. 

#Architecture

Api has only 3 requests, namely saveStatistics, retrieveStatistics, and deleteAllStatistics. 
Here are few examples of making requests to the server using http.Client structure.

```go
	client := &http.Client{}

	url := "http://0.0.0.0:7890/save_stats/"
	params := map[string]string{
		"date":   "2017-11-30",
		"views":  "120",
		"clicks": "240",
		"cost":   "34r 20k",
	}

	err, res := makePostMultipartRequest(client, url, params)
}
```
As you can see, request accepts 4 parameters -- date, views, clicks, and cost, with 3 of them being optional, namely views, clicks, and cost.
First, server checks against presence of mandatory date field and immediately spits out bad response in case if condition takes place. 
Afterwards, it's analysing format structure of input date. If the latter does not conform to 'YYYY-mm-dd', server responds with "BadRequest" message.
Here is code for makePostMultipartRequest function:
```go
  func makePostMultipartRequest(client *http.Client, url string, params map[string]string) (err error, res *http.Response) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	for key, value := range params {
		err := w.WriteField(key, value)
		if err != nil {
			log.Println(err)
		}
	}
	w.Close()

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return err, nil
	}

	req.Header.Set("Content-Type", w.FormDataContentType())
	res, err = client.Do(req)
	if err != nil {
		return err, nil
	}

	return nil, res
}
```
This function is responsible for filling out request body and making actual request to the server.
The next snippet greatly illustrates an example to make the request for retrieving statistics saved ealier 
```go
	from := "2017-11-30"
	to := "2077-12-25"

	url := "http://0.0.0.0:7890/retrieve_stats/?to="+to+"&from="+from
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil); if err != nil{t.Error(err)}
	res, err := client.Do(req); if err != nil{t.Error(err)}
```
All needed information is put into an url string, by the virtue of nature of GET requests. 

Lastly, I need to demostrate an example of deleting all statistics from database. That's the easiest one.

```go
	url := "http://0.0.0.0:7890/delete_stats/"

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", url, nil); if err != nil{t.Error(err)}
	res, err := client.Do(req); if err != nil{t.Error(err)}
```
As you might have already noticed, this request does not require any parameters.

#A bit of plans

This project was conceived with dual language implementation in mind. So, python version's coming soon ;)
