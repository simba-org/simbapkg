package sharedkernel

type AggregateRoot struct {
	domainEvents []DomainEvent
}

// ApplyDomain
//
//	@Description: 绑定事件，需要去重
//	@receiver ar
//	@param e
func (ar *AggregateRoot) ApplyDomain(e DomainEvent) {
	// 直接操作原始切片 ar.domainEvents
	for i, v := range ar.domainEvents {
		if v.Identity() == e.Identity() {
			ar.domainEvents[i] = e
			return
		}
	}
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
