package com.photomanager.photomanager.main.home.api

import android.net.Uri
import com.photomanager.photomanager.main.home.model.ImageDescriptor
import com.photomanager.photomanager.main.home.model.SearchCriteria

interface ImagesApi {
    suspend fun search(searchCriteria: SearchCriteria, indexRange: IntRange): List<ImageDescriptor>
    suspend fun uploadImage(fromUri: Uri, uri: Uri)
}