package server

import (
	"github.com/hudyweas/panshee/services/wallet_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/wallet_service/config"
	bscscanapi "github.com/hudyweas/panshee/services/wallet_service/external/bscscan"
	"github.com/sirupsen/logrus"
)

type Server struct {
	pb.UnimplementedPansheeWalletServiceServer
	config config.Config
	log *logrus.Logger

	bsc bscscanapi.BscAPI
}

func NewServer(config config.Config, log *logrus.Logger) (s *Server) {
	s = &Server{
		config: config,
		log: log,
	}

	s.bsc = bscscanapi.CreateBscScanAPI(config.BSCSCAN_KEY);

	s.log.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	return
}
