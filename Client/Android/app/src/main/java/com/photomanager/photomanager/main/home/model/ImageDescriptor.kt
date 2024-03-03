package com.photomanager.photomanager.main.home.model

import android.net.Uri
import java.util.Date

data class ImageDescriptor(
    val id: String, // Unique identifier
    val uri: Uri,
    val width: Int = 0,
    val height: Int = 0,
    val orientation: String? = null,
    val caption: String = "",
    val shotDate: Date = Date(),
    val lat: Double? = null,
    val lon: Double? = null,
    val camera: String = "",
    val focalLength: String? = null,
    val flash: String? = null,

)