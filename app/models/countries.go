package models

import (
	"github.com/goenning/vat"
)

// Country is a valid country within Fider
type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
	IsEU bool   `json:"isEU"`
}

var countries = []*Country{
	&Country{Code: "AF", Name: "Afghanistan", IsEU: vat.IsEU("AF")},
	&Country{Code: "AX", Name: "Åland Islands", IsEU: vat.IsEU("AX")},
	&Country{Code: "AL", Name: "Albania", IsEU: vat.IsEU("AL")},
	&Country{Code: "DZ", Name: "Algeria", IsEU: vat.IsEU("DZ")},
	&Country{Code: "AS", Name: "American Samoa", IsEU: vat.IsEU("AS")},
	&Country{Code: "AD", Name: "Andorra", IsEU: vat.IsEU("AD")},
	&Country{Code: "AO", Name: "Angola", IsEU: vat.IsEU("AO")},
	&Country{Code: "AI", Name: "Anguilla", IsEU: vat.IsEU("AI")},
	&Country{Code: "AQ", Name: "Antarctica", IsEU: vat.IsEU("AQ")},
	&Country{Code: "AG", Name: "Antigua and Barbuda", IsEU: vat.IsEU("AG")},
	&Country{Code: "AR", Name: "Argentina", IsEU: vat.IsEU("AR")},
	&Country{Code: "AM", Name: "Armenia", IsEU: vat.IsEU("AM")},
	&Country{Code: "AW", Name: "Aruba", IsEU: vat.IsEU("AW")},
	&Country{Code: "AU", Name: "Australia", IsEU: vat.IsEU("AU")},
	&Country{Code: "AT", Name: "Austria", IsEU: vat.IsEU("AT")},
	&Country{Code: "AZ", Name: "Azerbaijan", IsEU: vat.IsEU("AZ")},
	&Country{Code: "BS", Name: "Bahamas", IsEU: vat.IsEU("BS")},
	&Country{Code: "BH", Name: "Bahrain", IsEU: vat.IsEU("BH")},
	&Country{Code: "BD", Name: "Bangladesh", IsEU: vat.IsEU("BD")},
	&Country{Code: "BB", Name: "Barbados", IsEU: vat.IsEU("BB")},
	&Country{Code: "BY", Name: "Belarus", IsEU: vat.IsEU("BY")},
	&Country{Code: "BE", Name: "Belgium", IsEU: vat.IsEU("BE")},
	&Country{Code: "BZ", Name: "Belize", IsEU: vat.IsEU("BZ")},
	&Country{Code: "BJ", Name: "Benin", IsEU: vat.IsEU("BJ")},
	&Country{Code: "BM", Name: "Bermuda", IsEU: vat.IsEU("BM")},
	&Country{Code: "BT", Name: "Bhutan", IsEU: vat.IsEU("BT")},
	&Country{Code: "BO", Name: "Bolivia, Plurinational State of", IsEU: vat.IsEU("BO")},
	&Country{Code: "BQ", Name: "Bonaire, Sint Eustatius and Saba", IsEU: vat.IsEU("BQ")},
	&Country{Code: "BA", Name: "Bosnia and Herzegovina", IsEU: vat.IsEU("BA")},
	&Country{Code: "BW", Name: "Botswana", IsEU: vat.IsEU("BW")},
	&Country{Code: "BV", Name: "Bouvet Island", IsEU: vat.IsEU("BV")},
	&Country{Code: "BR", Name: "Brazil", IsEU: vat.IsEU("BR")},
	&Country{Code: "IO", Name: "British Indian Ocean Territory", IsEU: vat.IsEU("IO")},
	&Country{Code: "BN", Name: "Brunei Darussalam", IsEU: vat.IsEU("BN")},
	&Country{Code: "BG", Name: "Bulgaria", IsEU: vat.IsEU("BG")},
	&Country{Code: "BF", Name: "Burkina Faso", IsEU: vat.IsEU("BF")},
	&Country{Code: "BI", Name: "Burundi", IsEU: vat.IsEU("BI")},
	&Country{Code: "KH", Name: "Cambodia", IsEU: vat.IsEU("KH")},
	&Country{Code: "CM", Name: "Cameroon", IsEU: vat.IsEU("CM")},
	&Country{Code: "CA", Name: "Canada", IsEU: vat.IsEU("CA")},
	&Country{Code: "CV", Name: "Cape Verde", IsEU: vat.IsEU("CV")},
	&Country{Code: "KY", Name: "Cayman Islands", IsEU: vat.IsEU("KY")},
	&Country{Code: "CF", Name: "Central African Republic", IsEU: vat.IsEU("CF")},
	&Country{Code: "TD", Name: "Chad", IsEU: vat.IsEU("TD")},
	&Country{Code: "CL", Name: "Chile", IsEU: vat.IsEU("CL")},
	&Country{Code: "CN", Name: "China", IsEU: vat.IsEU("CN")},
	&Country{Code: "CX", Name: "Christmas Island", IsEU: vat.IsEU("CX")},
	&Country{Code: "CC", Name: "Cocos (Keeling) Islands", IsEU: vat.IsEU("CC")},
	&Country{Code: "CO", Name: "Colombia", IsEU: vat.IsEU("CO")},
	&Country{Code: "KM", Name: "Comoros", IsEU: vat.IsEU("KM")},
	&Country{Code: "CG", Name: "Congo", IsEU: vat.IsEU("CG")},
	&Country{Code: "CD", Name: "Congo, the Democratic Republic of the", IsEU: vat.IsEU("CD")},
	&Country{Code: "CK", Name: "Cook Islands", IsEU: vat.IsEU("CK")},
	&Country{Code: "CR", Name: "Costa Rica", IsEU: vat.IsEU("CR")},
	&Country{Code: "CI", Name: "Côte d'Ivoire", IsEU: vat.IsEU("CI")},
	&Country{Code: "HR", Name: "Croatia", IsEU: vat.IsEU("HR")},
	&Country{Code: "CU", Name: "Cuba", IsEU: vat.IsEU("CU")},
	&Country{Code: "CW", Name: "Curaçao", IsEU: vat.IsEU("CW")},
	&Country{Code: "CY", Name: "Cyprus", IsEU: vat.IsEU("CY")},
	&Country{Code: "CZ", Name: "Czech Republic", IsEU: vat.IsEU("CZ")},
	&Country{Code: "DK", Name: "Denmark", IsEU: vat.IsEU("DK")},
	&Country{Code: "DJ", Name: "Djibouti", IsEU: vat.IsEU("DJ")},
	&Country{Code: "DM", Name: "Dominica", IsEU: vat.IsEU("DM")},
	&Country{Code: "DO", Name: "Dominican Republic", IsEU: vat.IsEU("DO")},
	&Country{Code: "EC", Name: "Ecuador", IsEU: vat.IsEU("EC")},
	&Country{Code: "EG", Name: "Egypt", IsEU: vat.IsEU("EG")},
	&Country{Code: "SV", Name: "El Salvador", IsEU: vat.IsEU("SV")},
	&Country{Code: "GQ", Name: "Equatorial Guinea", IsEU: vat.IsEU("GQ")},
	&Country{Code: "ER", Name: "Eritrea", IsEU: vat.IsEU("ER")},
	&Country{Code: "EE", Name: "Estonia", IsEU: vat.IsEU("EE")},
	&Country{Code: "ET", Name: "Ethiopia", IsEU: vat.IsEU("ET")},
	&Country{Code: "FK", Name: "Falkland Islands (Malvinas)", IsEU: vat.IsEU("FK")},
	&Country{Code: "FO", Name: "Faroe Islands", IsEU: vat.IsEU("FO")},
	&Country{Code: "FJ", Name: "Fiji", IsEU: vat.IsEU("FJ")},
	&Country{Code: "FI", Name: "Finland", IsEU: vat.IsEU("FI")},
	&Country{Code: "FR", Name: "France", IsEU: vat.IsEU("FR")},
	&Country{Code: "GF", Name: "French Guiana", IsEU: vat.IsEU("GF")},
	&Country{Code: "PF", Name: "French Polynesia", IsEU: vat.IsEU("PF")},
	&Country{Code: "TF", Name: "French Southern Territories", IsEU: vat.IsEU("TF")},
	&Country{Code: "GA", Name: "Gabon", IsEU: vat.IsEU("GA")},
	&Country{Code: "GM", Name: "Gambia", IsEU: vat.IsEU("GM")},
	&Country{Code: "GE", Name: "Georgia", IsEU: vat.IsEU("GE")},
	&Country{Code: "DE", Name: "Germany", IsEU: vat.IsEU("DE")},
	&Country{Code: "GH", Name: "Ghana", IsEU: vat.IsEU("GH")},
	&Country{Code: "GI", Name: "Gibraltar", IsEU: vat.IsEU("GI")},
	&Country{Code: "GR", Name: "Greece", IsEU: vat.IsEU("GR")},
	&Country{Code: "GL", Name: "Greenland", IsEU: vat.IsEU("GL")},
	&Country{Code: "GD", Name: "Grenada", IsEU: vat.IsEU("GD")},
	&Country{Code: "GP", Name: "Guadeloupe", IsEU: vat.IsEU("GP")},
	&Country{Code: "GU", Name: "Guam", IsEU: vat.IsEU("GU")},
	&Country{Code: "GT", Name: "Guatemala", IsEU: vat.IsEU("GT")},
	&Country{Code: "GG", Name: "Guernsey", IsEU: vat.IsEU("GG")},
	&Country{Code: "GN", Name: "Guinea", IsEU: vat.IsEU("GN")},
	&Country{Code: "GW", Name: "Guinea-Bissau", IsEU: vat.IsEU("GW")},
	&Country{Code: "GY", Name: "Guyana", IsEU: vat.IsEU("GY")},
	&Country{Code: "HT", Name: "Haiti", IsEU: vat.IsEU("HT")},
	&Country{Code: "HM", Name: "Heard Island and McDonald Islands", IsEU: vat.IsEU("HM")},
	&Country{Code: "VA", Name: "Holy See (Vatican City State)", IsEU: vat.IsEU("VA")},
	&Country{Code: "HN", Name: "Honduras", IsEU: vat.IsEU("HN")},
	&Country{Code: "HK", Name: "Hong Kong", IsEU: vat.IsEU("HK")},
	&Country{Code: "HU", Name: "Hungary", IsEU: vat.IsEU("HU")},
	&Country{Code: "IS", Name: "Iceland", IsEU: vat.IsEU("IS")},
	&Country{Code: "IN", Name: "India", IsEU: vat.IsEU("IN")},
	&Country{Code: "ID", Name: "Indonesia", IsEU: vat.IsEU("ID")},
	&Country{Code: "IR", Name: "Iran, Islamic Republic of", IsEU: vat.IsEU("IR")},
	&Country{Code: "IQ", Name: "Iraq", IsEU: vat.IsEU("IQ")},
	&Country{Code: "IE", Name: "Ireland", IsEU: vat.IsEU("IE")},
	&Country{Code: "IM", Name: "Isle of Man", IsEU: vat.IsEU("IM")},
	&Country{Code: "IL", Name: "Israel", IsEU: vat.IsEU("IL")},
	&Country{Code: "IT", Name: "Italy", IsEU: vat.IsEU("IT")},
	&Country{Code: "JM", Name: "Jamaica", IsEU: vat.IsEU("JM")},
	&Country{Code: "JP", Name: "Japan", IsEU: vat.IsEU("JP")},
	&Country{Code: "JE", Name: "Jersey", IsEU: vat.IsEU("JE")},
	&Country{Code: "JO", Name: "Jordan", IsEU: vat.IsEU("JO")},
	&Country{Code: "KZ", Name: "Kazakhstan", IsEU: vat.IsEU("KZ")},
	&Country{Code: "KE", Name: "Kenya", IsEU: vat.IsEU("KE")},
	&Country{Code: "KI", Name: "Kiribati", IsEU: vat.IsEU("KI")},
	&Country{Code: "KP", Name: "Korea, Democratic People's Republic of", IsEU: vat.IsEU("KP")},
	&Country{Code: "KR", Name: "Korea, Republic of", IsEU: vat.IsEU("KR")},
	&Country{Code: "KW", Name: "Kuwait", IsEU: vat.IsEU("KW")},
	&Country{Code: "KG", Name: "Kyrgyzstan", IsEU: vat.IsEU("KG")},
	&Country{Code: "LA", Name: "Lao People's Democratic Republic", IsEU: vat.IsEU("LA")},
	&Country{Code: "LV", Name: "Latvia", IsEU: vat.IsEU("LV")},
	&Country{Code: "LB", Name: "Lebanon", IsEU: vat.IsEU("LB")},
	&Country{Code: "LS", Name: "Lesotho", IsEU: vat.IsEU("LS")},
	&Country{Code: "LR", Name: "Liberia", IsEU: vat.IsEU("LR")},
	&Country{Code: "LY", Name: "Libya", IsEU: vat.IsEU("LY")},
	&Country{Code: "LI", Name: "Liechtenstein", IsEU: vat.IsEU("LI")},
	&Country{Code: "LT", Name: "Lithuania", IsEU: vat.IsEU("LT")},
	&Country{Code: "LU", Name: "Luxembourg", IsEU: vat.IsEU("LU")},
	&Country{Code: "MO", Name: "Macao", IsEU: vat.IsEU("MO")},
	&Country{Code: "MK", Name: "Macedonia, the former Yugoslav Republic of", IsEU: vat.IsEU("MK")},
	&Country{Code: "MG", Name: "Madagascar", IsEU: vat.IsEU("MG")},
	&Country{Code: "MW", Name: "Malawi", IsEU: vat.IsEU("MW")},
	&Country{Code: "MY", Name: "Malaysia", IsEU: vat.IsEU("MY")},
	&Country{Code: "MV", Name: "Maldives", IsEU: vat.IsEU("MV")},
	&Country{Code: "ML", Name: "Mali", IsEU: vat.IsEU("ML")},
	&Country{Code: "MT", Name: "Malta", IsEU: vat.IsEU("MT")},
	&Country{Code: "MH", Name: "Marshall Islands", IsEU: vat.IsEU("MH")},
	&Country{Code: "MQ", Name: "Martinique", IsEU: vat.IsEU("MQ")},
	&Country{Code: "MR", Name: "Mauritania", IsEU: vat.IsEU("MR")},
	&Country{Code: "MU", Name: "Mauritius", IsEU: vat.IsEU("MU")},
	&Country{Code: "YT", Name: "Mayotte", IsEU: vat.IsEU("YT")},
	&Country{Code: "MX", Name: "Mexico", IsEU: vat.IsEU("MX")},
	&Country{Code: "FM", Name: "Micronesia, Federated States of", IsEU: vat.IsEU("FM")},
	&Country{Code: "MD", Name: "Moldova, Republic of", IsEU: vat.IsEU("MD")},
	&Country{Code: "MC", Name: "Monaco", IsEU: vat.IsEU("MC")},
	&Country{Code: "MN", Name: "Mongolia", IsEU: vat.IsEU("MN")},
	&Country{Code: "ME", Name: "Montenegro", IsEU: vat.IsEU("ME")},
	&Country{Code: "MS", Name: "Montserrat", IsEU: vat.IsEU("MS")},
	&Country{Code: "MA", Name: "Morocco", IsEU: vat.IsEU("MA")},
	&Country{Code: "MZ", Name: "Mozambique", IsEU: vat.IsEU("MZ")},
	&Country{Code: "MM", Name: "Myanmar", IsEU: vat.IsEU("MM")},
	&Country{Code: "NA", Name: "Namibia", IsEU: vat.IsEU("NA")},
	&Country{Code: "NR", Name: "Nauru", IsEU: vat.IsEU("NR")},
	&Country{Code: "NP", Name: "Nepal", IsEU: vat.IsEU("NP")},
	&Country{Code: "NL", Name: "Netherlands", IsEU: vat.IsEU("NL")},
	&Country{Code: "NC", Name: "New Caledonia", IsEU: vat.IsEU("NC")},
	&Country{Code: "NZ", Name: "New Zealand", IsEU: vat.IsEU("NZ")},
	&Country{Code: "NI", Name: "Nicaragua", IsEU: vat.IsEU("NI")},
	&Country{Code: "NE", Name: "Niger", IsEU: vat.IsEU("NE")},
	&Country{Code: "NG", Name: "Nigeria", IsEU: vat.IsEU("NG")},
	&Country{Code: "NU", Name: "Niue", IsEU: vat.IsEU("NU")},
	&Country{Code: "NF", Name: "Norfolk Island", IsEU: vat.IsEU("NF")},
	&Country{Code: "MP", Name: "Northern Mariana Islands", IsEU: vat.IsEU("MP")},
	&Country{Code: "NO", Name: "Norway", IsEU: vat.IsEU("NO")},
	&Country{Code: "OM", Name: "Oman", IsEU: vat.IsEU("OM")},
	&Country{Code: "PK", Name: "Pakistan", IsEU: vat.IsEU("PK")},
	&Country{Code: "PW", Name: "Palau", IsEU: vat.IsEU("PW")},
	&Country{Code: "PS", Name: "Palestinian Territory, Occupied", IsEU: vat.IsEU("PS")},
	&Country{Code: "PA", Name: "Panama", IsEU: vat.IsEU("PA")},
	&Country{Code: "PG", Name: "Papua New Guinea", IsEU: vat.IsEU("PG")},
	&Country{Code: "PY", Name: "Paraguay", IsEU: vat.IsEU("PY")},
	&Country{Code: "PE", Name: "Peru", IsEU: vat.IsEU("PE")},
	&Country{Code: "PH", Name: "Philippines", IsEU: vat.IsEU("PH")},
	&Country{Code: "PN", Name: "Pitcairn", IsEU: vat.IsEU("PN")},
	&Country{Code: "PL", Name: "Poland", IsEU: vat.IsEU("PL")},
	&Country{Code: "PT", Name: "Portugal", IsEU: vat.IsEU("PT")},
	&Country{Code: "PR", Name: "Puerto Rico", IsEU: vat.IsEU("PR")},
	&Country{Code: "QA", Name: "Qatar", IsEU: vat.IsEU("QA")},
	&Country{Code: "RE", Name: "Réunion", IsEU: vat.IsEU("RE")},
	&Country{Code: "RO", Name: "Romania", IsEU: vat.IsEU("RO")},
	&Country{Code: "RU", Name: "Russian Federation", IsEU: vat.IsEU("RU")},
	&Country{Code: "RW", Name: "Rwanda", IsEU: vat.IsEU("RW")},
	&Country{Code: "BL", Name: "Saint Barthélemy", IsEU: vat.IsEU("BL")},
	&Country{Code: "SH", Name: "Saint Helena, Ascension and Tristan da Cunha", IsEU: vat.IsEU("SH")},
	&Country{Code: "KN", Name: "Saint Kitts and Nevis", IsEU: vat.IsEU("KN")},
	&Country{Code: "LC", Name: "Saint Lucia", IsEU: vat.IsEU("LC")},
	&Country{Code: "MF", Name: "Saint Martin (French part)", IsEU: vat.IsEU("MF")},
	&Country{Code: "PM", Name: "Saint Pierre and Miquelon", IsEU: vat.IsEU("PM")},
	&Country{Code: "VC", Name: "Saint Vincent and the Grenadines", IsEU: vat.IsEU("VC")},
	&Country{Code: "WS", Name: "Samoa", IsEU: vat.IsEU("WS")},
	&Country{Code: "SM", Name: "San Marino", IsEU: vat.IsEU("SM")},
	&Country{Code: "ST", Name: "Sao Tome and Principe", IsEU: vat.IsEU("ST")},
	&Country{Code: "SA", Name: "Saudi Arabia", IsEU: vat.IsEU("SA")},
	&Country{Code: "SN", Name: "Senegal", IsEU: vat.IsEU("SN")},
	&Country{Code: "RS", Name: "Serbia", IsEU: vat.IsEU("RS")},
	&Country{Code: "SC", Name: "Seychelles", IsEU: vat.IsEU("SC")},
	&Country{Code: "SL", Name: "Sierra Leone", IsEU: vat.IsEU("SL")},
	&Country{Code: "SG", Name: "Singapore", IsEU: vat.IsEU("SG")},
	&Country{Code: "SX", Name: "Sint Maarten (Dutch part)", IsEU: vat.IsEU("SX")},
	&Country{Code: "SK", Name: "Slovakia", IsEU: vat.IsEU("SK")},
	&Country{Code: "SI", Name: "Slovenia", IsEU: vat.IsEU("SI")},
	&Country{Code: "SB", Name: "Solomon Islands", IsEU: vat.IsEU("SB")},
	&Country{Code: "SO", Name: "Somalia", IsEU: vat.IsEU("SO")},
	&Country{Code: "ZA", Name: "South Africa", IsEU: vat.IsEU("ZA")},
	&Country{Code: "GS", Name: "South Georgia and the South Sandwich Islands", IsEU: vat.IsEU("GS")},
	&Country{Code: "SS", Name: "South Sudan", IsEU: vat.IsEU("SS")},
	&Country{Code: "ES", Name: "Spain", IsEU: vat.IsEU("ES")},
	&Country{Code: "LK", Name: "Sri Lanka", IsEU: vat.IsEU("LK")},
	&Country{Code: "SD", Name: "Sudan", IsEU: vat.IsEU("SD")},
	&Country{Code: "SR", Name: "Suriname", IsEU: vat.IsEU("SR")},
	&Country{Code: "SJ", Name: "Svalbard and Jan Mayen", IsEU: vat.IsEU("SJ")},
	&Country{Code: "SZ", Name: "Swaziland", IsEU: vat.IsEU("SZ")},
	&Country{Code: "SE", Name: "Sweden", IsEU: vat.IsEU("SE")},
	&Country{Code: "CH", Name: "Switzerland", IsEU: vat.IsEU("CH")},
	&Country{Code: "SY", Name: "Syrian Arab Republic", IsEU: vat.IsEU("SY")},
	&Country{Code: "TW", Name: "Taiwan, Province of China", IsEU: vat.IsEU("TW")},
	&Country{Code: "TJ", Name: "Tajikistan", IsEU: vat.IsEU("TJ")},
	&Country{Code: "TZ", Name: "Tanzania, United Republic of", IsEU: vat.IsEU("TZ")},
	&Country{Code: "TH", Name: "Thailand", IsEU: vat.IsEU("TH")},
	&Country{Code: "TL", Name: "Timor-Leste", IsEU: vat.IsEU("TL")},
	&Country{Code: "TG", Name: "Togo", IsEU: vat.IsEU("TG")},
	&Country{Code: "TK", Name: "Tokelau", IsEU: vat.IsEU("TK")},
	&Country{Code: "TO", Name: "Tonga", IsEU: vat.IsEU("TO")},
	&Country{Code: "TT", Name: "Trinidad and Tobago", IsEU: vat.IsEU("TT")},
	&Country{Code: "TN", Name: "Tunisia", IsEU: vat.IsEU("TN")},
	&Country{Code: "TR", Name: "Turkey", IsEU: vat.IsEU("TR")},
	&Country{Code: "TM", Name: "Turkmenistan", IsEU: vat.IsEU("TM")},
	&Country{Code: "TC", Name: "Turks and Caicos Islands", IsEU: vat.IsEU("TC")},
	&Country{Code: "TV", Name: "Tuvalu", IsEU: vat.IsEU("TV")},
	&Country{Code: "UG", Name: "Uganda", IsEU: vat.IsEU("UG")},
	&Country{Code: "UA", Name: "Ukraine", IsEU: vat.IsEU("UA")},
	&Country{Code: "AE", Name: "United Arab Emirates", IsEU: vat.IsEU("AE")},
	&Country{Code: "GB", Name: "United Kingdom", IsEU: vat.IsEU("GB")},
	&Country{Code: "US", Name: "United States", IsEU: vat.IsEU("US")},
	&Country{Code: "UM", Name: "United States Minor Outlying Islands", IsEU: vat.IsEU("UM")},
	&Country{Code: "UY", Name: "Uruguay", IsEU: vat.IsEU("UY")},
	&Country{Code: "UZ", Name: "Uzbekistan", IsEU: vat.IsEU("UZ")},
	&Country{Code: "VU", Name: "Vanuatu", IsEU: vat.IsEU("VU")},
	&Country{Code: "VE", Name: "Venezuela, Bolivarian Republic of", IsEU: vat.IsEU("VE")},
	&Country{Code: "VN", Name: "Viet Nam", IsEU: vat.IsEU("VN")},
	&Country{Code: "VG", Name: "Virgin Islands, British", IsEU: vat.IsEU("VG")},
	&Country{Code: "VI", Name: "Virgin Islands, U.S.", IsEU: vat.IsEU("VI")},
	&Country{Code: "WF", Name: "Wallis and Futuna", IsEU: vat.IsEU("WF")},
	&Country{Code: "EH", Name: "Western Sahara", IsEU: vat.IsEU("EH")},
	&Country{Code: "YE", Name: "Yemen", IsEU: vat.IsEU("YE")},
	&Country{Code: "ZM", Name: "Zambia", IsEU: vat.IsEU("ZM")},
	&Country{Code: "ZW", Name: "Zimbabwe", IsEU: vat.IsEU("ZW")},
}

// GetAllCountries returns a list of all known countries
func GetAllCountries() []*Country {
	return countries
}
