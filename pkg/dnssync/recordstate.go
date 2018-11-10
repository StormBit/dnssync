package dnssync

// RecordState specifies the state of a DNS Record
type RecordState struct {
    Name string     `json:"name"`
    Type string     `json:"type"`
    Value []string  `json:"value"`
}