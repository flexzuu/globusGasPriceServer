`env GOOS=linux GOARCH=arm GOARM=5 go build`
`eval (docker-machine env black-pearl)`
`docker-compose up --build`

TODO:
- [x] Fix DB and go type missmatches. 
- [ ] Build filler go tool
- [ ] Setup dockerfile with cron (https://www.ekito.fr/people/run-a-cron-job-with-docker/) & filler
- [ ] Add dockerfile to compose file
