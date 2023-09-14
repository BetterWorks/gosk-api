package models

import (
	"time"

	"github.com/BetterWorks/gosk-api/internal/core/jsonapi"
	"github.com/BetterWorks/gosk-api/internal/core/pagination"
	"github.com/google/uuid"
)

// ExampleInputData defines the subset of Example domain model attributes that are accepted
// for input data request binding
type ExampleInputData struct {
	Deleted     bool    `json:"deleted" validate:"omitempty,boolean"`
	Description *string `json:"description" validate:"omitempty,min=3,max=999"`
	Enabled     bool    `json:"enabled"  validate:"omitempty,boolean"`
	Status      *uint32 `json:"status" validate:"omitempty,numeric"`
	Title       string  `json:"title" validate:"required,omitempty,min=2,max=255"`
}

// ExampleDomainModel an Example domain model that contains one or more ExampleObject(s)
// and related metadata
type ExampleDomainModel struct {
	Data []ExampleObject
	Meta *ModelMetadata
	Solo bool
}

type ModelMetadata struct {
	Paging pagination.PageMetadata
}

// ExampleObject
type ExampleObject struct {
	Attributes ExampleObjectAttributes
	Meta       any
	Related    any
}

// Example defines an Example domain model for application logic
type ExampleObjectAttributes struct {
	ID          uuid.UUID  `json:"-"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Status      *uint32    `json:"status"`
	Enabled     bool       `json:"enabled"`
	Deleted     bool       `json:"-"`
	CreatedOn   time.Time  `json:"created_on"`
	CreatedBy   uint32     `json:"created_by"`
	ModifiedOn  *time.Time `json:"modified_on"`
	ModifiedBy  *uint32    `json:"modified_by"`
}

func (m *ExampleDomainModel) FormatResponse() (*jsonapi.Response, error) {
	if m.Solo {
		resource := formatResource(&m.Data[0])
		response := &jsonapi.Response{Data: resource}
		return response, nil
	}

	meta := &jsonapi.ResponseMetadata{
		Paging: pagination.PageMetadata{
			Limit:  m.Meta.Paging.Limit,
			Offset: m.Meta.Paging.Offset,
			Total:  m.Meta.Paging.Total,
		},
	}

	data := make([]jsonapi.ResponseResource, 0, len(m.Data))
	for _, domo := range m.Data {
		resource := formatResource(&domo)
		data = append(data, resource)
	}
	response := &jsonapi.Response{
		Meta: meta,
		Data: data,
	}

	return response, nil
}

// serializeResource
func formatResource(domo *ExampleObject) jsonapi.ResponseResource {
	return jsonapi.ResponseResource{
		Type: "example", // TODO
		ID:   domo.Attributes.ID,
		// Meta: domo.Meta,
		Attributes: ExampleObjectAttributes{
			Title:       domo.Attributes.Title,
			Description: domo.Attributes.Description,
			Status:      domo.Attributes.Status,
			Enabled:     domo.Attributes.Enabled,
			Deleted:     domo.Attributes.Deleted,
			CreatedOn:   domo.Attributes.CreatedOn,
			CreatedBy:   domo.Attributes.CreatedBy,
			ModifiedOn:  domo.Attributes.ModifiedOn,
			ModifiedBy:  domo.Attributes.ModifiedBy,
		},
	}
}
