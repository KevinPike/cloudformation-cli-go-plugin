package errcode

import (
	"strings"
)

// Status ...
type Status int

func (c Status) String() string {
	if c < NotUpdatable || c > InternalFailure {
		return "Unknown"
	}

	statuses := []string{
		"Unknown",
		"NotUpdatable",
		"InvalidRequest",
		"AccessDenied",
		"InvalidCredentials",
		"AlreadyExists",
		"NotFound",
		"ResourceConflict",
		"Throttling",
		"ServiceLimitExceeded",
		"NotStabilized",
		"GeneralServiceException",
		"ServiceInternalError",
		"NetworkFailure",
		"InternalFailure",
	}

	return statuses[c]
}

//Convert string to Status
func Convert(s string) Status {
	status := strings.ToUpper(s)

	statuses := []string{
		"Unknown",
		"NotUpdatable",
		"InvalidRequest",
		"AccessDenied",
		"InvalidCredentials",
		"AlreadyExists",
		"NotFound",
		"ResourceConflict",
		"Throttling",
		"ServiceLimitExceeded",
		"NotStabilized",
		"GeneralServiceException",
		"ServiceInternalError",
		"NetworkFailure",
		"InternalFailure",
	}

	for i, v := range statuses {
		if status == strings.ToUpper(v) {
			return Status(i)
		}
	}

	return Unknown
}

const (
	// Unknown ...
	Unknown Status = iota

	// NotUpdatable is when the customer tried perform an update to a property that is CreateOnly. Only
	// applicable to Update Handler. (Terminal)
	NotUpdatable

	// InvalidRequest is a generic exception caused by invalid input from the customer. (Terminal)
	InvalidRequest

	//AccessDenied is when the customer has insufficient permissions to perform this action. (Terminal)
	AccessDenied

	//InvalidCredentials is when the customer's provided credentials were invalid. (Terminal)
	InvalidCredentials

	//AlreadyExists is when the specified resource already existed prior to the execution of the handler.
	//Only applicable to Create Handler (Terminal) Handlers MUST return this error
	//when duplicate creation requests are received.
	AlreadyExists

	//NotFound is when the specified resource does not exist, or is in a terminal, inoperable, and
	//irrecoverable state. (Terminal)
	NotFound

	//ResourceConflict is when the resource is temporarily unable to be acted upon; for example, if the
	//resource is currently undergoing an operation and cannot be acted upon until
	//that operation is finished (Retriable)
	ResourceConflict

	//Throttling is when the request was throttled by the downstream service. (Retriable)
	Throttling

	//ServiceLimitExceeded is when a non-transient resource limit was reached on the service side. (Terminal)
	ServiceLimitExceeded

	//NotStabilized is when the downstream resource failed to complete all of its ready state checks.
	//(Retriable)
	NotStabilized

	//GeneralServiceException is an exception from the downstream service that does not map to any other error
	//codes. (Terminal)
	GeneralServiceException

	//ServiceInternalError is when the downstream service returned an internal error, typically with a 5XX HTTP
	//Status code. (Retriable)
	ServiceInternalError

	//NetworkFailure is when the request was unable to be completed due to networking issues, such as
	//failure to receive a response from the server. (Retriable)
	NetworkFailure

	// InternalFailure is an unexpected error occurred within the handler, such as an NPE, etc.
	//(Terminal)
	InternalFailure
)
