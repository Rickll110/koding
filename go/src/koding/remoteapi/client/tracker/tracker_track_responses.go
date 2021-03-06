package tracker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// TrackerTrackReader is a Reader for the TrackerTrack structure.
type TrackerTrackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TrackerTrackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTrackerTrackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewTrackerTrackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTrackerTrackOK creates a TrackerTrackOK with default headers values
func NewTrackerTrackOK() *TrackerTrackOK {
	return &TrackerTrackOK{}
}

/*TrackerTrackOK handles this case with default header values.

Request processed successfully
*/
type TrackerTrackOK struct {
	Payload *models.DefaultResponse
}

func (o *TrackerTrackOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/Tracker.track][%d] trackerTrackOK  %+v", 200, o.Payload)
}

func (o *TrackerTrackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTrackerTrackUnauthorized creates a TrackerTrackUnauthorized with default headers values
func NewTrackerTrackUnauthorized() *TrackerTrackUnauthorized {
	return &TrackerTrackUnauthorized{}
}

/*TrackerTrackUnauthorized handles this case with default header values.

Unauthorized request
*/
type TrackerTrackUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *TrackerTrackUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/Tracker.track][%d] trackerTrackUnauthorized  %+v", 401, o.Payload)
}

func (o *TrackerTrackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
