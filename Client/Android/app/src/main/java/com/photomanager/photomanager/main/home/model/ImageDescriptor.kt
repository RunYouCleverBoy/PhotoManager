package com.photomanager.photomanager.main.home.model

data class ImageDescriptor(
    val id: String,
    val url: String,
    val isPublic: Boolean,
    val owner: String,
    val visibleTo: List<String>,
    val metadata: PhotoMetadata,
    val workFlow: WorkFlow,
    val similarTo: List<String>,
    val ancestor: String,
    val comments: List<Comment>,
    val tags: List<String>
)

data class PhotoMetadata(
    val shotDate: Long?,
    val modifiedDate: Long?,
    val camera: String?,
    val place: Place?,
    val exposure: String?,
    val fNumber: Float?,
    val iso: Int?,
    val description: String?
)

data class WorkFlow(
    val upvoteGrade: Int,
    val workflowStage: WorkflowStage,
    val albums: List<String>
)

data class Comment(
    val commenterID: String,
    val commenterName: String,
    val comment: String,
    val time: Long
)

enum class WorkflowStage(val value: String) {
    FOOTAGE("footage"),
    COLLECTION("collection")
}
