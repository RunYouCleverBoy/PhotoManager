package com.photomanager.photomanager.main.home.repository

import android.net.Uri
import com.photomanager.photomanager.main.home.api.PhotosApi
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.update
import javax.inject.Inject

class WorkImagesRepoImpl @Inject constructor(val api: PhotosApi) : WorkImagesRepo {
    private val mutableImages: MutableStateFlow<List<Uri>> = MutableStateFlow(emptyList())
    override val images: StateFlow<List<Uri>> = mutableImages

    override fun addImage(uri: Uri) {
        if (uri in mutableImages.value) return
        mutableImages.update { images -> images + uri }
    }

    override fun addImages(uris: List<Uri>) {
        val images = mutableImages.value
        val alreadyThere = images.toSet()
        mutableImages.value = images + uris.filter { it !in alreadyThere }
    }
    override fun removeImage(uri: Uri) {
        mutableImages.update { images -> images - uri }
    }

    override fun getImages(): List<Uri> = mutableImages.value

    override fun clearImages() { mutableImages.value = emptyList() }
}