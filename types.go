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
