package com.photomanager.photomanager.utils

import android.net.Uri
import android.os.Build
import java.net.URLEncoder

fun Uri.encode(): String = if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.TIRAMISU) {
    URLEncoder.encode(toString(), Charsets.UTF_8)
} else {
    @Suppress("DEPRECATION")
    URLEncoder.encode(toString())
}
