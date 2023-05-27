package domain

type Source struct {
	Link     string
	Category string
	Name     string
}

var Suara = Source{
	Link:     "https://www.suara.com/rss/bisnis",
	Category: "Bisnis",
	Name:     "Suara.com",
}

var CNN = Source{
	Link:     "https://www.cnnindonesia.com/nasional/rss",
	Category: "Nasional",
	Name:     "CNN Indonesia",
}

var CNBC = Source{
	Link:     "https://www.cnbcindonesia.com/news/rss",
	Category: "Berita Terkini",
	Name:     "CNBC Indonesia",
}

var Republika = Source{
	Link:     "https://www.republika.co.id/rss",
	Category: "Berita Terkini",
	Name:     "Republika",
}

var Tempo = Source{
	Link:     "https://rss.tempo.co/",
	Category: "Berita Terkini",
	Name:     "Tempo",
}

var Antara = Source{
	Link:     "https://www.antaranews.com/rss/terkini.xml",
	Category: "Berita Terkini",
	Name:     "Antara",
}

var Kumparan = Source{
	Link:     "https://lapi.kumparan.com/v2.0/rss",
	Category: "Berita Terkini",
	Name:     "Kumparan",
}

var Okezone = Source{
	Link:     "https://sindikasi.okezone.com/index.php/rss/0/RSS2.0",
	Category: "Berita Terkini",
	Name:     "Okezone",
}

var BBC = Source{
	Link: "http://feeds.bbci.co.uk/indonesia/rss.xml",
	Name: "BBC Indonesia",
}

var Vice = Source{
	Link: "https://www.vice.com/id_id/rss",
	Name: "Vice",
}

var VOA = Source{
	Link: "https://www.voaindonesia.com/api/zmgqoe$moi",
	Name: "VOA Indonesia",
}
