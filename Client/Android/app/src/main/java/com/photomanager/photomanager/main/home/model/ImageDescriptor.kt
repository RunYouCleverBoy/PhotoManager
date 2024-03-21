package com.photomanager.photomanager.main.home.model

data class ImageDescriptor(
    val id: String,
    val url: String,
    val isPublic: Boolean = false,
    val owner: String = "",
    val visibleTo: List<String> = emptyList(),
    val metadata: PhotoMetadata = PhotoMetadata(),
    val workFlow: WorkFlow = WorkFlow(),
    val similarTo: List<String> = emptyList(),
    val ancestor: String = "",
    val comments: List<Comment> = emptyList(),
    val tags: List<String> = emptyList()
)

data class PhotoMetadata(
    val shotDate: Long? = null,
    val modifiedDate: Long? = null,
    val camera: String? = null,
    val place: Place? = null,
    val exposure: String? = null,
    val fNumber: Float? = null,
    val iso: Int? = null,
    val description: String? = null
)

data class WorkFlow(
    val upvoteGrade: Int = 0,
    val workflowStage: WorkflowStage = WorkflowStage.FOOTAGE,
    val albums: List<String> = emptyList()
)

data class Comment(
    val commentId: String,
    val commenterID: String,
    val commenterName: String,
    val comment: String,
    val time: Long
)

enum class WorkflowStage(val value: String) {
    FOOTAGE("footage"),
    COLLECTION("collection")
}
