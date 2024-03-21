package com.photomanager.photomanager.main.home.api

import android.net.Uri
import com.photomanager.photomanager.main.home.api.mappers.toPhotoSearchOptions
import com.photomanager.photomanager.main.home.ktor.KtorFactory
import com.photomanager.photomanager.main.home.model.ImageDescriptor
import com.photomanager.photomanager.main.home.model.SearchCriteria
import io.ktor.client.HttpClient
import io.ktor.client.call.body
import io.ktor.client.request.post
import io.ktor.client.request.setBody
import io.ktor.http.isSuccess
import javax.inject.Inject

class ImagesApiImpl @Inject constructor(
    private val client: HttpClient, private val config: KtorFactory.Configuration
) : ImagesApi {
    override suspend fun search(
        searchCriteria: SearchCriteria, indexRange: IntRange
    ): List<ImageDescriptor> {
        val urlString = Uri.parse(config.baseUrl)
            .buildUpon()
            .appendPath("photos")
            .build().toString()
        val result = client.post(urlString) {
            setBody(searchCriteria.toPhotoSearchOptions())
        }

        return if (result.status.isSuccess()) {
            result.body() ?: emptyList()
        } else {
            emptyList()
        }
    }

    override suspend fun uploadImage(fromUri: Uri, uri: Uri) {

    }
}
