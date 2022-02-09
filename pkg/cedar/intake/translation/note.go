package translation

import (
	"encoding/json"

	wire "github.com/cmsgov/easi-app/pkg/cedar/intake/gen/models"
	intakemodels "github.com/cmsgov/easi-app/pkg/cedar/intake/models"
	"github.com/cmsgov/easi-app/pkg/models"
)

// TranslatableNote is a wrapper around our Note model for translating into the CEDAR Intake API schema
type TranslatableNote models.Note

// ObjectID is a unique identifier for a TranslatableNote
func (note *TranslatableNote) ObjectID() string {
	return note.ID.String()
}

// ObjectType is a human-readable identifier for the Note type, for use in logging
func (note *TranslatableNote) ObjectType() string {
	return "note"
}

// CreateIntakeModel translates a Note into an IntakeInput
func (note *TranslatableNote) CreateIntakeModel() (*wire.IntakeInput, error) {
	obj := intakemodels.EASINote{
		IntakeID:  note.SystemIntakeID.String(),
		AuthorEUA: note.AuthorEUAID,
		Content:   note.Content.ValueOrZero(),
	}

	blob, err := json.Marshal(&obj)
	if err != nil {
		return nil, err
	}

	result := wire.IntakeInput{
		ID:   pStr(note.ID.String()),
		Body: pStr(string(blob)),

		// invariants for this type
		Status:     pStr(wire.IntakeInputStatusFinal),
		BodyFormat: pStr(wire.IntakeInputBodyFormatJSON),
		Type:       pStr(wire.IntakeInputTypeEASINote),
		Schema:     versionStr(IntakeInputSchemaEASINoteV01),
	}

	if note.CreatedAt != nil {
		result.CreatedDate = pDateTime(note.CreatedAt)
		result.LastUpdate = pDateTime(note.CreatedAt)
	}

	return &result, nil
}