package com.photomanager.photomanager.main.home.repository

import android.net.Uri
import com.photomanager.photomanager.main.home.db.DatabaseHolder
import com.photomanager.photomanager.main.home.db.FootageEntity
import com.photomanager.photomanager.main.home.model.ImageDescriptor
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.async
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import kotlinx.coroutines.launch
import javax.inject.Inject

class ImagesRepoImpl @Inject constructor(private val databaseHolder: DatabaseHolder) : ImagesRepo {
    override fun getFootage(
        searchCriteria: ImagesRepo.SearchCriteria,
        indexRange: IntRange
    ): Flow<List<ImageDescriptor>> = flow {
        val dbResult = CoroutineScope(Dispatchers.IO).async {
            issueCallForDb(searchCriteria, indexRange)
        }
        val netResult = CoroutineScope(Dispatchers.IO).async {
            issueCallForNet(searchCriteria, indexRange)
        }
        val dbList = dbResult.await().map { it.toImageDescriptor() }
        emit(dbList)
        val netList = netResult.await()
        if (netList.isNotEmpty()) {
            emit(netList)
            CoroutineScope(Dispatchers.IO).launch {
                netList.forEach {
                    databaseHolder.database.footageDao().insertFootage(it.toFootageEntity())
                }
            }
        }
    }
    private suspend fun issueCallForDb(
        searchCriteria: ImagesRepo.SearchCriteria,
        indexRange: IntRange
    ): List<FootageEntity> {
        return databaseHolder.database.footageDao().getFootageBy(
            afterDate = searchCriteria.afterDate,
            beforeDate = searchCriteria.beforeDate,
            captionIncludes = searchCriteria.captionIncludes,
            startIndex = indexRange.first,
            size = if (!indexRange.isEmpty()) indexRange.last - indexRange.first + 1 else 0
        )
    }

    private fun issueCallForNet(searchCriteria: ImagesRepo.SearchCriteria, indexRange: IntRange): List<ImageDescriptor> {
        return emptyList() // TODO: Stub
    }

    override fun getCollection(
        searchCriteria: ImagesRepo.SearchCriteria,
        indexRange: IntRange
    ): List<ImageDescriptor> {
        return emptyList() // TODO: Stub
    }

    override fun addToFootage(imageDescriptor: List<ImageDescriptor>) {
        TODO("Not yet implemented")
    }

    override fun addToCollection(imageDescriptor: List<ImageDescriptor>) {
        TODO("Not yet implemented")
    }

    private fun FootageEntity.toImageDescriptor(): ImageDescriptor {
        return ImageDescriptor(
            uri = Uri.parse(uri),
            caption = caption,
            shotDate = date,
            orientation = orientation,
            lat = lat,
            lon = lon,
            width = width,
            height = height,
            camera = camera,
            focalLength = focalLength,

            )
    }

    private fun ImageDescriptor.toFootageEntity(): FootageEntity {
        return FootageEntity(
            uri = uri.toString(),
            caption = caption,
            date = shotDate,
            orientation = orientation,
            lat = lat,
            lon = lon,
            width = width,
            height = height,
            camera = camera,
            focalLength = focalLength,
        )
    }
}