package com.photomanager.photomanager.utils

val IntRange.size get() = if (!isEmpty()) last - first + 1 else 0