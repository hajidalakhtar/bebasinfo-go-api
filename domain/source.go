package domain

type Source struct {
	Link string
	Name string
}

var Suara = Source{
	Link: "https://www.suara.com/rss/bisnis",
	Name: "Suara.com",
}

var CNN = Source{
	Link: "https://www.cnnindonesia.com/nasional/rss",
	Name: "CNN Indonesia",
}

var CNBC = Source{
	Link: "https://www.cnbcindonesia.com/news/rss",
	Name: "CNBC Indonesia",
}

var Republika = Source{
	Link: "https://www.republika.co.id/rss",
	Name: "Republika",
}

var Tempo = Source{
	Link: "https://rss.tempo.co/",
	Name: "Tempo",
}

var Antara = Source{
	Link: "https://www.antaranews.com/rss/terkini.xml",
	Name: "Antara",
}

var Kumparan = Source{
	Link: "https://lapi.kumparan.com/v2.0/rss",
	Name: "Kumparan",
}

var Okezone = Source{
	Link: "https://sindikasi.okezone.com/index.php/rss",
	Name: "Okezone",
}

var Liputan6 = Source{
	Link: "https://feed.liputan6.com/rss",
	Name: "Liputan6",
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
