+++
author = "Hugo Authors"
title = "Easy gallery"
date = "2020-12-17"
# description = "A brief description of Hugo Shortcodes"
tags = [
    "shortcodes",
    "includes",
    "gallery",
]
enableToc=false
+++

This page shows samples of gallery and carousel.
<!--more-->
Code uses 
[hugo-easy-gallery](https://github.com/liwenyip/hugo-easy-gallery).
<!--more-->

---
{{< load-photoswipe >}}

## Carousel image gallery

{{< gallery_tag tag="animals" full="item" thumb="100" width="100" />}}

## Another sample

{{< gallery_tag tag="flowers" />}}

## Source of this page

{{< source "content/en/blog/gallery.md" >}}
