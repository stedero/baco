package model

import (
	"basta/ravo/baco/rio"
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"unicode/utf16"

	"github.com/h2non/filetype"
)

type wrapper struct {
	MijnVelden *ravoRecord `json:"mijnVelden"`
}

// ravoHTML is used for embedded HTML
type ravoHTML struct {
	InnerXML string `xml:",innerxml" json:"html"`
}

// BijlageType defines MIME type and extension
type BijlageType struct {
	Type      string
	Subtype   string
	Value     string
	Extension string
	Filename  string `json:",omitempty"`
	Data      string
	Binary    []byte `json:"-"`
}

// ravoRecord defines a Ravo XML record.
type ravoRecord struct {
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
	Probleemstelling                          ravoHTML
	KlantklachtNr                             string `xml:"KlantklachtNr"`
	NotificationNr                            string `xml:"NotificationNr"`
	VoorstelVoorVerbetering                   ravoHTML
	PrioriteitKlant                           string `xml:"PrioriteitKlant"`
	PrioriteitAanvrager                       int64  `xml:"PrioriteitAanvrager"`
	CommercielePrioriteit                     int64  `xml:"CommercielePrioriteit"`
	MeerkostenGeschat                         int64  `xml:"MeerkostenGeschat"`
	MinderkostenGeschat                       int64  `xml:"MinderkostenGeschat"`
	NotitiesRWVbehandelaarNaam                string `xml:"NotitiesRWVbehandelaarNaam"`
	NotitiesRWVbehandelaar                    ravoHTML
	BetrokkenEngineering                      string         `xml:"BetrokkenEngineering"`
	BetrokkenBedrijfsburo                     string         `xml:"BetrokkenBedrijfsburo"`
	BetrokkenInkoopLogistiek                  string         `xml:"BetrokkenInkoopLogistiek"`
	BetrokkenProductie                        string         `xml:"BetrokkenProductie"`
	BetrokkenSales                            string         `xml:"BetrokkenSales"`
	BetrokkenAfterSales                       string         `xml:"BetrokkenAfterSales"`
	BetrokkenTraining                         string         `xml:"BetrokkenTraining"`
	InschattingUren                           int64          `xml:"InschattingUren"`
	Normuur                                   int64          `xml:"Normuur"`
	InschattingEngTotaal                      int64          `xml:"InschattingEngTotaal"`
	InschattingOverig                         int64          `xml:"InschattingOverig"`
	NormuurOverig                             int64          `xml:"NormuurOverig"`
	InschattingOverigTotaal                   int64          `xml:"InschattingOverigTotaal"`
	TerugkoppelingDoor                        string         `xml:"TerugkoppelingDoor"`
	TerugkoppelingDoorParaaf                  string         `xml:"TerugkoppelingDoorParaaf"`
	TerugkoppelingDoorDatum                   string         `xml:"TerugkoppelingDoorDatum"`
	Geaccepteerd                              bool           `xml:"Geaccepteerd"`
	GeaccepteerdOVProdMan                     bool           `xml:"GeaccepteerdOVProdMan"`
	GeaccepteerdOVOverig                      bool           `xml:"GeaccepteerdOVOverig"`
	GeaccepteerdPrio                          int64          `xml:"GeaccepteerdPrio"`
	GeaccepteerdDatum                         string         `xml:"GeaccepteerdDatum"`
	GeaccepteerdDatumSAP                      string         `xml:"GeaccepteerdDatumSAP"`
	GeaccepteerdOVProdManReden                string         `xml:"GeaccepteerdOVProdManReden"`
	GeaccepteerdOVProdManDatum                string         `xml:"GeaccepteerdOVProdManDatum"`
	GeaccepteerdOVOverigReden                 string         `xml:"GeaccepteerdOVOverigReden"`
	GeaccepteerdOVOverigDatum                 string         `xml:"GeaccepteerdOVOverigDatum"`
	Afgewezen                                 bool           `xml:"Afgewezen"`
	AfgewezenReden                            string         `xml:"AfgewezenReden"`
	AfgewezenDatum                            string         `xml:"AfgewezenDatum"`
	VerkoopgebiedAnders                       string         `xml:"VerkoopgebiedAnders"`
	Cabine                                    bool           `xml:"Cabine"`
	Container                                 bool           `xml:"Container"`
	Chassis                                   bool           `xml:"Chassis"`
	Optie                                     bool           `xml:"Optie"`
	SparePart                                 bool           `xml:"SparePart"`
	Europa                                    bool           `xml:"Europa"`
	NAmerika                                  bool           `xml:"NAmerika"`
	ROW                                       bool           `xml:"ROW"`
	R540                                      bool           `xml:"R540"`
	R560                                      bool           `xml:"R560"`
	SnelheidAnders                            bool           `xml:"SnelheidAnders"`
	Euro6                                     bool           `xml:"EURO-6" json:"EURO-6"`
	StageIII                                  bool           `xml:"STAGE-III" json:"STAGE-III"`
	Tier3USA                                  bool           `xml:"TIER3-USA" json:"TIER3-USA"`
	Tier4USA                                  bool           `xml:"TIER4-USA" json:"TIER4-USA"`
	MotorAandrijving                          bool           `xml:"MotorAandrijving"`
	Hydrauliek                                bool           `xml:"Hydrauliek"`
	Elektra                                   bool           `xml:"Elektra"`
	Constructie                               bool           `xml:"Constructie"`
	Montage                                   bool           `xml:"Montage"`
	Oppervlaktebehandeling                    bool           `xml:"Oppervlaktebehandeling"`
	XMLBijlage                                []string       `xml:"Bijlage" json:"-"`
	Bijlagen                                  []*BijlageType `json:"Bijlagen"`
	MeerMinder                                string         `xml:"MeerMinder"`
	ArtikelTekeningNr                         []struct {
		ArtikelTekening string `xml:"ArtikelTekening"`
		ArtikelNaam     string `xml:"ArtikelNaam"`
	} `xml:"ArtikelTekeningNr"`
	Betreft     string `xml:"Betreft"`
	Afgehandeld bool   `xml:"Afgehandeld"`
}

// Convert XML to JSON
func Convert(r io.Reader, w io.Writer) error {
	rec, err := readRavoRecord(r)
	if err != nil {
		return err
	}
	rec.createBijlagen()
	// rec.hexDumpBijlageHeader()
	// rec.saveBijlagen()
	return rec.writeJSON(w)
}

// readRavoRecord read a Ravo XML record.
func readRavoRecord(r io.Reader) (*ravoRecord, error) {
	var rr ravoRecord
	decoder := xml.NewDecoder(r)
	err := decoder.Decode(&rr)
	if err != nil {
		return nil, err
	}
	return &rr, nil
}

// writeJSON writes a Ravo record as JSON
func (rr *ravoRecord) writeJSON(w io.Writer) error {
	data, err := json.MarshalIndent(wrapper{rr}, "", "  ")
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

func (rr *ravoRecord) createBijlagen() {
	bijlagen := make([]*BijlageType, 0, len(rr.XMLBijlage))
	for seq, xmlBijlage := range rr.XMLBijlage {
		infoPathData, _ := base64.StdEncoding.DecodeString(xmlBijlage)
		binary := infoPathData[64:]
		t, error := filetype.Match(binary)
		if error != nil {
			t = filetype.Unknown
		}
		filename := extractFilename(infoPathData)
		if filename == "" {
			filename = fmt.Sprintf("bijlage%03d.%s", seq+1, t.Extension)
		}
		bijlage := &BijlageType{
			Type:      t.MIME.Type,
			Subtype:   t.MIME.Subtype,
			Value:     t.MIME.Value,
			Extension: t.Extension,
			Filename:  filename,
			Data:      base64.StdEncoding.EncodeToString(binary),
			Binary:    binary}
		bijlagen = append(bijlagen, bijlage)
	}
	rr.Bijlagen = bijlagen
}

func (rr *ravoRecord) hexDumpBijlageHeader() {
	for _, bijlage := range rr.Bijlagen {
		fmt.Printf("\n%s\n", bijlage.Filename)
		fmt.Printf("%s", hex.Dump(bijlage.Binary[:128]))
	}
}

func (rr *ravoRecord) saveBijlagen() {
	for _, bijlage := range rr.Bijlagen {
		writer := rio.CreateFile(fmt.Sprintf("/Users/steef/Desktop/Basta/%s", bijlage.Filename))
		writer.Write(bijlage.Binary)
		writer.Close()
	}
}

func extractFilename(data []byte) string {
	var fileSize uint32
	var fileNameLength uint32
	buf := bytes.NewReader(data[16:])
	err := binary.Read(buf, binary.LittleEndian, &fileSize)
	if err != nil {
		return ""
	}
	err = binary.Read(buf, binary.LittleEndian, &fileNameLength)
	if err != nil {
		return ""
	}
	ints := make([]uint16, fileNameLength-1)
	if err := binary.Read(buf, binary.LittleEndian, &ints); err != nil {
		return ""
	}
	return string(utf16.Decode(ints))
}
