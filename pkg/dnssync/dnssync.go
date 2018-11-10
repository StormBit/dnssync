package dnssync

import (
	"log"

	"github.com/cloudflare/cloudflare-go"

	"github.com/stormbit/dnssync/pkg/diff"
)

// DNSSync contains state for the sync operation
type DNSSync struct { 
	api 	*cloudflare.API
	dryMode bool
	zoneID 	string
}

// ApplyZoneState applies state for any records within the zone
func (c *DNSSync) ApplyZoneState(desiredState ZoneState) {
	for _, recordState := range desiredState.Records {
		c.ApplyRecordState(recordState)
	}
}

// ApplyRecordState applies a desired Record's state
func (c *DNSSync) ApplyRecordState(desiredState RecordState) {

	log.Println("Evaluating:", desiredState.Name, "type", desiredState.Type)

	// Get current state.
	var currentState []string
	records, _ := c.api.DNSRecords(c.zoneID, cloudflare.DNSRecord{
		Name: desiredState.Name,
		Type: desiredState.Type,
	})
	for _, record := range records {
		currentState = append(currentState, record.Content)
	}

	// Diff the state.
	toCreate, toDelete := diff.DiffState(currentState, desiredState.Value)

	if (len(toCreate) == 0 && len(toDelete) == 0) {
		log.Println("Nothing to do!")
		return
	}

	// Create new records.
	for _, record := range toCreate {
		log.Println("Creating:", desiredState.Name, "with value", record)

		if (c.dryMode) {
			break
		}

		_, err := c.api.CreateDNSRecord(c.zoneID, cloudflare.DNSRecord{
			Name: desiredState.Name,
			Type: desiredState.Type,
			Content: record,
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	// Delete old ones.
	for _, record := range toDelete {
		log.Println("Deleting:", desiredState.Name, "with value", record)

		if (c.dryMode) {
			break
		}

		records, err := c.api.DNSRecords(c.zoneID, cloudflare.DNSRecord{
			Name: desiredState.Name,
			Type: desiredState.Type,
			Content: record,
		})
		if err != nil {
			log.Fatal(err)
		}


		for _, recordToDelete := range records {
			err := c.api.DeleteDNSRecord(c.zoneID, recordToDelete.ID)
			if err != nil {
				log.Fatal(err)
			}
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}

// NewCloudflareDNSSync returns a DNSSync instance
func NewCloudflareDNSSync(zone string, apiKey string, apiEmail string) *DNSSync {
	// Construct a new API object
	api, err := cloudflare.New(apiKey, apiEmail)
	if err != nil {
		log.Fatal(err)
	}

	id, err := api.ZoneIDByName(zone)
	if err != nil {
		log.Fatal(err)
	}

	return &DNSSync{
		zoneID: id,
		api: api,
	}
}
