package com.photomanager.photomanager.main.home.repository

import android.net.Uri
import com.photomanager.photomanager.main.home.api.ImagesApi
import com.photomanager.photomanager.main.home.db.dao.PhotoDao
import com.photomanager.photomanager.main.home.db.mappers.toImageDescriptor
import com.photomanager.photomanager.main.home.db.mappers.toPhoto
import com.photomanager.photomanager.main.home.model.ImageDescriptor
import com.photomanager.photomanager.main.home.model.SearchCriteria
import com.photomanager.photomanager.main.home.model.WorkflowStage
import com.photomanager.photomanager.utils.size
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import kotlinx.coroutines.withContext
import javax.inject.Inject
import com.photomanager.photomanager.main.home.db.WorkFlowStage as DbWorkFlowStage


class PhotoRepoImpl @Inject constructor(
    private val photoDao: PhotoDao,
    private val photoApi: ImagesApi,
    private val imageProcessorRepo: ImageProcessorRepo
) : PhotoRepo {
    override suspend fun getPhotosByCriteria(
        searchCriteria: SearchCriteria,
        range: IntRange
    ): Flow<List<ImageDescriptor>> =
        flow {
            val photos = queryDbBySearchCriteria(searchCriteria, range)
            emit(photos)
            val serverData = photoApi.search(searchCriteria)
            val serverPhotos = serverData.map{it.toPhoto()}
            photoDao.insertPhotos(serverPhotos)
            val photosWithServer = queryDbBySearchCriteria(searchCriteria, range)
            emit(photosWithServer)
        }

    private suspend fun queryDbBySearchCriteria(
        searchCriteria: SearchCriteria,
        range: IntRange
    ) = photoDao.getPhotosByMetadata(
        shotBefore = searchCriteria.dateRange.last,
        shotAfter = searchCriteria.dateRange.first,
        camera = searchCriteria.camera,
        description = searchCriteria.descriptionIncludes,
        place = searchCriteria.locationNameContains,
        workflowStage = searchCriteria.stage?.name,
        gradeAtLeast = searchCriteria.rating?.first ?: -1000,
        gradeAtMost = searchCriteria.rating?.last ?: 1000,
        minLatitude = searchCriteria.latitudeRange.start,
        maxLatitude = searchCriteria.latitudeRange.endInclusive,
        minLongitude = searchCriteria.longitudeRange.start,
        maxLongitude = searchCriteria.longitudeRange.endInclusive,
        offset = range.first,
        limit = range.size
    ).map { it.toImageDescriptor() }

    override suspend fun importPhotos(photos: List<Uri>, footage: WorkflowStage) {
        val imageDescriptors = withContext(Dispatchers.IO) {
            photos.map { imageProcessorRepo.processExif(it, WorkflowStage.FOOTAGE) }
        }
        photoDao.insertPhotos(imageDescriptors.map { it.toPhoto() })
    }

    override suspend fun addPhotos(photos: List<ImageDescriptor>) {
        photoDao.insertPhotos(photos.map { it.toPhoto() })
    }

    override suspend fun getSize(searchCriteria: SearchCriteria): Int {
        return photoDao.getPhotosCount(
            shotBefore = searchCriteria.dateRange.last,
            shotAfter = searchCriteria.dateRange.first,
            camera = searchCriteria.camera,
            description = searchCriteria.descriptionIncludes,
            place = searchCriteria.locationNameContains,
            workflowStage = searchCriteria.stage?.name,
            gradeAtLeast = searchCriteria.rating?.first ?: -1000,
            gradeAtMost = searchCriteria.rating?.last ?: 1000,
            minLatitude = searchCriteria.latitudeRange.start,
            maxLatitude = searchCriteria.latitudeRange.endInclusive,
            minLongitude = searchCriteria.longitudeRange.start,
            maxLongitude = searchCriteria.longitudeRange.endInclusive,
        )
    }

    override suspend fun addToCollection(ids: List<String>) {
        val photos = photoDao.getPhotosByIds(ids)
        val convertedPhotos = photos.map { photo ->
            photo.copy(
                id = imageProcessorRepo.idForPhoto(photo.id, photo.url, WorkflowStage.COLLECTION),
                flow = photo.flow.copy(workflowStage = DbWorkFlowStage.COLLECTION.value)
            )
        }
        photoDao.insertPhotos(convertedPhotos)
    }
}