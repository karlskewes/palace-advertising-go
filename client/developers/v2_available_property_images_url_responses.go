// Code generated by go-swagger; DO NOT EDIT.

package developers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/karlskewes/palace-advertising-go/models"
)

// V2AvailablePropertyImagesURLReader is a Reader for the V2AvailablePropertyImagesURL structure.
type V2AvailablePropertyImagesURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2AvailablePropertyImagesURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2AvailablePropertyImagesURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewV2AvailablePropertyImagesURLOK creates a V2AvailablePropertyImagesURLOK with default headers values
func NewV2AvailablePropertyImagesURLOK() *V2AvailablePropertyImagesURLOK {
	return &V2AvailablePropertyImagesURLOK{}
}

/*V2AvailablePropertyImagesURLOK handles this case with default header values.

Images (URL Links) and meta data related to 'Available Property'
*/
type V2AvailablePropertyImagesURLOK struct {
	Payload []*models.V2AvailablePropertyImagesURL
}

func (o *V2AvailablePropertyImagesURLOK) Error() string {
	return fmt.Sprintf("[GET /v2AvailablePropertyImagesURL/JSON/{PropertyCode}][%d] v2AvailablePropertyImagesUrlOK  %+v", 200, o.Payload)
}

func (o *V2AvailablePropertyImagesURLOK) GetPayload() []*models.V2AvailablePropertyImagesURL {
	return o.Payload
}

func (o *V2AvailablePropertyImagesURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
