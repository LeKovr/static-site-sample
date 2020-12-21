+++
author = "Hugo Authors"
title = "Include sources"
date = "2020-12-17"
# description = "A brief description of Hugo Shortcodes"
tags = [
    "shortcodes",
    "includes",
]
+++

This page shows how to extend and use Hugo's [Built-in Shortcodes](https://gohugo.io/content-management/shortcodes/#use-hugo-s-built-in-shortcodes)
for file source content including.
<!--more-->

---

## Highlight file

{{< source "media/assign.go" >}}

Note linenumbers with unique anchors

## Gist sample

{{< gist LeKovr 9b079a413e429758e53d hookex.go >}}

## Include file

{{% include "media/assign.md" %}}

## Include file as text

```markdown
{{% include "media/assign.md" %}}
```


## Source of this page

{{< source "content/en/blog/include.md" >}}
