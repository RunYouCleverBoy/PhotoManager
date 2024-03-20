package com.photomanager.photomanager.main.home.repository

import com.photomanager.photomanager.main.home.api.model.PhotoModel
import com.photomanager.photomanager.main.home.model.SearchCriteria
import kotlinx.coroutines.flow.Flow

interface PhotoRepo {
    suspend fun getPhotosByCriteria(searchCriteria: SearchCriteria): Flow<List<PhotoModel>>
    suspend fun addPhotos(photos: List<PhotoModel>)
}