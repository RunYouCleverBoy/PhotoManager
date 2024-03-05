package com.photomanager.photomanager.main.home.api

import android.net.Uri
import com.photomanager.photomanager.main.home.model.SearchCriteria

interface ImagesApi {
    suspend fun getFootage(searchCriteria: SearchCriteria, indexRange: IntRange): List<ImageApiDescriptor>
    suspend fun getCollection(searchCriteria: SearchCriteria, indexRange: IntRange): List<ImageApiDescriptor>
    suspend fun markMovedToFootage(images: List<String>): List<ImageApiDescriptor>
    suspend fun uploadImage(uri: Uri)
}