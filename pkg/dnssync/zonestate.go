package dnssync

// ZoneState is a collection of RecordStates
type ZoneState struct {
	Name string 			`json:"name"`
    Records []RecordState	`json:"records"`
}