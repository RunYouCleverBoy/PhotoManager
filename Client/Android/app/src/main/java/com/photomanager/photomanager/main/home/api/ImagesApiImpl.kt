package com.photomanager.photomanager.main.home.api

import android.net.Uri
import com.photomanager.photomanager.main.home.ktor.KtorFactory
import com.photomanager.photomanager.main.home.model.SearchCriteria
import io.ktor.client.HttpClient
import io.ktor.client.call.body
import io.ktor.client.request.get
import javax.inject.Inject

class ImagesApiImpl @Inject constructor(
    private val client: HttpClient, private val config: KtorFactory.Configuration
) : ImagesApi {
    override suspend fun getFootage(
        searchCriteria: SearchCriteria, indexRange: IntRange
    ): List<ImageApiDescriptor> {
        val urlString = Uri.parse(config.baseUrl)
            .buildUpon()
            .appendPath("photos")
            .build().toString()
        val result = client.get(urlString).body<String>()
        return emptyList()
    }

    override suspend fun getCollection(
        searchCriteria: SearchCriteria, indexRange: IntRange
    ): List<ImageApiDescriptor> {
        return emptyList()
    }

    override suspend fun markMovedToFootage(images: List<String>): List<ImageApiDescriptor> {
        return emptyList()
    }

    override suspend fun uploadImage(uri: Uri) {

    }
}