# string-encryptor
This repository consisit of 2 microservices: string-enctyptor and string-randomizer.<br/>
String encryptor is grpc server. It can be started by **go run main.go** command, number of workers for workerpool can be defined by **WORKERS** environmental variable (example WORKERS=10 go run main.go). Encryptor microservice works on port 8080.<br/>
String randomizer is http server which creates a number of wanted random strings. It works on port 8081. In order to define number of strings http post request should be done. JSON file should be in format:<br/>
<pre>
{
  "n":"10"
}
 </pre>
 , where n is the number of wanted strings. Response will be array of encrypted strings in JSON format. Server can be started by command **go run main.go**.
