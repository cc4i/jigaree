package service

import (
	"context"
	pb "weather/protos"
)

type WeatheService struct {
	pb.UnimplementedWeatherServiceServer
}

func (s *WeatheService) WeatherbyCity(ctx context.Context, in *pb.CityWeatherRequest) (*pb.CityWeatherResponse, error) {

	return &pb.CityWeatherResponse{}, nil
}
