export $(grep -v '^#' .env | xargs) && air
