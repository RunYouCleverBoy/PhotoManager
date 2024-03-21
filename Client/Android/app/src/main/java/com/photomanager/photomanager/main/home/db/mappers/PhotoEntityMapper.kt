package com.photomanager.photomanager.main.home.db.mappers

import com.photomanager.photomanager.main.home.db.Comments
import com.photomanager.photomanager.main.home.db.Photo
import com.photomanager.photomanager.main.home.model.Comment
import com.photomanager.photomanager.main.home.model.ImageDescriptor
import com.photomanager.photomanager.main.home.model.PhotoMetadata
import com.photomanager.photomanager.main.home.model.Place
import com.photomanager.photomanager.main.home.model.WorkFlow
import com.photomanager.photomanager.main.home.model.WorkflowStage
import com.photomanager.photomanager.main.home.db.WorkFlowStage as DbWorkFlowStage

fun Photo.toImageDescriptor(): ImageDescriptor {
    return ImageDescriptor(
        id = this.id,
        url = this.url,
        owner = this.owner,
        comments = comments.comments.map { comment ->
            Comment(
                commentId = comment.id,
                commenterID = comment.id,
                commenterName = comment.userName,
                comment = comment.text,
                time = comment.date
            )
        },
        metadata = PhotoMetadata(
            shotDate = this.metadata.shotDate,
            modifiedDate = this.metadata.modifiedDate,
            camera = this.metadata.camera,
            exposure = this.metadata.exposure,
            fNumber = this.metadata.fNumber,
            iso = this.metadata.iso,
            description = this.metadata.description,
            place = Place(
                name = this.place.name,
                city = this.place.city,
                country = this.place.country,
                latitude = this.place.latitude,
                longitude = this.place.longitude
            )
        ),
        isPublic = false,
        visibleTo = emptyList(),
        ancestor = "",
        similarTo = emptyList(),
        tags = emptyList(),
        workFlow = WorkFlow(
            upvoteGrade = this.flow.upvoteGrade,
            workflowStage = flow.workflowStage.toModel(),
            albums = emptyList()
        ),
    )
}

fun ImageDescriptor.toPhoto(): Photo {
    return Photo(
        id = this.id,
        url = this.url,
        owner = this.owner,
        comments = Comments(comments.map { comment ->
            com.photomanager.photomanager.main.home.db.Comment(
                id = comment.commentId,
                userId = comment.commenterID,
                userName = comment.commenterName,
                text = comment.comment,
                date = comment.time
            )
        }),
        metadata = com.photomanager.photomanager.main.home.db.PhotoMetadata(
            shotDate = this.metadata.shotDate ?: 0,
            modifiedDate = this.metadata.modifiedDate ?: 0,
            camera = this.metadata.camera ?: "",
            exposure = this.metadata.exposure ?: "",
            fNumber = this.metadata.fNumber ?: 0f,
            iso = this.metadata.iso ?: 0,
            description = this.metadata.description ?: ""
        ),
        place = com.photomanager.photomanager.main.home.db.Place(
            name = this.metadata.place?.name ?: "",
            city = this.metadata.place?.city ?: "",
            country = this.metadata.place?.country ?: "",
            latitude = this.metadata.place?.latitude ?: 0.0,
            longitude = this.metadata.place?.longitude ?: 0.0
        ),
        flow = com.photomanager.photomanager.main.home.db.WorkFlow(
            upvoteGrade = this.workFlow.upvoteGrade,
            workflowStage = this.workFlow.workflowStage.value
        )
    )
}

private fun String.toModel(): WorkflowStage {
    return when (this) {
        DbWorkFlowStage.FOOTAGE.value -> WorkflowStage.FOOTAGE
        DbWorkFlowStage.COLLECTION.value -> WorkflowStage.COLLECTION
        else -> WorkflowStage.FOOTAGE
    }
}
