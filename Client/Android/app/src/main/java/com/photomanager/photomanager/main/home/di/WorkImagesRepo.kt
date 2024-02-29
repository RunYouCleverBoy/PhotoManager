package com.photomanager.photomanager.main.home.di

import android.net.Uri
import kotlinx.coroutines.flow.StateFlow

interface WorkImagesRepo {
    val images: StateFlow<List<Uri>>
    fun addImage(uri: Uri)
    fun addImages(uris: List<Uri>)
    fun removeImage(uri: Uri)
    fun getImages(): List<Uri>
    fun clearImages()
}