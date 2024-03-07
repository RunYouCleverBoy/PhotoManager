package com.photomanager.photomanager.main.home.api.mappers

import android.net.Uri
import com.photomanager.photomanager.main.home.model.ImageDescriptor
import models.PhotoMetadata
import models.PhotoModel
import models.WorkFlow
import models.WorkflowStage
import java.util.Date

fun PhotoModel.toImageDescriptor(): ImageDescriptor = ImageDescriptor(
    id = id,
    uri = Uri.parse(url),
    width = 0, // width is not available in PhotoModel
    height = 0, // height is not available in PhotoModel
    orientation = null, // orientation is not available in PhotoModel
    caption = metadata.description ?: "",
    shotDate = Date(metadata.shotDate ?: 0),
    lat = metadata.location?.latitude,
    lon = metadata.location?.longitude,
    camera = metadata.camera ?: "",
    focalLength = null, // focalLength is not available in PhotoModel
    flash = null, // flash is not available in PhotoModel
)

fun ImageDescriptor.toPhotoModel(): PhotoModel = PhotoModel(
    id = id,
    url = uri.toString(),
    isPublic = true, // isPublic is not available in ImageDescriptor
    owner = "", // owner is not available in ImageDescriptor
    visibleTo = emptyList(), // visibleTo is not available in ImageDescriptor
    metadata = PhotoMetadata(
        shotDate = shotDate.time,
        modifiedDate = null, // modifiedDate is not available in ImageDescriptor
        camera = camera,
        location = null, // location is not available in ImageDescriptor
        place = null, // place is not available in ImageDescriptor
        exposure = null, // exposure is not available in ImageDescriptor
        fNumber = null, // fNumber is not available in ImageDescriptor
        iso = null, // iso is not available in ImageDescriptor
        description = caption
    ),
    workFlow = WorkFlow(0, workflowStage = WorkflowStage.FOOTAGE, albums = emptyList()), // workFlow is not available in ImageDescriptor
    similarTo = emptyList(), // similarTo is not available in ImageDescriptor
    ancestor = "", // ancestor is not available in ImageDescriptor
    comments = emptyList(), // comments is not available in ImageDescriptor
    tags = emptyList() // tags is not available in ImageDescriptor
)