{{""|xmlheader|unescaped}}
<rss version="2.0">
    <channel>
        <title>{{"rss.title"|get}}</title>
        <link>{{"link"|get}}</link>
        <description>{{"rss.desc"|get}}</description>
        <lastBuildDate>{{.LastBuildDate}}</lastBuildDate>
        <language>{{"rss.lang"|get}}</language>
       {{range .Config}}
        <item>
            <title>{{.Title}}</title>
            <link>{{.Link}}</link>
            <author>{{.Author}}</author>
            <pubDate>{{.Date}}</pubDate>
            <description>{{.Desc}}</description>
        </item>
       {{end}}
    </channel>
</rss>