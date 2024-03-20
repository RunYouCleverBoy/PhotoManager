package com.photomanager.photomanager.main.home.api.model

import com.photomanager.photomanager.main.home.api.Geolocation
import com.photomanager.photomanager.main.home.api.ObjectId
import com.photomanager.photomanager.main.home.api.Place
import com.photomanager.photomanager.main.home.api.typeconverters.WorkflowStageSerializer
import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable


@Serializable(with = WorkflowStageSerializer::class)
enum class WorkflowStage(val value: String) {
   FOOTAGE("footage"),
   COLLECTION("collection"),
   ALBUM("album")
}

@Serializable
data class WorkFlow(
   @SerialName("upvote_grade") val upvoteGrade: Int,
   @SerialName("workflow_stage") val workflowStage: WorkflowStage,
   @SerialName("albums") val albums: List<ObjectId>
)

@Serializable
data class PhotoMetadata(
   @SerialName("shot_date") val shotDate: Long?,
   @SerialName("process_date") val modifiedDate: Long?,
   @SerialName("camera") val camera: String?,
   @SerialName("location") val location: Geolocation?,
   @SerialName("place") val place: Place?,
   @SerialName("exposure") val exposure: String?,
   @SerialName("f_number") val fNumber: Float?,
   @SerialName("iso") val iso: Int?,
   @SerialName("description") val description: String?
)

@Serializable
data class Comments(
   @SerialName("commenter_id") val commenterID: ObjectId,
   @SerialName("commenter_name") val commenterName: String,
   @SerialName("comment") val comment: String,
   @SerialName("time") val time: Long
)

@Serializable
data class PhotoModel(
   @SerialName("id") val id: ObjectId,
   @SerialName("url") val url: String,
   @SerialName("is_public") val isPublic: Boolean,
   @SerialName("owner") val owner: ObjectId,
   @SerialName("visible_to") val visibleTo: List<ObjectId>,
   @SerialName("metadata") val metadata: PhotoMetadata,
   @SerialName("workflow") val workFlow: WorkFlow,
   @SerialName("similar_to") val similarTo: List<ObjectId>,
   @SerialName("ancestor") val ancestor: ObjectId,
   @SerialName("comments") val comments: List<Comments>,
   @SerialName("tags") val tags: List<String>
)

@Serializable
data class PhotoSearchLocation(
   @SerialName("geolocation") val geolocation: Geolocation,
   @SerialName("radius") val radius: Double
)

@Serializable
data class UpvoteGradeRange(
   @SerialName("min") val min: Int,
   @SerialName("max") val max: Int
)

@Serializable
data class PhotoSearchOwnedPhotoFilter(
   @SerialName("only_mine") val onlyMine: Boolean?,
   @SerialName("is_public") val isPublic: Boolean?,
   @SerialName("upvote_grade") val upvoteGrade: UpvoteGradeRange?,
   @SerialName("workflow_stage") val workflowStage: WorkflowStage?
)

@Serializable
data class PhotoSearchOptions(
   @SerialName("shot_after") val shotAfter: Long?,
   @SerialName("shot_before") val shotBefore: Long?,
   @SerialName("modified_around") val modifiedAround: Long?,
   @SerialName("camera") val camera: String?,
   @SerialName("location") val location: PhotoSearchLocation?,
   @SerialName("location_contains") val locationContains: String?,
   @SerialName("comments_containing") val commentsContaining: String?,
   @SerialName("owned_photo_filter") val ownedPhotoFilter: PhotoSearchOwnedPhotoFilter?
)

@Serializable
data class AlbumSearchCriteria(
   @SerialName("owner_id") val ownerID: ObjectId?,
   @SerialName("name") val nameRegex: String?,
   @SerialName("description") val descriptionRegex: String?,
   @SerialName("visibility_to") val visibilityTo: ObjectId?
)