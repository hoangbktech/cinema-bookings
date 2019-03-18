#!/bin/bash

#Movie
protoc --proto_path=movie-service/api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:movie-service/pkg/api/v1 movie-service.proto

#Cinema
protoc --proto_path=cinema-service/api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:cinema-service/pkg/api/v1 cinema-service.proto

#User
protoc --proto_path=user-service/api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:user-service/pkg/api/v1 user-service.proto

#Booking
protoc --proto_path=booking-service/api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:booking-service/pkg/api/v1 booking-service.proto
protoc --proto_path=booking-service/api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:booking-service/pkg/api/v1 booking-service.proto
protoc --proto_path=booking-service/api/proto/v1 --proto_path=third_party --swagger_out=logtostderr=true:booking-service/api/swagger/v1 booking-service.proto

#Notification
protoc --proto_path=notification-service/api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:notification-service/pkg/api/v1 notification-service.proto

#Document Service
protoc --proto_path=document-service/api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:document-service/pkg/api/v1 document-service.proto

