package com.photomanager.photomanager.main.home.api

import android.net.Uri
import com.photomanager.photomanager.main.home.model.ImageDescriptor
import com.photomanager.photomanager.main.home.model.SearchCriteria

interface ImagesApi {
    suspend fun search(searchCriteria: SearchCriteria, indexRange: IntRange = 0..1000): List<ImageDescriptor>
    suspend fun uploadImage(fromUri: Uri, uri: Uri)
}