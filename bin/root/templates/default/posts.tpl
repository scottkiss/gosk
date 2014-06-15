{{template "header"}}
<div class="top-nav">
			<ul>
				<li><a href="/" >Index</a></li>
				<li><a href="/tag.html">Tags</a></li>
				<li><a href="/category.html">Categories</a></li>
				<li><a href="/archive.html">Archive</a></li>
				{{range .Nav}}
				<li><a href="{{.Href}}" target="{{.Target}}">{{.Name}}</a></li>
				{{end}}
				<li><a href="/rss.xml" class="rss" title="feed"></a></li>
			</ul>
</div>
<div style="clear:both;"></div>
<div class="main">
	<div class="main-inner">
		 <div id="article-title">
		 	<a href="/{{.fi.Link}}">{{.fi.Title}}</a>
		 </div>
		 <div id="article-meta">Author {{.fi.Author}}  | Posted {{.fi.Date}} </div>

		  <div id="article-tags">
		  {{range .Tags}}
		  <a class="tag" href="/tag.html#{{.Name}}">
		  {{.Name}}</a> 
		  {{end}}
		  </div>
		 <div id="article-content"> {{.fi.Content|unescaped}} </div>
		<hr/>
	<div class="comments">
	 <h3>Comments</h3>
		<div id="disqus_thread"></div>
    		<script type="text/javascript">
        	/* * * CONFIGURATION VARIABLES: EDIT BEFORE PASTING INTO YOUR WEBPAGE * * */
        	var disqus_shortname = {{"disqus_shortname"|get}}; // required: replace example with your forum shortname

        /* * * DON'T EDIT BELOW THIS LINE * * */
        (function() {
            var dsq = document.createElement('script'); dsq.type = 'text/javascript'; dsq.async = true;
            dsq.src = '//' + disqus_shortname + '.disqus.com/embed.js';
            (document.getElementsByTagName('head')[0] || document.getElementsByTagName('body')[0]).appendChild(dsq);
        })();
    </script>
    <noscript>Please enable JavaScript to view the <a href="http://disqus.com/?ref_noscript">comments powered by Disqus.</a></noscript>
    <a href="http://disqus.com" class="dsq-brlink">comments powered by <span class="logo-disqus">Disqus</span></a>
    
	</div>
</div>

</div>

<script type="text/javascript" src="/assets/themes/{{"theme"|get}}/jquery.js"></script>

{{template "footer"}}