```
http://www.pcre.org/current/doc/html/pcre2syntax.html#TOC1
\N{U+hh..} character with Unicode code point hh.. (Unicode mode only)
```


```
\N{U+hh..} is synonymous with \x{hh..}
```

输入

```
<?php
$isMatch = preg_match_all('/[\x{4e00}]+/u', '你好..一23',$match);
if($isMatch && is_array($match[0]) ){
  $result =   implode("",$match[0]);
};
print_r($result);
```


中文匹配的正则总是很麻烦,而且不同的语言写法也很不一样,网上的答案也是非常多,有些是utf-8 的正则,有些是gbk的,有的是unicode什么的.如果没有弄清楚这些概念,会有很多特别的问题出现.
