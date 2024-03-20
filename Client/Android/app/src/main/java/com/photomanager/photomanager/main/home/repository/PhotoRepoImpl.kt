import com.photomanager.photomanager.main.home.api.model.PhotoModel
import com.photomanager.photomanager.main.home.db.dao.PhotoDao
import com.photomanager.photomanager.main.home.model.SearchCriteria
import com.photomanager.photomanager.main.home.repository.PhotoRepo
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import javax.inject.Inject



class PhotoRepoImpl @Inject constructor(private val photoDao: PhotoDao): PhotoRepo {

override suspend fun getPhotosByCriteria(searchCriteria: SearchCriteria): Flow<List<PhotoModel>> = flow {
    val photos = photoDao.getPhotosByMetadata(
        shotBefore = searchCriteria.dateRange.last,
        shotAfter = searchCriteria.dateRange.first,
        camera = searchCriteria.camera,
        description = searchCriteria.descriptionIncludes,
        place = searchCriteria.locationNameContains,
        workflowStage = searchCriteria.stage?.name,
        gradeAtLeast = searchCriteria.rating ?: -100,
        gradeAtMost = searchCriteria.gradeAtMost,
        minLatitude = searchCriteria.minLatitude,
        maxLatitude = searchCriteria.maxLatitude,
        minLongitude = searchCriteria.minLongitude,
        maxLongitude = searchCriteria.maxLongitude,
        offset = searchCriteria.offset,
        limit = searchCriteria.limit
    )
}
//    override suspend fun getPhotosByCriteria(searchCriteria: SearchCriteria): Flow<List<PhotoModel>> = flow {
//        val photos = photoDao.getPhotosByMetadata()
//    }
    override suspend fun addPhotos(photos: List<PhotoModel>) {
        TODO("Not yet implemented")
    }
}