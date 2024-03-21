package com.photomanager.photomanager.main.home.repository

import android.content.Context
import android.net.Uri
import androidx.exifinterface.media.ExifInterface
import com.photomanager.photomanager.main.home.db.Photo
import com.photomanager.photomanager.main.home.model.ImageDescriptor
import com.photomanager.photomanager.main.home.model.PhotoMetadata
import com.photomanager.photomanager.main.home.model.Place
import com.photomanager.photomanager.main.home.model.WorkFlow
import com.photomanager.photomanager.main.home.model.WorkflowStage
import com.photomanager.photomanager.utils.GeoLocationUtils
import dagger.hilt.android.qualifiers.ApplicationContext
import java.text.SimpleDateFormat
import java.util.Date
import java.util.Locale
import java.util.UUID
import javax.inject.Inject

class ImageProcessorRepoImpl @Inject constructor(
    @ApplicationContext val context: Context,
    private val geoLocationUtils: GeoLocationUtils,
) : ImageProcessorRepo {
    override suspend fun processExif(uri: Uri, asWorkflowStage: WorkflowStage): ImageDescriptor {
        val exif = context.contentResolver.openInputStream(uri).use { inputStream ->
            inputStream?.let { ExifInterface(it) }
        }

        val lat = exif?.getAttribute(ExifInterface.TAG_GPS_LATITUDE)?.toDoubleOrNull()
        val lon = exif?.getAttribute(ExifInterface.TAG_GPS_LONGITUDE)?.toDoubleOrNull()
        val date = exif?.getAttribute(ExifInterface.TAG_DATETIME)
        val exposure = exif?.getAttribute(ExifInterface.TAG_EXPOSURE_TIME)
        val fNumber = exif?.getAttribute(ExifInterface.TAG_F_NUMBER)?.toFloatOrNull()
        val make = exif?.getAttribute(ExifInterface.TAG_MAKE)
        val model = exif?.getAttribute(ExifInterface.TAG_MODEL)

        val address = if (lat != null && lon != null) {
            geoLocationUtils.reverseGeolocation(lat, lon)
        } else {
            null
        }

        val place = if (address != null) {
            Place(
                country = address.countryName,
                city = address.locality,
                name = address.featureName,
                latitude = lat ?: 0.0,
                longitude = lon ?: 0.0,
            )
        } else {
            null
        }

        return ImageDescriptor(
            id = idForPhoto(null, asWorkflowStage),
            url = uri.toString(),
            metadata = PhotoMetadata(
                shotDate = (date?.parseDateTime() ?: Date()).time,
                modifiedDate = System.currentTimeMillis(),
                camera = "$make ; $model",
                exposure = exposure,
                fNumber = fNumber,
                place = place
            ),
            workFlow = WorkFlow(
                upvoteGrade = 0,
                workflowStage = asWorkflowStage,
                albums = emptyList()
            ),
            isPublic = false,
            owner = "",
        )
    }

    override fun idForPhoto(photo: Photo?, collection: WorkflowStage): String {
        val id = when {
            photo == null -> UUID.randomUUID().toString()
            photo.id.length >= 10 -> photo.id.substring(1)
            photo.url.isNotBlank() -> UUID.fromString(photo.url).toString()
            else -> UUID.randomUUID().toString()
        }
        return when (collection) {
            WorkflowStage.FOOTAGE -> "f$id"
            WorkflowStage.COLLECTION -> "c$id"
        }
    }

    private fun String.parseDateTime(): Date? {
        val format = SimpleDateFormat("yyyy:MM:dd HH:mm:ss", Locale.getDefault())
        return try {
            format.parse(this)
        } catch (e: Exception) {
            null
        }
    }
}