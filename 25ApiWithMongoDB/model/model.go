package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Netflix struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bosn:"_id,omitempty"`
	Movie       string             `json:"movie,omitempty"`
	Banner      string             `json:"banner"`
	Casts       *Casts             `json:"casts"`
	ReleaseDate string             `json:"releasedate"`
	Watched     bool               `json:"watched,omitempty"`
}

type Casts struct {
	Directors []string `json:"directors"`
	Producers []string `json:"producers"`
	Platforms []string `json:"platforms"`
	Actors    []string `json:"actors"`
}
