package model

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
)

// RavoRecord defines a Ravo XML record.
type RavoRecord struct {
	Klant             bool `xml:"klant"`
	Sales             bool `xml:"Sales"`
	AfterSalesService bool `xml:"AfterSalesService"`
	Productie         bool `xml:"Productie"`
	Kwaliteit         bool `xml:"Kwaliteit"`
	Inkoop            bool `xml:"Inkoop"`
	Engineering       bool `xml:"Engineering"`
}

// ReadRavoRecord read a Ravo XML record.
func ReadRavoRecord(r io.Reader) *RavoRecord {
	var ravoRecord RavoRecord
	decoder := xml.NewDecoder(r)
	err := decoder.Decode(&ravoRecord)
	if err != nil {
		log.Fatalf("error unmarshaling Ravo record: %v", err)
	}
	return &ravoRecord
}

// WriteJSON writes a Ravo record as JSON
func (rr *RavoRecord) WriteJSON(w io.Writer) {
	data, err := json.MarshalIndent(rr, "", "    ")
	if err != nil {
		log.Fatalf("failed to marshal to JSON: %v", err)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Fatalf("failed to write JSON: %v", err)
	}
}
