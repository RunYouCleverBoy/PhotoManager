package com.photomanager.photomanager.main.home.repository

import android.net.Uri
import com.photomanager.photomanager.main.home.model.ImageDescriptor
import com.photomanager.photomanager.main.home.model.WorkflowStage

interface ImageProcessorRepo {
    suspend fun processExif(uri: Uri, asWorkflowStage: WorkflowStage): ImageDescriptor
    fun idForPhoto(photoId: String?, photoUri: String?, collection: WorkflowStage): String
}