package com.photomanager.photomanager.main.home.repository

import com.photomanager.photomanager.main.home.model.ImageDescriptor
import com.photomanager.photomanager.main.home.model.SearchCriteria
import kotlinx.coroutines.flow.Flow

interface ImagesRepo {
    sealed interface Channel {
        suspend fun get(searchCriteria: SearchCriteria, indexRange: IntRange): Flow<List<ImageDescriptor>>
        suspend fun getSize(searchCriteria: SearchCriteria): Int

        abstract class Footage : Channel {
            abstract suspend fun add(imageDescriptor: List<ImageDescriptor>)
        }
        abstract class Collection : Channel {
            abstract suspend fun add(ids: List<String>)
        }
    }

    val footage: Channel.Footage
    val collection: Channel.Collection
}