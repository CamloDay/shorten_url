i have not included the compiled files, so to build them run
go build main.go
go build cli/shorten/shorten_cli.go
go build cli/redirect/redirect_cli.go

to start the service
./main --ADDRESS='localhost:8080'
ADDRESS defaults to 'localhost:8080'
this will block the terminal by listening

shorten ORIGINAL url using
./shorten_cli --ADDRESS='localhost:8080' --ORIGINAL='http://abc.com'
ADDRESS defaults to 'localhost:8080'
ORIGINAL is the url that you want to shorten
response will be in a json body with the shortened url under a 'short' tag
example response {"status_code":200,"description":"Ok","short":"http://localhost:8080/adc6adf6-e99b-4fb4-8867-4d1a1996f6ec"}

redirect to the original url
./redirect_cli --SHORT='http://localhost:8080/e7e44623-afb7-4b92-9c4c-9f673c243ad8'
pass the SHORT in from the shorten response
