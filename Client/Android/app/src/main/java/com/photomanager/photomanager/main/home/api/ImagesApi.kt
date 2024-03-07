package com.photomanager.photomanager.main.home.api

import android.net.Uri
import com.photomanager.photomanager.main.home.model.ImageDescriptor
import com.photomanager.photomanager.main.home.model.SearchCriteria

interface ImagesApi {
    suspend fun getFootage(searchCriteria: SearchCriteria, indexRange: IntRange): List<ImageDescriptor>
    suspend fun getCollection(searchCriteria: SearchCriteria, indexRange: IntRange): List<ImageDescriptor>
    suspend fun markMovedToFootage(images: List<String>): List<ImageDescriptor>
    suspend fun uploadImage(fromUri: Uri, uri: Uri)
}