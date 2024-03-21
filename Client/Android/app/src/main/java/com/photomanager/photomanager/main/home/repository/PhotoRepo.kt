package com.photomanager.photomanager.main.home.repository

import android.net.Uri
import com.photomanager.photomanager.main.home.model.ImageDescriptor
import com.photomanager.photomanager.main.home.model.SearchCriteria
import com.photomanager.photomanager.main.home.model.WorkflowStage
import kotlinx.coroutines.flow.Flow

interface PhotoRepo {
    suspend fun getPhotosByCriteria(searchCriteria: SearchCriteria, range: IntRange): Flow<List<ImageDescriptor>>
    suspend fun importPhotos(photos: List<Uri>, footage: WorkflowStage)
    suspend fun addPhotos(photos: List<ImageDescriptor>)
    suspend fun getSize(searchCriteria: SearchCriteria): Int
    suspend fun addToCollection(ids: List<String>)
}