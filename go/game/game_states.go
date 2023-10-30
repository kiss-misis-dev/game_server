package game

type State int8

const (
	StateUndefined State = iota
	StateScheduled
	StateRegistrationOpen
	StateRegistrationClosed
	StatePlay
	StatePause
	StateRound
	StateVerify
	StateUnconfirmed
	StateClosed
)

type Event int8

const (
	EventUndefined Event = iota
	EventOpenRegistration
	EventCloseRegistration
	EventBeginGame
	EventEndGame
	EventPauseGame
	EventResumeGame
	EventBeginRound
	EventEndRound
	EventStartVerification
	EventStopVerification
	EventConfirm
)

type InvalidEvent struct {
	CurrentState  State
	OccurredEvent Event
}

func (error InvalidEvent) Error() string {
	return "invalid game event"
}

func (currentState *State) HandleEvent(event Event) (err error) {
	isValid := switchEvent(currentState, event)
	if !isValid {
		err = InvalidEvent{*currentState, event}
	}
	return
}

func switchEvent(currentState *State, event Event) (isValid bool) {
	switch event {
	case EventUndefined:
		isValid = false
	case EventOpenRegistration:
		isValid = handleEventOpenRegistration(currentState)
	case EventCloseRegistration:
		isValid = handleEventCloseRegistration(currentState)
	case EventBeginGame:
		isValid = handleEventBeginGame(currentState)
	case EventEndGame:
		isValid = handleEventEndGame(currentState)
	case EventPauseGame:
		isValid = handleEventPauseGame(currentState)
	case EventResumeGame:
		isValid = handleEventResumeGame(currentState)
	case EventBeginRound:
		isValid = handleEventBeginRound(currentState)
	case EventEndRound:
		isValid = handleEventEndRound(currentState)
	case EventStartVerification:
		isValid = handleEventStartVerification(currentState)
	case EventStopVerification:
		isValid = handleEventStopVerification(currentState)
	case EventConfirm:
		isValid = handleEventConfirm(currentState)
	}
	return
}

func handleEventOpenRegistration(currentState *State) (isValid bool) {
	if *currentState == StateScheduled {
		*currentState = StateRegistrationOpen
		isValid = true
	} else {
		isValid = false
	}
	return
}

func handleEventCloseRegistration(currentState *State) (isValid bool) {
	if *currentState == StateRegistrationOpen {
		*currentState = StateRegistrationClosed
		isValid = true
	} else {
		isValid = false
	}
	return
}

func handleEventBeginGame(currentState *State) (isValid bool) {
	if *currentState == StateRegistrationClosed {
		*currentState = StatePlay
		isValid = true
	} else {
		isValid = false
	}
	return
}

func handleEventEndGame(currentState *State) (isValid bool) {
	if *currentState == StatePlay {
		*currentState = StateUnconfirmed
		isValid = true
	} else {
		isValid = false
	}
	return
}

func handleEventPauseGame(currentState *State) (isValid bool) {
	if *currentState == StatePlay {
		*currentState = StatePause
		isValid = true
	} else {
		isValid = false
	}
	return
}

func handleEventResumeGame(currentState *State) (isValid bool) {
	if *currentState == StatePause {
		*currentState = StatePlay
		isValid = true
	} else {
		isValid = false
	}
	return
}

func handleEventBeginRound(currentState *State) (isValid bool) {
	if *currentState == StatePlay {
		*currentState = StateRound
		isValid = true
	} else {
		isValid = false
	}
	return
}

func handleEventEndRound(currentState *State) (isValid bool) {
	if *currentState == StateRound {
		*currentState = StatePlay
		isValid = true
	} else {
		isValid = false
	}
	return
}

func handleEventStartVerification(currentState *State) (isValid bool) {
	if *currentState == StatePlay {
		*currentState = StateVerify
		isValid = true
	} else {
		isValid = false
	}
	return
}

func handleEventStopVerification(currentState *State) (isValid bool) {
	if *currentState == StateVerify {
		*currentState = StatePlay
		isValid = true
	} else {
		isValid = false
	}
	return
}

func handleEventConfirm(currentState *State) (isValid bool) {
	if *currentState == StateUnconfirmed {
		*currentState = StateClosed
		isValid = true
	} else {
		isValid = false
	}
	return
}
