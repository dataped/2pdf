package mapdf

type Filelists struct {
	Aplikasi      string `json:"aplikasi,omitempty" bson:"aplikasi"`
	User          string `json:"user,omitempty" bson:"user"`
	JenisDocument string `json:"jenisdokumen,omitempty" bson:"jenisdokumen"`
	Halaman       uint8  `json:"halaman,omitempty" bson:"halaman"`
	Filename      string `json:"filename,omitempty" bson:"filename"`
	Mime          string `json:"mime,omitempty" bson:"mime"`
	Urutan        uint8  `json:"urutan,omitempty" bson:"urutan"`
	Content       string `json:"content,omitempty" bson:"content"`
}

type LoginHeader struct {
	Login  string `reqHeader:"login"`
	Origin string `reqHeader:"origin"`
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

type ResponseKomarFile struct {
	Code    uint8  `json:"code"`
	Success bool   `json:"success"`
	Status  string `json:"status"`
	Data    string `json:"data"`
}
