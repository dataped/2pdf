package mapdf

type Filelists struct {
	Aplikasi      string
	User          string
	JenisDocument string
	Halaman       uint8
	Filename      string
}

type LoginHeader struct {
	Login string `reqHeader:"login"`
}

type Ktp struct {
	Province      string `json:"province"`
	District      string `json:"district"`
	Nik           string `json:"nik"`
	Name          string `json:"name"`
	PlaceDob      string `json:"place_dob"`
	Gender        string `json:"gender"`
	Address1      string `json:"address_1"`
	Address2      string `json:"address_2"`
	Address3      string `json:"address_3"`
	Address4      string `json:"address_4"`
	Religion      string `json:"religion"`
	MarriedStatus string `json:"married_status"`
	Occupation    string `json:"occupation"`
	Nationality   string `json:"nationality"`
	ValidUntil    string `json:"valid_until"`
}
