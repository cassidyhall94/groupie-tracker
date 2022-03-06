Deployed at: https://groupie-tracker-01.herokuapp.com/

## Building docker for running locally

- Run `/build.sh`
- It'll tell you what to do

## Building and running on Heroku

If you haven't logged in then do this first:
- `heroku container:login`

To build and release:

- `heroku container:push web -a groupie-tracker-01`
- `heroku container:release web -a groupie-tracker-01`