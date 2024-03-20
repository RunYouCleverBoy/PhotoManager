package com.photomanager.photomanager.main.home.db

import androidx.room.ColumnInfo
import androidx.room.Embedded
import androidx.room.Entity
import androidx.room.Junction
import androidx.room.PrimaryKey
import androidx.room.Relation
import androidx.room.TypeConverters

@Entity(tableName = "photos", primaryKeys = ["id"])
@TypeConverters(CompositeTypeConverters::class)
data class Photo(
    @ColumnInfo(name="id") val id: String,
    @ColumnInfo(name = "owner") val owner: String,
    @ColumnInfo(name = "url") val url: String,
    @ColumnInfo(name = "comments") val comments: Comments,
    @Embedded val metadata: PhotoMetadata,
    @Embedded val place: Place,
    @Embedded val flow: WorkFlow,
)

data class PhotoMetadata(
    @ColumnInfo(name = "shot_date") val shotDate: Long,
    @ColumnInfo(name = "modified_date") val modifiedDate: Long,
    @ColumnInfo(name = "camera") val camera: String,
    @ColumnInfo(name = "exposure") val exposure: String,
    @ColumnInfo(name = "f_number") val fNumber: Float,
    @ColumnInfo(name = "iso") val iso: Int,
    @ColumnInfo(name = "description") val description: String,
)

data class Place(
    @ColumnInfo("name") val name: String,
    @ColumnInfo("city") val city: String,
    @ColumnInfo("country") val country: String,
    @ColumnInfo("latitude") val latitude: Double,
    @ColumnInfo("longitude") val longitude: Double
)

data class WorkFlow(
    @ColumnInfo("upvote_grade") val upvoteGrade: Int,
    @ColumnInfo("workflow_stage") val workflowStage: String,
)

@Entity(tableName = "photo_albums")
data class PhotoAlbum(
    @PrimaryKey val id: String,
    @ColumnInfo(name = "cover_image_url") val coverImageUrl: String,
    @ColumnInfo(name = "name") val name: String,
    @ColumnInfo(name = "description") val description: String,
    @ColumnInfo(name = "owner") val owner: String,
    @ColumnInfo(name = "visibility") val visibility: UserVisibility,
)

@Entity(tableName = "photos_in_albums", primaryKeys = ["photo_id", "album_id"])
data class PhotosAndAlbumCross(
    @PrimaryKey(autoGenerate = true) val id: String,
    @ColumnInfo(name = "photo_id") val photoId: Int,
    @ColumnInfo(name = "album_id") val albumId: Int,
)

@Entity(tableName = "photo_tags")
data class PhotoAndTags(
    @PrimaryKey(autoGenerate = true) val id: String,
    @ColumnInfo(name = "photo_id") val photoId: Int,
    @ColumnInfo(name = "tag") val tag: String,
)

data class AlbumWithPhotos(
    @Embedded val album: PhotoAlbum,
    @Relation(
        parentColumn = "album_id",
        entityColumn = "photo_id",
        associateBy = Junction(PhotosAndAlbumCross::class)
    )
    val photos: List<Photo>,
)

data class PhotoWithTags(
    @Embedded val photo: Photo,
    @Relation(
        parentColumn = "id",
        entityColumn = "photo_id"
    )
    val tags: List<String>,
)
