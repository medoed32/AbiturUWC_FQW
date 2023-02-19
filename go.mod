module AbiturUWC_FQW

go 1.20

require (
	github.com/julienschmidt/httprouter v1.3.0
	internal/app/authorizationApp v0.0.0
)

replace internal/app/authorizationApp => ./internal/app/authorizationApp
