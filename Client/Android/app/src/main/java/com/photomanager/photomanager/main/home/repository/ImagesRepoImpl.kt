package com.photomanager.photomanager.main.home.repository

import com.photomanager.photomanager.main.home.api.ImageApiDescriptor
import com.photomanager.photomanager.main.home.api.ImagesApi
import com.photomanager.photomanager.main.home.db.CollectionDao
import com.photomanager.photomanager.main.home.db.FootageDao
import com.photomanager.photomanager.main.home.model.ImageDescriptor
import com.photomanager.photomanager.main.home.model.SearchCriteria
import com.photomanager.photomanager.utils.size
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.async
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import kotlinx.coroutines.launch
import javax.inject.Inject

class ImagesRepoImpl @Inject constructor(
    api: ImagesApi,
    footageDao: FootageDao,
    collectionDao: CollectionDao
) : ImagesRepo {
    override val footage: ImagesRepo.Channel.Footage = FootageRepoImpl(footageDao, api)
    override val collection: ImagesRepo.Channel.Collection = CollectionRepoImpl(collectionDao, footageDao, api)
}

class FootageRepoImpl(private val dao: FootageDao, private val api: ImagesApi) :
    ImagesRepo.Channel.Footage() {
    override suspend fun add(imageDescriptor: List<ImageDescriptor>) {
        dao.insertFootage(imageDescriptor.map { it.toFootageEntity() })
    }

    override suspend fun get(
        searchCriteria: SearchCriteria,
        indexRange: IntRange
    ): Flow<List<ImageDescriptor>> {
        val scope = CoroutineScope(Dispatchers.IO)
        return flow {
            val footage = scope.async {
                // get data from server
                api.getFootage(searchCriteria, indexRange).map { it.toImageDescriptor() }
            }
            var fromDb = footageEntities(searchCriteria, indexRange)
            emit(fromDb)

            dao.insertFootage(footage.await().map { it.toFootageEntity() })

            fromDb = footageEntities(searchCriteria, indexRange)
            emit(fromDb)
        }
    }

    override suspend fun getSize(searchCriteria: SearchCriteria): Int {
        return dao.countFootageBy(
            searchCriteria.afterDate,
            searchCriteria.beforeDate,
            searchCriteria.captionIncludes
        )
    }

    private suspend fun footageEntities(
        searchCriteria: SearchCriteria,
        indexRange: IntRange
    ) = dao.getFootageBy(
        afterDate = searchCriteria.afterDate,
        beforeDate = searchCriteria.beforeDate,
        captionIncludes = searchCriteria.captionIncludes,
        startIndex = indexRange.first,
        size = indexRange.size
    ).map { it.toImageDescriptor() }
}

class CollectionRepoImpl(
    private val dao: CollectionDao,
    private val footageDao: FootageDao,
    private val api: ImagesApi
) : ImagesRepo.Channel.Collection() {
    override suspend fun add(ids: List<String>) {
        val scope = CoroutineScope(Dispatchers.IO)
        scope.launch {
            val imageDescriptors =
                footageDao.getFootageByIds(ids).map { it.toImageDescriptor() }
            val collection: List<ImageApiDescriptor> =
                api.markMovedToFootage(imageDescriptors.map { it.id })
            dao.insertCollection(collection.map {
                it.toImageDescriptor().toCollectionEntity()
            })
        }
    }

    override suspend fun get(
        searchCriteria: SearchCriteria,
        indexRange: IntRange
    ): Flow<List<ImageDescriptor>> {
        val scope = CoroutineScope(Dispatchers.IO)
        return flow {
            val footage = scope.async {
                // get data from server
                api.getCollection(searchCriteria, indexRange).map { it.toImageDescriptor() }
            }
            var fromDb = collectionEntities(searchCriteria, indexRange)
            emit(fromDb)

            footageDao.insertFootage(footage.await().map { it.toFootageEntity() })

            fromDb = collectionEntities(searchCriteria, indexRange)
            emit(fromDb)
        }
    }

    override suspend fun getSize(searchCriteria: SearchCriteria): Int {
        return dao.countCollectionBy(
            searchCriteria.afterDate,
            searchCriteria.beforeDate,
            searchCriteria.captionIncludes
        )
    }

    private suspend fun collectionEntities(
        searchCriteria: SearchCriteria,
        indexRange: IntRange
    ) = dao.getCollectionBy(
        afterDate = searchCriteria.afterDate,
        beforeDate = searchCriteria.beforeDate,
        captionIncludes = searchCriteria.captionIncludes,
        startIndex = indexRange.first,
        size = indexRange.size
    ).map { it.toImageDescriptor() }
}
