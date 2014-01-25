{{template "header"}}
<div class="top-nav">
			<ul>
				<li><a href="/" >首页</a></li>
				<li><a href="/tag.html">标签</a></li>
				<li><a href="/category.html">分类</a></li>
				<li><a href="/archive.html">归档</a></li>
				{{range .Nav}}
				<li><a href="{{.Href}}" target="{{.Target}}">{{.Name}}</a></li>
				{{end}}
				<li><a href="/rss.xml" class="rss" title="订阅"></a></li>
			</ul>
</div>
<div style="clear:both;"></div>
<div class="main">
	<div class="main-inner">
		 <div id="article-title">
		 	<a href="/{{.Link}}">{{.Title}}</a>
		 </div>
		 <div id="article-meta">作者 {{.Author}}  | 发布于 {{.Date}} </div>

		  <div id="article-tags">
		  {{range .Tags}}
		  <a class="tag" href="/tag.html#{{.Name}}">
		  {{.Name}}</a> 
		  {{end}}
		  </div>
		 <div id="article-content"> {{.Content|unescaped}} </div>

		 <div class="bshare-custom icon-medium"><a title="分享到" href="http://www.bShare.cn/" id="bshare-shareto" class="bshare-more">分享到</a><a title="分享到QQ空间" class="bshare-qzone"></a><a title="分享到新浪微博" class="bshare-sinaminiblog"></a><a title="分享到人人网" class="bshare-renren"></a><a title="分享到腾讯微博" class="bshare-qqmb"></a><a title="分享到网易微博" class="bshare-neteasemb"></a><a title="更多平台" class="bshare-more bshare-more-icon more-style-addthis"></a><span class="BSHARE_COUNT bshare-share-count">0</span></div><script type="text/javascript" charset="utf-8" src="http://static.bshare.cn/b/buttonLite.js#style=-1&amp;uuid=&amp;pophcol=2&amp;lang=zh"></script><script type="text/javascript" charset="utf-8" src="http://static.bshare.cn/b/bshareC0.js"></script>
	</div>
	
	<div class="comments">
	<hr>
    <!-- duoshuo div -->
	<div class="well">  
	<div class="ds-thread"></div></div><br/>
	<!-- Duoshuo Comment BEGIN -->
	<script type="text/javascript">  
	var duoshuoQuery = {short_name:"{{"duoshuoShortname"|get}}"};  
	(function() {    
	var ds = document.createElement('script');    
	ds.type = 'text/javascript';
	ds.async = true;    
	ds.src = 'http://static.duoshuo.com/embed.js';    
	ds.charset = 'UTF-8';    
	(document.getElementsByTagName('head')[0]||
 	document.getElementsByTagName('body')[0]).appendChild(ds); 
	})(); 
 </script>
 <!-- Duoshuo Comment END -->
</div>
</div>

</div>

<script type="text/javascript" src="/assets/themes/{{"theme"|get}}/jquery.js"></script>
<script type="text/javascript">
 		$(function(){
 			$('.prettyprint').each(function(){
 				$(this).html(htmlEncode($(this).html()));
 			});
 		});
		function htmlEncode(str){
  			var s = "";
  			if (str.length == 0) return "";
  				s = str.replace(/&/g, "&gt;");
  				s = s.replace(/</g, "&lt;");
  				s = s.replace(/>/g, "&gt;");
	  			s = s.replace(/ /g, "&nbsp;");
	  			s = s.replace(/\'/g, "&#39;");
	 			s = s.replace(/\"/g, "&quot;"); 
	  			s = s.replace(/\n/g, "<br>");
  				return s;   
			}
 </script>
{{template "footer"}}