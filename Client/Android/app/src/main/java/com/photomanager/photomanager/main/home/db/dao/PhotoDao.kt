package com.photomanager.photomanager.main.home.db.dao

import androidx.room.Dao
import androidx.room.Delete
import androidx.room.Insert
import androidx.room.OnConflictStrategy
import androidx.room.Query
import androidx.room.Transaction
import com.photomanager.photomanager.main.home.db.AlbumWithPhotos
import com.photomanager.photomanager.main.home.db.Photo
import com.photomanager.photomanager.main.home.db.PhotoAndTags
import kotlinx.coroutines.flow.Flow

@Dao
interface PhotoDao {
    @Transaction
    @Query("SELECT * FROM photo_albums WHERE album_id = :albumId")
    fun getPhotosByAlbum(albumId: Long): Flow<List<AlbumWithPhotos>>

    @Query("SELECT * FROM photos WHERE owner = :owner")
    fun getByOwner(owner: String): Flow<List<Photo>>

    @Query(
        "SELECT * FROM photos WHERE " +
                "(shot_date BETWEEN :shotAfter AND :shotBefore) " +
                "AND (:camera IS NULL OR camera LIKE '%'||:camera||'%') " +
                "AND (:description IS NULL OR exposure LIKE '%'||:description||'%')" +
                "AND (:place IS NULL OR name LIKE '%'||:place||'%' OR city LIKE '%'||:place||'%' OR country LIKE '%'||:place||'%') " +
                "AND (:workflowStage IS NULL OR workflow_stage LIKE '%'||:workflowStage||'%') " +
                "AND (latitude BETWEEN :minLatitude AND :maxLatitude) " +
                "AND (longitude BETWEEN :minLongitude AND :maxLongitude) " +
                "AND (upvote_grade BETWEEN :gradeAtLeast AND :gradeAtMost) " +
                "ORDER BY shot_date DESC " +
                "LIMIT :limit OFFSET :offset"
    )
    suspend fun getPhotosByMetadata(
        shotBefore: Long,
        shotAfter: Long,
        camera: String?,
        description: String?,
        place: String?,
        workflowStage: String?,
        gradeAtLeast: Int,
        gradeAtMost: Int,
        minLatitude: Double,
        maxLatitude: Double,
        minLongitude: Double,
        maxLongitude: Double,
        offset: Int,
        limit: Int
    ): List<Photo>

    @Query(
        "SELECT COUNT() FROM photos WHERE " +
                "(shot_date BETWEEN :shotAfter AND :shotBefore) " +
                "AND (:camera IS NULL OR camera LIKE '%'||:camera||'%') " +
                "AND (:description IS NULL OR exposure LIKE '%'||:description||'%')" +
                "AND (:place IS NULL OR name LIKE '%'||:place||'%' OR city LIKE '%'||:place||'%' OR country LIKE '%'||:place||'%') " +
                "AND (:workflowStage IS NULL OR workflow_stage LIKE '%'||:workflowStage||'%') " +
                "AND (latitude BETWEEN :minLatitude AND :maxLatitude) " +
                "AND (longitude BETWEEN :minLongitude AND :maxLongitude) " +
                "AND (upvote_grade BETWEEN :gradeAtLeast AND :gradeAtMost) "
    )
    suspend fun getPhotosCount(
        shotBefore: Long,
        shotAfter: Long,
        camera: String?,
        description: String?,
        place: String?,
        workflowStage: String?,
        gradeAtLeast: Int,
        gradeAtMost: Int,
        minLatitude: Double,
        maxLatitude: Double,
        minLongitude: Double,
        maxLongitude: Double,
    ): Int

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun upsertPhoto(photo: Photo)

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertPhotos(photos: List<Photo>)

    @Query("DELETE FROM photos WHERE photo_id = :photoId")
    suspend fun deletePhoto(photoId: String)

    @Query("INSERT INTO photos_in_albums (album_id, photo_id) VALUES (:albumId, :photoId)")
    suspend fun addPhotoToAlbum(albumId: String, photoId: String)

    @Query("DELETE FROM photos_in_albums WHERE album_id = :albumId AND photo_id = :photoId")
    suspend fun removePhotoFromAlbum(albumId: String, photoId: String)

    @Query("SELECT tag FROM photo_tags WHERE photo_id = :photoId ORDER BY tag")
    suspend fun getTagsByPhoto(photoId: String): List<String>

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun addTagToPhoto(photoAndTags: PhotoAndTags)

    @Delete
    suspend fun removeTagFromPhoto(photoAndTags: PhotoAndTags)

    @Query("SELECT P.* FROM photos as P JOIN photo_tags as T ON T.photo_id == P.photo_id WHERE T.tag LIKE '%'||:tag||'%'")
    suspend fun getPhotosByTag(tag: String): List<Photo>

    @Query("SELECT * FROM photos WHERE photo_id IN (:ids)")
    suspend fun getPhotosByIds(ids: List<String>): List<Photo>
}