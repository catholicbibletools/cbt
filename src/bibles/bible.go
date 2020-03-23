package bibles

/*import "plugin"

func loadPluginSymbol(path string, name string) interface{} {
	plug, err1 := plugin.Open(path)
	if err1 != nil {
		panic(err1)
	}
	sb, err2 := plug.Lookup(name)
	if err2 != nil {
		panic(err2)
	}
	return sb
}

func loadBibleBooksFrom(path string) BibleBook {
	sb := loadPluginSymbol(path+"/books.so", "BibleBook")
	var bb BibleBook
	bb, ok := sb.(BibleBook)
	if !ok {
		panic("symbol BibleBook not found at " + path + "/books.so")
	}
	return bb
}

func loadBibleInfoFrom(path string, code string) BibleInfo {
	sb := loadPluginSymbol(path+"/"+code+"/bible.so", "BibleInfo")
	var bi BibleInfo
	bi, ok := sb.(BibleInfo)
	if !ok {
		panic("symbol BibleInfo not found at " + path + "/" + code + "/bible.so")
	}
	return bi
}

// interfaces

type BibleBook interface {
	LoadVPGNChapters(CurrentBook *Book)
	LoadVPEXChapters(CurrentBook *Book)
	LoadVPLVChapters(CurrentBook *Book)
	LoadVPNMChapters(CurrentBook *Book)
	LoadVPDTChapters(CurrentBook *Book)
	LoadVHIEChapters(CurrentBook *Book)
	LoadVHIDChapters(CurrentBook *Book)
	LoadVH1SChapters(CurrentBook *Book)
	LoadVH2SChapters(CurrentBook *Book)
	LoadVH1RChapters(CurrentBook *Book)
	LoadVH2RChapters(CurrentBook *Book)
	LoadVH1PChapters(CurrentBook *Book)
	LoadVH2PChapters(CurrentBook *Book)
	LoadVHEDChapters(CurrentBook *Book)
	LoadVHNHChapters(CurrentBook *Book)
	LoadVH1MChapters(CurrentBook *Book)
	LoadVH2MChapters(CurrentBook *Book)
	LoadVMISChapters(CurrentBook *Book)
	LoadVMIEChapters(CurrentBook *Book)
	LoadVMEZChapters(CurrentBook *Book)
	LoadVNOSChapters(CurrentBook *Book)
	LoadVNILChapters(CurrentBook *Book)
	LoadVNAMChapters(CurrentBook *Book)
	LoadVNABChapters(CurrentBook *Book)
	LoadVNINChapters(CurrentBook *Book)
	LoadVNMIChapters(CurrentBook *Book)
	LoadVNNHChapters(CurrentBook *Book)
	LoadVNHAChapters(CurrentBook *Book)
	LoadVNSOChapters(CurrentBook *Book)
	LoadVNAGChapters(CurrentBook *Book)
	LoadVNZAChapters(CurrentBook *Book)
	LoadVNMLChapters(CurrentBook *Book)
	LoadVMDNChapters(CurrentBook *Book)
	LoadVSIBChapters(CurrentBook *Book)
	LoadVSPRChapters(CurrentBook *Book)
	LoadVSEEChapters(CurrentBook *Book)
	LoadVSCCChapters(CurrentBook *Book)
	LoadVHRTChapters(CurrentBook *Book)
	LoadVMLMChapters(CurrentBook *Book)
	LoadVHETChapters(CurrentBook *Book)
	LoadVHTBChapters(CurrentBook *Book)
	LoadVHIHChapters(CurrentBook *Book)
	LoadVMBRChapters(CurrentBook *Book)
	LoadVSSPChapters(CurrentBook *Book)
	LoadVSEIChapters(CurrentBook *Book)
	LoadVSPSChapters(CurrentBook *Book)
	LoadNEMTChapters(CurrentBook *Book)
	LoadNEMCChapters(CurrentBook *Book)
	LoadNELCChapters(CurrentBook *Book)
	LoadNEIOChapters(CurrentBook *Book)
	LoadNTAAChapters(CurrentBook *Book)
	LoadNPRMChapters(CurrentBook *Book)
	LoadNP1CChapters(CurrentBook *Book)
	LoadNP2CChapters(CurrentBook *Book)
	LoadNPGAChapters(CurrentBook *Book)
	LoadNPEPChapters(CurrentBook *Book)
	LoadNPPPChapters(CurrentBook *Book)
	LoadNPCLChapters(CurrentBook *Book)
	LoadNPPMChapters(CurrentBook *Book)
	LoadNP1EChapters(CurrentBook *Book)
	LoadNP2EChapters(CurrentBook *Book)
	LoadNP1IChapters(CurrentBook *Book)
	LoadNP2IChapters(CurrentBook *Book)
	LoadNPTTChapters(CurrentBook *Book)
	LoadNTHBChapters(CurrentBook *Book)
	LoadNCIBChapters(CurrentBook *Book)
	LoadNC1PChapters(CurrentBook *Book)
	LoadNC2PChapters(CurrentBook *Book)
	LoadNCIDChapters(CurrentBook *Book)
	LoadNC1IChapters(CurrentBook *Book)
	LoadNC2IChapters(CurrentBook *Book)
	LoadNC3IChapters(CurrentBook *Book)
	LoadNTAIChapters(CurrentBook *Book)
}

type BibleInfo interface {
	LoadInfo(CurrentBible *Bible)
}*/

// bible

func IsValidBible(code string) bool {
	switch code {
	case "en.ccb", "en.drbcr", "es.eldpdd", "es.lbl", "es.lsb", "pt.bs", "fr.lbdcc", "it.lb", "la.nvbse":
		return true
	}
	return false
}

type Bible struct {
	code string
	name string
	//path  string
	Codes []string
	Names []string
	Books []Book
}

func NewBible(code string) *Bible {
	result := &Bible{}
	switch code {
	case "en.ccb":
		EN_CCB_LoadInfo(result)
	case "en.drbcr":
		EN_DRBCR_LoadInfo(result)
	case "es.eldpdd":
		ES_ELDPDD_LoadInfo(result)
	case "es.lbl":
		ES_LBL_LoadInfo(result)
	case "es.lsb":
		ES_LSB_LoadInfo(result)
	case "pt.bs":
		PT_BS_LoadInfo(result)
	case "fr.lbdcc":
		FR_LBDCC_LoadInfo(result)
	case "it.lb":
		IT_LB_LoadInfo(result)
	case "la.nvbse":
		LA_NVBSE_LoadInfo(result)
	}
	return result
}

/*func NewBible(path string, code string) *Bible {
	result := &Bible{}
	switch code {
	case "en.ccb":
		EN_CCB_LoadInfo(result)
	case "en.drbcr":
		EN_DRBCR_LoadInfo(result)
	case "es.eldpdd":
		ES_ELDPDD_LoadInfo(result)
	case "es.lbl":
		ES_LBL_LoadInfo(result)
	case "es.lsb":
		ES_LSB_LoadInfo(result)
	case "pt.bs":
		PT_BS_LoadInfo(result)
	case "fr.lbdcc":
		FR_LBDCC_LoadInfo(result)
	case "it.lb":
		IT_LB_LoadInfo(result)
	case "la.nvbse":
		LA_NVBSE_LoadInfo(result)
	default:
		loadBibleInfoFrom(path, code).LoadInfo(result)
	}
	result.path = path + "/" + result.code
	return result
}*/

func (me *Bible) GetCode() string {
	return me.code
}

func (me *Bible) GetName() string {
	return me.name
}

func (me *Bible) SetInfo(code string, name string) {
	me.code = code
	me.name = name
}

func (me *Bible) AddBook(code string, name string) {
	me.Codes = append(me.Codes, code)
	me.Names = append(me.Names, name)
}

/*func (me *Bible) loadBibleBooks() BibleBook {
	return loadBibleBooksFrom(me.path)
}*/

func (me *Bible) LoadBooks() {
	switch me.code {
	case "en.ccb":
		EN_CCB_LoadBooks(me)
		return
	case "en.drbcr":
		EN_DRBCR_LoadBooks(me)
		return
	case "es.eldpdd":
		ES_ELDPDD_LoadBooks(me)
		return
	case "es.lbl":
		ES_LBL_LoadBooks(me)
		return
	case "es.lsb":
		ES_LSB_LoadBooks(me)
		return
	case "pt.bs":
		PT_BS_LoadBooks(me)
		return
	case "fr.lbdcc":
		FR_LBDCC_LoadBooks(me)
		return
	case "it.lb":
		IT_LB_LoadBooks(me)
		return
	case "la.nvbse":
		LA_NVBSE_LoadBooks(me)
		return
	}
	/*bb := me.loadBibleBooks()
	for x := range me.Codes {
		var cb *Book = NewBook(me.Codes[x], me.Names[x])
		switch me.Codes[x] {
		case "VPGN":
			bb.LoadVPGNChapters(cb)
		case "VPEX":
			bb.LoadVPEXChapters(cb)
		case "VPLV":
			bb.LoadVPLVChapters(cb)
		case "VPNM":
			bb.LoadVPNMChapters(cb)
		case "VPDT":
			bb.LoadVPDTChapters(cb)
		case "VHIE":
			bb.LoadVHIEChapters(cb)
		case "VHID":
			bb.LoadVHIDChapters(cb)
		case "VH1S":
			bb.LoadVH1SChapters(cb)
		case "VH2S":
			bb.LoadVH2SChapters(cb)
		case "VH1R":
			bb.LoadVH1RChapters(cb)
		case "VH2R":
			bb.LoadVH2RChapters(cb)
		case "VH1P":
			bb.LoadVH1PChapters(cb)
		case "VH2P":
			bb.LoadVH2PChapters(cb)
		case "VHED":
			bb.LoadVHEDChapters(cb)
		case "VHNH":
			bb.LoadVHNHChapters(cb)
		case "VH1M":
			bb.LoadVH1MChapters(cb)
		case "VH2M":
			bb.LoadVH2MChapters(cb)
		case "VMIS":
			bb.LoadVMISChapters(cb)
		case "VMIE":
			bb.LoadVMIEChapters(cb)
		case "VMEZ":
			bb.LoadVMEZChapters(cb)
		case "VNOS":
			bb.LoadVNOSChapters(cb)
		case "VNIL":
			bb.LoadVNILChapters(cb)
		case "VNAM":
			bb.LoadVNAMChapters(cb)
		case "VNAB":
			bb.LoadVNABChapters(cb)
		case "VNIN":
			bb.LoadVNINChapters(cb)
		case "VNMI":
			bb.LoadVNMIChapters(cb)
		case "VNNH":
			bb.LoadVNNHChapters(cb)
		case "VNHA":
			bb.LoadVNHAChapters(cb)
		case "VNSO":
			bb.LoadVNSOChapters(cb)
		case "VNAG":
			bb.LoadVNAGChapters(cb)
		case "VNZA":
			bb.LoadVNZAChapters(cb)
		case "VNML":
			bb.LoadVNMLChapters(cb)
		case "VMDN":
			bb.LoadVMDNChapters(cb)
		case "VSIB":
			bb.LoadVSIBChapters(cb)
		case "VSPR":
			bb.LoadVSPRChapters(cb)
		case "VSEE":
			bb.LoadVSEEChapters(cb)
		case "VSCC":
			bb.LoadVSCCChapters(cb)
		case "VHRT":
			bb.LoadVHRTChapters(cb)
		case "VMLM":
			bb.LoadVMLMChapters(cb)
		case "VHET":
			bb.LoadVHETChapters(cb)
		case "VHTB":
			bb.LoadVHTBChapters(cb)
		case "VHIH":
			bb.LoadVHIHChapters(cb)
		case "VMBR":
			bb.LoadVMBRChapters(cb)
		case "VSSP":
			bb.LoadVSSPChapters(cb)
		case "VSEI":
			bb.LoadVSEIChapters(cb)
		case "VSPS":
			bb.LoadVSPSChapters(cb)
		case "NEMT":
			bb.LoadNEMTChapters(cb)
		case "NEMC":
			bb.LoadNEMCChapters(cb)
		case "NELC":
			bb.LoadNELCChapters(cb)
		case "NEIO":
			bb.LoadNEIOChapters(cb)
		case "NTAA":
			bb.LoadNTAAChapters(cb)
		case "NPRM":
			bb.LoadNPRMChapters(cb)
		case "NP1C":
			bb.LoadNP1CChapters(cb)
		case "NP2C":
			bb.LoadNP2CChapters(cb)
		case "NPGA":
			bb.LoadNPGAChapters(cb)
		case "NPEP":
			bb.LoadNPEPChapters(cb)
		case "NPPP":
			bb.LoadNPPPChapters(cb)
		case "NPCL":
			bb.LoadNPCLChapters(cb)
		case "NPPM":
			bb.LoadNPPMChapters(cb)
		case "NP1E":
			bb.LoadNP1EChapters(cb)
		case "NP2E":
			bb.LoadNP2EChapters(cb)
		case "NP1I":
			bb.LoadNP1IChapters(cb)
		case "NP2I":
			bb.LoadNP2IChapters(cb)
		case "NPTT":
			bb.LoadNPTTChapters(cb)
		case "NTHB":
			bb.LoadNTHBChapters(cb)
		case "NCIB":
			bb.LoadNCIBChapters(cb)
		case "NC1P":
			bb.LoadNC1PChapters(cb)
		case "NC2P":
			bb.LoadNC2PChapters(cb)
		case "NCID":
			bb.LoadNCIDChapters(cb)
		case "NC1I":
			bb.LoadNC1IChapters(cb)
		case "NC2I":
			bb.LoadNC2IChapters(cb)
		case "NC3I":
			bb.LoadNC3IChapters(cb)
		case "NTAI":
			bb.LoadNTAIChapters(cb)
		}
		me.Books = append(me.Books, *cb)
	}*/
}

func (me *Bible) GetBook(code string) *Book {
	for n := range me.Books {
		var b = me.Books[n]
		if b.Code == code {
			return &b
		}
	}
	return nil
}
