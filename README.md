`env GOOS=linux GOARCH=arm GOARM=5 go build`
`eval (docker-machine env black-pearl)`
`docker-compose up --build`

DONE: Fix DB and go type missmatches.