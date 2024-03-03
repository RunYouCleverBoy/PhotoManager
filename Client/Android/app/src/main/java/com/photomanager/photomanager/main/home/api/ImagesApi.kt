package com.photomanager.photomanager.main.home.api

import com.photomanager.photomanager.main.home.model.SearchCriteria

interface ImagesApi {
    fun getFootage(searchCriteria: SearchCriteria, indexRange: IntRange): List<ImageApiDescriptor>
    fun getCollection(searchCriteria: SearchCriteria, indexRange: IntRange): List<ImageApiDescriptor>
    fun markMovedToFootage(images: List<String>): List<ImageApiDescriptor>
}