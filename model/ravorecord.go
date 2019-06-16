package model

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
)

// RavoBool defines Ravo boolean type as encoded in XML
type RavoBool struct {
	Nil   bool `xml:"nil,attr" json:"nil"`
	Value bool `xml:",innerxml" json:"value"`
}

type RavoHTML struct {
	InnerXML string `xml:",innerxml"`
}

// RavoRecord defines a Ravo XML record.
type RavoRecord struct {
	Klant                                     bool   `xml:"klant"`
	Sales                                     bool   `xml:"Sales"`
	AfterSalesService                         bool   `xml:"AfterSalesService"`
	Productie                                 bool   `xml:"Productie"`
	Kwaliteit                                 bool   `xml:"Kwaliteit"`
	Inkoop                                    bool   `xml:"Inkoop"`
	Engineering                               bool   `xml:"Engineering"`
	Overige                                   bool   `xml:"Overige"`
	Overige2                                  bool   `xml:"Overige2"`
	IngediendDoor                             string `xml:"IngediendDoor"`
	RWVnummer                                 int64  `xml:"RWVnummer"`
	DatumIngediend                            string `xml:"DatumIngediend"`
	VeiligheidsRisico                         bool   `xml:"VeiligheidsRisico"`
	MilieuSchadeRisico                        bool   `xml:"MilieuSchadeRisico"`
	Klantklacht                               bool   `xml:"Klantklacht"`
	WettelijkVereist                          bool   `xml:"WettelijkVereist"`
	GarantieAanspraak                         bool   `xml:"GarantieAanspraak"`
	NietConformFunctieEisen                   bool   `xml:"NietConformFunctieEisen"`
	NietConformKwaliteitEisen                 bool   `xml:"NietConformKwaliteitEisen"`
	EfficiencyVeranderingMontageMinuten       bool   `xml:"EfficiencyVeranderingMontageMinuten"`
	KostprijsVerandering                      bool   `xml:"KostprijsVerandering"`
	ProductStandaardisatie                    bool   `xml:"ProductStandaardisatie"`
	ProductwijzigingLeverancier               bool   `xml:"ProductwijzigingLeverancier"`
	ProductwijzigingIntern                    bool   `xml:"ProductwijzigingIntern"`
	Anders                                    bool   `xml:"Anders"`
	VeiligheidsrisicoAantal                   RavoBool
	MilieuschadeRisicoAantal                  RavoBool
	KlantklachtAantal                         RavoBool
	WettelijkVereistAantal                    RavoBool
	GarantieAanspraakAantal                   RavoBool
	NietConformFunctieEisenAantal             RavoBool
	NietConformKwaliteitEisenAantal           int64 `xml:"NietConformKwaliteitEisenAantal"`
	EfficiencyVeranderingMontageMinutenAantal RavoBool
	KostprijsVeranderingAantal                int64 `xml:"KostprijsVeranderingAantal"`
	ProductStandaardisatieAantal              RavoBool
	ProductwijzigingLeverancierAantal         RavoBool
	ProductwijzigingInternAantal              RavoBool
	AndersAantal                              RavoBool
	AndersWat                                 string `xml:"AndersWat"`
	Probleemstelling                          RavoHTML
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
