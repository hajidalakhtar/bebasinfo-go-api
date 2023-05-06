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

//export const RSS_CNN_NEWS: string = "https://www.cnnindonesia.com/{type}/rss"
//export const RSS_CNBC_NEWS: string = "https://www.cnbcindonesia.com/{type}/rss"
//export const RSS_REPUBLIKA_NEWS: string = "https://www.republika.co.id/rss/"
//export const RSS_TEMPO_NEWS: string = "http://rss.tempo.co/"
//export const RSS_ANTARA_NEWS: string = "https://www.antaranews.com/rss/"
//export const RSS_KUMPARAN_NEWS: string = "https://lapi.kumparan.com/v2.0/rss/"
//export const RSS_OKEZONE = {
//breaking: "https://sindikasi.okezone.com/index.php/rss/0/RSS2.0",
//news: "https://sindikasi.okezone.com/index.php/rss/1/RSS2.0",
//sport: "https://sindikasi.okezone.com/index.php/rss/2/RSS2.0",
//economy: "https://sindikasi.okezone.com/index.php/rss/11/RSS2.0",
//lifestyle: "https://sindikasi.okezone.com/index.php/rss/12/RSS2.0",
//celebrity: "https://sindikasi.okezone.com/index.php/rss/13/RSS2.0",
//bola: "https://sindikasi.okezone.com/index.php/rss/14/RSS2.0",
//techno: "https://sindikasi.okezone.com/index.php/rss/16/RSS2.0",
//}
//export const RSS_LIPUTAN6: string = "https://feed.liputan6.com/rss"
//export const RSS_BBC: string = "https://feeds.bbci.co.uk/indonesia/{type}/rss.xml"
//export const RSS_TRIBUN: string = "https://{zone}.tribunnews.com/rss/"
//export const RSS_JAWAPOS: string = "https://www.jawapos.com/{type}/feed/"
//export const RSS_VICE: string = "https://www.vice.com/id/rss?locale=id_id{page}"
//export const RSS_SUARA: string = "https://www.suara.com/rss/{type}"
//export const RSS_VOA: string = "https://www.voaindonesia.com/api/zmgqoe$moi"
