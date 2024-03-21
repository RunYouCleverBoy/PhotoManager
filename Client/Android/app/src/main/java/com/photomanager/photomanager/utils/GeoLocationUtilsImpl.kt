package com.photomanager.photomanager.utils

import android.content.Context
import android.location.Address
import android.location.Geocoder
import android.os.Build
import dagger.hilt.android.qualifiers.ApplicationContext
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import java.util.Locale
import javax.inject.Inject
import kotlin.coroutines.resume
import kotlin.coroutines.suspendCoroutine

interface GeoLocationUtils {
    suspend fun reverseGeolocation(lat: Double, lon: Double): Address?
}

class GeoLocationUtilsImpl @Inject constructor(@ApplicationContext context: Context): GeoLocationUtils {
    private val geocoder = Geocoder(context, Locale.getDefault())
    override suspend fun reverseGeolocation(lat: Double, lon: Double): Address? {
        return getFromLatLon(geocoder, lat, lon)
    }

    private suspend fun getFromLatLon(
        geocoder: Geocoder,
        lat: Double,
        lon: Double
    ): Address? {
        return if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.TIRAMISU) {
            suspendCoroutine { cont ->
                geocoder.getFromLocation(lat, lon, 1) { response ->
                    cont.resume(response.firstOrNull())
                }
            }
        } else {
            withContext(Dispatchers.IO) {
                @Suppress("DEPRECATION")
                geocoder.getFromLocation(lat, lon, 1)?.firstOrNull()
            }
        }
    }
}