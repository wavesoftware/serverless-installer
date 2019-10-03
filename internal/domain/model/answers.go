package model

// Jeager - a struct that represents a Jeager installation options
type Jeager struct {
	Install bool
}

// Elasticsearch - a struct that represents a Elastic Search installation options
type Elasticsearch struct {
	Install bool
}

// Kiali - a struct that represents a Kiali installation options
type Kiali struct {
	Install bool
}

// Istio - a struct that represents a Istio installation options
type Istio struct {
	Install bool
}

// Serverless - a struct that represents a Serverless installation options
type Serverless struct {
	Install bool
}

// Answers - a struct that represents answers
type Answers struct {
	Jeager        Jeager
	Elasticsearch Elasticsearch
	Kiali         Kiali
	Istio         Istio
	Serverless    Serverless
}
