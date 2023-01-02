package main

// const addr = "redis:6379" for docker-compose
const addr = "127.0.0.1:6379"
const password = "somepwd"
const baseURL = "http://somesite.com"
const hlTemplate = `
<!doctype html>
<html lang="en">
    <head>
        <title>pb</title>
        <link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/styles/%s.min.css">
        <link rel="apple-touch-icon" sizes="57x57" href="//img.aimisaka.site/pb/apple-icon-57x57.png">
        <link rel="apple-touch-icon" sizes="60x60" href="//img.aimisaka.site/pb/apple-icon-60x60.png">
        <link rel="apple-touch-icon" sizes="72x72" href="//img.aimisaka.site/pb/apple-icon-72x72.png">
        <link rel="apple-touch-icon" sizes="76x76" href="//img.aimisaka.siteapple-icon-76x76.png">
        <link rel="apple-touch-icon" sizes="114x114" href="//img.aimisaka.site/pb/apple-icon-114x114.png">
        <link rel="apple-touch-icon" sizes="120x120" href="//img.aimisaka.site/pb/apple-icon-120x120.png">
        <link rel="apple-touch-icon" sizes="144x144" href="//img.aimisaka.site/pb/apple-icon-144x144.png">
        <link rel="apple-touch-icon" sizes="152x152" href="//img.aimisaka.site/pb/apple-icon-152x152.png">
        <link rel="apple-touch-icon" sizes="180x180" href="//img.aimisaka.site/pb/apple-icon-180x180.png">
        <link rel="icon" type="image/png" sizes="192x192"  href="//img.aimisaka.site/pb/android-icon-192x192.png">
        <link rel="icon" type="image/png" sizes="32x32" href="//img.aimisaka.site/pb/favicon-32x32.png">
        <link rel="icon" type="image/png" sizes="96x96" href="//img.aimisaka.site/pb/favicon-96x96.png">
        <link rel="icon" type="image/png" sizes="16x16" href="//img.aimisaka.site/pb/favicon-16x16.png">
        <link rel="manifest" href="//img.aimisaka.site/pb/manifest.json">
        <meta name="msapplication-TileColor" content="#ffffff">
        <meta name="msapplication-TileImage" content="//img.aimisaka.site/pb/ms-icon-144x144.png">
        <meta name="theme-color" content="#ffffff">
        <style>
            body {
                color: #c9d1d9; 
                background: #0d1117;
            }
        </style>
    </head>
    <body>
        <pre><code class="language-%s">
%s
        </code></pre>
        <script src="//cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/highlight.min.js"></script>
        <script>hljs.highlightAll();</script>
    </body>
</html>
`
