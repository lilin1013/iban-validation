package iban

type CountrySetting struct {
	// Length of IBAN code for this country
	Length int

	// Format of BBAN part of IBAN for this country
	Format string
}

var countries = map[string]CountrySetting{
	"AD": CountrySetting{Length: 24, Format: "F04F04A12"},
	"AE": CountrySetting{Length: 23, Format: "F03F16"},
	"AL": CountrySetting{Length: 28, Format: "F08A16"},
	"AT": CountrySetting{Length: 20, Format: "F05F11"},
	"AZ": CountrySetting{Length: 28, Format: "U04A20"},
	"BA": CountrySetting{Length: 20, Format: "F03F03F08F02"},
	"BE": CountrySetting{Length: 16, Format: "F03F07F02"},
	"BG": CountrySetting{Length: 22, Format: "U04F04F02A08"},
	"GB": CountrySetting{Length: 22, Format: "U04F06F08"},
	"SE": CountrySetting{Length: 24, Format: "F03F16F01"},
}

var countryReg = map[string]string{
	"F": "[0-9]",
	"A": "[0-9A-Z]",
	"U": "[A-Z]",
}
