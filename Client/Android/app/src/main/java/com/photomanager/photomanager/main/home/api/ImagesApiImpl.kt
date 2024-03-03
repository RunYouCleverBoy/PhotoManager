package com.photomanager.photomanager.main.home.api

import com.photomanager.photomanager.main.home.model.SearchCriteria
import io.ktor.client.HttpClient
import javax.inject.Inject

class ImagesApiImpl @Inject constructor(
    private val ktorClient: HttpClient
) : ImagesApi {
    override fun getFootage(
        searchCriteria: SearchCriteria,
        indexRange: IntRange
    ): List<ImageApiDescriptor> {
        return emptyList()
    }

    override fun getCollection(
        searchCriteria: SearchCriteria,
        indexRange: IntRange
    ): List<ImageApiDescriptor> {
        return emptyList()
    }

    override fun markMovedToFootage(images: List<String>): List<ImageApiDescriptor> {
        return emptyList()
    }
}