{{define "header"}}
<!DOCTYPE html>
<html>   
	<head>        
	<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"/>        
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8"/>        
	<meta name="description" content="{{"meta.description"|get}}"/>        
	<meta name="keywords" content="{{"meta.keywords"|get}}"/>        
	<meta name="author" content="{{"meta.author"|get}}"/>        
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1"/>
	<title>{{"title"|get}} - {{"subtitle"|get}}</title>
	<link rel="stylesheet" href="/assets/themes/{{"theme"|get}}/plugin/prettify/normalize.css"/>
	<link rel="stylesheet" href="/assets/themes/{{"theme"|get}}/plugin/prettify/prettify_{{"codetheme"|get}}.css"/>
	<link rel="stylesheet" href="/assets/themes/{{"theme"|get}}/main.css"/>
	<link rel="shortcut icon" href="/fav.ico"/>
	<script type="text/javascript" src="/assets/themes/{{"theme"|get}}/plugin/prettify/prettify.js"></script>
	</head>
	<body onload="prettyPrint()">
		<div id="header">
			<div id="header-inner"> 
				<div id="title"><a href="/">{{"title"|get}}</a>
				</div>            
				<div id="subtitle">{{"subtitle"|get}}</div>        
			</div>   
		</div>
{{end}}