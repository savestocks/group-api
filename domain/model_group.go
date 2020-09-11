package domain

import (
    "time"
)

//Group represents a model
type Group struct {
    ID string `json:"id"`
    CreatedAt time.Time `json:"dateCreated,omitempty"`
    UpdatedAt time.Time `json:"dateUpdated,omitempty"`
    Name string `json:"name"`
}