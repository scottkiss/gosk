{{template "header"}}
<div class="top-nav">
	<ul>
		<li><a href="/" class="on-sel">Index</a></li>
		<li><a href="/tag.html">Tags</a></li>
		<li><a href="/category.html">Categories</a></li>
		<li><a href="/archive.html">Archive</a></li>
		{{range .nav}}
		<li><a href="{{.Href}}" target="{{.Target}}">{{.Name}}</a></li>
		{{end}}
		<li><a href="/rss.xml" class="rss" title="feed"></a></li>
	</ul>
</div>
<div style="clear:both;"></div>
<div class="main">
	<div class="main-inner">
		<div class="article-list">
		{{range .ar}}
			<div class="article">
				<p class="title"><a href="/articles/{{.Link}}">{{.Title}}</a></p>
				<p class="abstract">&lt;abstract&gt;: {{.Abstract}}&nbsp;&nbsp;<a href="/articles/{{.Link}}">Read more</a></p>
				<p class="meta">Author {{.Author}} | Posted {{.Date}} | Tags 
				{{range .Tags}}
				<a class="tag" href="/tag.html#{{.Name}}">{{.Name}}</a>
				{{end}}
				</p>
			</div> 	
		{{end}}
		</div>
	</div>
</div>
{{template "footer"}}