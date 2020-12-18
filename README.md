## How to run the API

* import the packages by running go get -d
* If above step is successful go to `go-api` folder and run
     ```
     go build -ldflags="-X 'main.Version=v0.0.1' -X 'main.MinVersion=`date -u +.%Y%m%d.%H%M%S`' -X 'main.BuildTime=$(date)'"
  
  If no errors go-api.exe will be created on windows machine.
  
* It shouldn't throw any error if it does go back to step 1 and do all the steps again.
* Run it to a given port (ex:8080) by using:
     ```bigquery
     go-api --port=8080
  
* To test if server is up do this url
```bigquery
    http://localhost:8080/