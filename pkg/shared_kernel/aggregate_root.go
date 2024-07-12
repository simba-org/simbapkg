package sharedkernel

type AggregateRoot struct {
	domainEvents []DomainEvent
}

func (ar *AggregateRoot) ApplyDomain(e DomainEvent) {
	ar.domainEvents = append(ar.domainEvents, e)
}

func (ar *AggregateRoot) DomainEvents() []DomainEvent {
	return ar.domainEvents
}

func (ar *AggregateRoot) RemoveDomainEvents(eventIdentity string) {
	for i := 0; i < len(ar.domainEvents); i++ {
		if ar.domainEvents[i].Identity() == eventIdentity {
			ar.domainEvents = append(ar.domainEvents[:i], ar.domainEvents[i+1:]...)
			i--
		}
	}
}
