package com.photomanager.photomanager.main.home.repository

import com.photomanager.photomanager.main.home.model.ImageDescriptor
import kotlinx.coroutines.flow.Flow
import java.util.Date

interface ImagesRepo {
    open class SearchCriteria(
        val captionIncludes: String? = null,
        val afterDate: Date? = null,
        val beforeDate: Date? = null
    )

    fun getFootage(
        searchCriteria: SearchCriteria, indexRange: IntRange
    ): Flow<List<ImageDescriptor>>

    fun getCollection(
        searchCriteria: SearchCriteria,
        indexRange: IntRange
    ): List<ImageDescriptor>

    fun addToFootage(imageDescriptor: List<ImageDescriptor>)
    fun addToCollection(imageDescriptor: List<ImageDescriptor>)
}