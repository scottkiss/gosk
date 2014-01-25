{{template "header"}}
<div class="top-nav">
	<ul>
		<li><a href="/">Index</a></li>
		<li><a href="/tag.html">Tags</a></li>
		<li><a href="/category.html">Categories</a></li>
		<li><a href="/archive.html">Archive</a></li>
		{{range .nav}}
		<li><a href="{{.Href}}" target="{{.Target}}" >{{.Name}}</a></li>
		{{end}}
		<li><a href="/rss.xml" class="rss" title="feed"></a></li>
	</ul>
</div>
<div style="clear:both;"></div>
<div class="main">
	<div class="main-inner">
		 <h1>{{.p.Title}}</h1>
		 <div id="page-content">{{.p.Content|unescaped}}</div>
	</div>
</div>
{{template "footer"}}