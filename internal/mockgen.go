package internal

import _ "github.com/golang/mock/mockgen/model"

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks consumer-producer/internal/app/repo EventRepo
//go:generate mockgen -destination=./mocks/send_mock.go -package=mocks consumer-producer/internal/app/sender EventSender
