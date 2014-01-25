{{template "header"}}
<div class="top-nav">
			<ul>
				<li><a href="/" >首页</a></li>
				<li><a href="/tag.html" class="on-sel">标签</a></li>
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
		<div id="tags-main">
			{{range $k,$v := .tag}}
        	<a href="/tag.html#{{$k}}">{{$k}}<span class="count">×{{$v.Length}}</span></a>
       		{{end}}
		 	<div style="clear:both;"></div>
		 </div>
        <div id="tag-index">
        {{range $k,$v := .tag}}
        	<h1><a name="{{$k}}">{{$k}}</a></h1>
			{{range $v.Articles}}
            <p><a href="/articles/{{.Link}}">{{.Title}}</a></p>
            {{end}}
       	{{end}}
        </div>
		</div>
</div>
{{template "footer"}}