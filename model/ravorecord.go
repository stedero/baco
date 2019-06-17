package model

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
)

type wrapper struct {
	MijnVelden *RavoRecord `json:"mijnVelden"`
}

// RavoHTML is used for embeded HTML
type RavoHTML struct {
	InnerXML string `xml:",innerxml" json:"html"`
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
	VeiligheidsrisicoAantal                   int64  `xml:"VeiligheidsrisicoAantal"`
	MilieuschadeRisicoAantal                  int64  `xml:"MilieuschadeRisicoAantal"`
	KlantklachtAantal                         int64  `xml:"KlantklachtAantal"`
	WettelijkVereistAantal                    int64  `xml:"WettelijkVereistAantal"`
	GarantieAanspraakAantal                   int64  `xml:"GarantieAanspraakAantal"`
	NietConformFunctieEisenAantal             int64  `xml:"NietConformFunctieEisenAantal"`
	NietConformKwaliteitEisenAantal           int64  `xml:"NietConformKwaliteitEisenAantal"`
	EfficiencyVeranderingMontageMinutenAantal int64  `xml:"EfficiencyVeranderingMontageMinutenAantal"`
	KostprijsVeranderingAantal                int64  `xml:"KostprijsVeranderingAantal"`
	ProductStandaardisatieAantal              int64  `xml:"ProductStandaardisatieAantal"`
	ProductwijzigingLeverancierAantal         int64  `xml:"ProductwijzigingLeverancierAantal"`
	ProductwijzigingInternAantal              int64  `xml:"ProductwijzigingInternAantal"`
	AndersAantal                              int64  `xml:"AndersAantal"`
	AndersWat                                 string `xml:"AndersWat"`
	VeiligheidsrisicoImpact                   string `xml:"VeiligheidsrisicoImpact"`
	MilieuschadeRisicoImpact                  string `xml:"MilieuschadeRisicoImpact"`
	KlantklachtImpact                         string `xml:"KlantklachtImpact"`
	WettelijkVereistImpact                    string `xml:"WettelijkVereistImpact"`
	GarantieAanspraakImpact                   string `xml:"GarantieAanspraakImpact"`
	NietConformFunctieEisenImpact             string `xml:"NietConformFunctieEisenImpact"`
	NietConformKwaliteitEisenImpact           int64  `xml:"NietConformKwaliteitEisenImpact"`
	AndereLeverancier                         string `xml:"AndereLeverancier"`
	AndereLeverancier2                        string `xml:"AndereLeverancier2"`
	VoorNadeel                                int64  `xml:"VoorNadeel"`
	VoorNadeel2                               int64  `xml:"VoorNadeel2"`
	Probleemstelling                          RavoHTML
	KlantklachtNr                             string `xml:"KlantklachtNr"`
	NotificationNr                            string `xml:"NotificationNr"`
	VoorstelVoorVerbetering                   RavoHTML
	PrioriteitKlant                           string `xml:"PrioriteitKlant"`
	PrioriteitAanvrager                       int64  `xml:"PrioriteitAanvrager"`
	CommercielePrioriteit                     int64  `xml:"CommercielePrioriteit"`
	MeerkostenGeschat                         int64  `xml:"MeerkostenGeschat"`
	MinderkostenGeschat                       int64  `xml:"MinderkostenGeschat"`
	NotitiesRWVbehandelaarNaam                string `xml:"NotitiesRWVbehandelaarNaam"`
	NotitiesRWVbehandelaar                    string `xml:"NotitiesRWVbehandelaar"`
	BetrokkenEngineering                      string `xml:"BetrokkenEngineering"`
	BetrokkenBedrijfsburo                     string `xml:"BetrokkenBedrijfsburo"`
	BetrokkenInkoopLogistiek                  string `xml:"BetrokkenInkoopLogistiek"`
	BetrokkenProductie                        string `xml:"BetrokkenProductie"`
	BetrokkenSales                            string `xml:"BetrokkenSales"`
	BetrokkenAfterSales                       string `xml:"BetrokkenAfterSales"`
	BetrokkenTraining                         string `xml:"BetrokkenTraining"`
	InschattingUren                           int64  `xml:"InschattingUren"`
	Normuur                                   int64  `xml:"Normuur"`
	InschattingEngTotaal                      int64  `xml:"InschattingEngTotaal"`
	InschattingOverig                         int64  `xml:"InschattingOverig"`
	NormuurOverig                             int64  `xml:"NormuurOverig"`
	InschattingOverigTotaal                   int64  `xml:"InschattingOverigTotaal"`
	TerugkoppelingDoor                        string `xml:"TerugkoppelingDoor"`
	TerugkoppelingDoorParaaf                  string `xml:"TerugkoppelingDoorParaaf"`
	TerugkoppelingDoorDatum                   string `xml:"TerugkoppelingDoorDatum"`
	Geaccepteerd                              bool   `xml:"Geaccepteerd"`
	GeaccepteerdOVProdMan                     bool   `xml:"GeaccepteerdOVProdMan"`
	GeaccepteerdOVOverig                      bool   `xml:"GeaccepteerdOVOverig"`
	GeaccepteerdPrio                          int64  `xml:"GeaccepteerdPrio"`
	GeaccepteerdDatum                         string `xml:"GeaccepteerdDatum"`
	GeaccepteerdDatumSAP                      string `xml:"GeaccepteerdDatumSAP"`
	GeaccepteerdOVProdManReden                string `xml:"GeaccepteerdOVProdManReden"`
	GeaccepteerdOVProdManDatum                string `xml:"GeaccepteerdOVProdManDatum"`
	GeaccepteerdOVOverigReden                 string `xml:"GeaccepteerdOVOverigReden"`
	GeaccepteerdOVOverigDatum                 string `xml:"GeaccepteerdOVOverigDatum"`
	Afgewezen                                 bool   `xml:"Afgewezen"`
	AfgewezenReden                            string `xml:"AfgewezenReden"`
	AfgewezenDatum                            string `xml:"AfgewezenDatum"`
	VerkoopgebiedAnders                       string `xml:"VerkoopgebiedAnders"`
	Cabine                                    bool   `xml:"Cabine"`
	Container                                 bool   `xml:"Container"`
	Chassis                                   bool   `xml:"Chassis"`
	Optie                                     bool   `xml:"Optie"`
	SparePart                                 bool   `xml:"SparePart"`
	Europa                                    bool   `xml:"Europa"`
	NAmerika                                  bool   `xml:"NAmerika"`
	ROW                                       bool   `xml:"ROW"`
	R540                                      bool   `xml:"R540"`
	R560                                      bool   `xml:"R560"`
	SnelheidAnders                            bool   `xml:"SnelheidAnders"`
	Euro6                                     bool   `xml:"EURO-6" json:"EURO-6"`
	StageIII                                  bool   `xml:"STAGE-III" json:"STAGE-III"`
	Tier3USA                                  bool   `xml:"TIER3-USA" json:"TIER3-USA"`
	Tier4USA                                  bool   `xml:"TIER4-USA" json:"TIER4-USA"`
	MotorAandrijving                          bool   `xml:"MotorAandrijving"`
	Hydrauliek                                bool   `xml:"Hydrauliek"`
	Elektra                                   bool   `xml:"Elektra"`
	Constructie                               bool   `xml:"Constructie"`
	Montage                                   bool   `xml:"Montage"`
	Oppervlaktebehandeling                    bool   `xml:"Oppervlaktebehandeling"`
	Bijlage                                   string `xml:"Bijlage"`
	MeerMinder                                string `xml:"MeerMinder"`
	ArtikelTekeningNr                         struct {
		ArtikelTekening string `xml:"ArtikelTekening"`
		ArtikelNaam     string `xml:"ArtikelNaam"`
	} `xml:"ArtikelTekeningNr"`
	Betreft     string `xml:"Betreft"`
	Afgehandeld bool   `xml:"Afgehandeld"`
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
	data, err := json.MarshalIndent(wrapper{rr}, "", "    ")
	if err != nil {
		log.Fatalf("failed to marshal to JSON: %v", err)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Fatalf("failed to write JSON: %v", err)
	}
}
