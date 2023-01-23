package services

import (
	"github.com/sirupsen/logrus"
)

type Metrics struct {
	//
}

func NewMetrics() *Metrics {
	return &Metrics{}
}

func (m Metrics) Decorate(function IDecoratedFunction) IDecoratedFunction {
	//
	logrus.Info("Metric stored in the prometheus")
	return function
}
