package com.photomanager.photomanager.main.home.repository

import android.net.Uri
import androidx.exifinterface.media.ExifInterface
import com.photomanager.photomanager.main.home.model.ImageDescriptor
import java.text.SimpleDateFormat
import java.util.Date
import java.util.Locale
import javax.inject.Inject

class ImageProcessorRepoImpl @Inject constructor() : ImageProcessorRepo {
    override fun processImage(uri: Uri): ImageDescriptor {
        val exif = ExifInterface(uri.path!!)
        val lat = exif.getAttribute(ExifInterface.TAG_GPS_LATITUDE)
        val lon = exif.getAttribute(ExifInterface.TAG_GPS_LONGITUDE)
        val date = exif.getAttribute(ExifInterface.TAG_DATETIME)
        val width = exif.getAttribute(ExifInterface.TAG_IMAGE_WIDTH)
        val height = exif.getAttribute(ExifInterface.TAG_IMAGE_LENGTH)
        val orientation = exif.getAttribute(ExifInterface.TAG_ORIENTATION)
        val make = exif.getAttribute(ExifInterface.TAG_MAKE)
        val model = exif.getAttribute(ExifInterface.TAG_MODEL)
        val flash = exif.getAttribute(ExifInterface.TAG_FLASH)
        val focalLength = exif.getAttribute(ExifInterface.TAG_FOCAL_LENGTH)
        return ImageDescriptor(
            uri = uri,
            lat = lat?.toDoubleOrNull(),
            lon = lon?.toDoubleOrNull(),
            shotDate = date?.parseDateTime() ?: Date(),
            width = width?.toIntOrNull() ?: 0,
            height = height?.toIntOrNull() ?: 0,
            orientation = orientation,
            caption = "",
            flash = flash,
            focalLength = focalLength,
            camera = listOfNotNull(make, model).joinToString(" - ")
        )
    }

    private fun String.parseDateTime(): Date? {
        val format = SimpleDateFormat("yyyy:MM:dd HH:mm:ss", Locale.getDefault())
        return format.parse(this)
    }
}