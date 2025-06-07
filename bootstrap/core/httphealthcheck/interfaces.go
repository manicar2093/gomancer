package httphealthcheck

type Checkable interface {
	ServiceHealth() (HealthStatusData, error)
	ServiceName() string
}
