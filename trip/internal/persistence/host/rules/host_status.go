package rules

import (
	"fmt"
	hostModel "trip/internal/persistence/host/model"
)

type StatusTransition struct {
	From hostModel.Status
	To   hostModel.Status
}

var AllowedStatusTransitions = map[StatusTransition]struct{}{
	{hostModel.StatusDraft, hostModel.StatusPending}:       {},
	{hostModel.StatusPending, hostModel.StatusRejected}:    {},
	{hostModel.StatusPending, hostModel.StatusApproved}:    {},
	{hostModel.StatusRejected, hostModel.StatusPending}:    {},
	{hostModel.StatusRejected, hostModel.StatusApproved}:   {},
	{hostModel.StatusApproved, hostModel.StatusSuspension}: {},
	{hostModel.StatusApproved, hostModel.StatusDeleted}:    {},
	{hostModel.StatusSuspension, hostModel.StatusApproved}: {},
	{hostModel.StatusSuspension, hostModel.StatusDeleted}:  {},
}

var AcceptableStatusTransitions = map[StatusTransition]struct{}{
	{hostModel.StatusPending, hostModel.StatusApproved}:  {},
	{hostModel.StatusRejected, hostModel.StatusApproved}: {},
}

func IsStatusTransitionAllowed(from, to hostModel.Status) bool {
	_, ok := AllowedStatusTransitions[StatusTransition{from, to}]
	return ok
}

func IsStatusTransitionAcceptable(from, to hostModel.Status) bool {
	_, ok := AcceptableStatusTransitions[StatusTransition{from, to}]
	return ok
}

func ValidateStatusTransition(from, to hostModel.Status) error {
	if from == to {
		return nil
	}

	if !IsStatusTransitionAllowed(from, to) {
		return fmt.Errorf("invalid status transition: %d → %d", from, to)
	}

	return nil
}
