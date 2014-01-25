{{template "header"}}
<div class="top-nav">
	<ul>
		<li><a href="/" class="on-sel">首页</a></li>
		<li><a href="/tag.html">标签</a></li>
		<li><a href="/category.html">分类</a></li>
		<li><a href="/archive.html">归档</a></li>
		{{range .nav}}
		<li><a href="{{.Href}}" target="{{.Target}}">{{.Name}}</a></li>
		{{end}}
		<li><a href="/rss.xml" class="rss" title="订阅"></a></li>
	</ul>
</div>
<div style="clear:both;"></div>
<div class="main">
	<div class="main-inner">
		<div class="article-list">
		{{range .ar}}
			<div class="article">
				<p class="title"><a href="/articles/{{.Link}}">{{.Title}}</a></p>
				<p class="abstract">&lt;摘要&gt;: {{.Abstract}}&nbsp;&nbsp;<a href="/articles/{{.Link}}">阅读更多</a></p>
				<p class="meta">作者 {{.Author}} | 发布于 {{.Date}} | Tags 
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