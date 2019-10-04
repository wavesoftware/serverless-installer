package model

// Install - should install given tool
type Install struct {
	Enabled bool
}

// Jeager - a struct that represents a Jeager installation options
type Jeager struct {
	Install
}

// Elasticsearch - a struct that represents a Elastic Search installation options
type Elasticsearch struct {
	Install
}

// Kiali - a struct that represents a Kiali installation options
type Kiali struct {
	Install
}

// Istio - a struct that represents a Istio installation options
type Istio struct {
	Install
}

// Serverless - a struct that represents a Serverless installation options
type Serverless struct {
	Install
}

// Answers - a struct that represents answers
type Answers struct {
	Jeager
	Elasticsearch
	Kiali
	Istio
	Serverless
}
