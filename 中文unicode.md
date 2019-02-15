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
