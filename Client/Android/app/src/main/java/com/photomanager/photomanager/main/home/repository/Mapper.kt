package com.photomanager.photomanager.main.home.repository

import android.net.Uri
import com.photomanager.photomanager.main.home.api.ImageApiDescriptor
import com.photomanager.photomanager.main.home.db.CollectionEntity
import com.photomanager.photomanager.main.home.db.FootageEntity
import com.photomanager.photomanager.main.home.model.ImageDescriptor

fun ImageApiDescriptor.toImageDescriptor(): ImageDescriptor {
    return ImageDescriptor(
        id = uniqueId,
        caption = caption,
        uri = Uri.parse(uri),
    )
}

fun ImageDescriptor.toApiDescriptor(): ImageApiDescriptor {
    return ImageApiDescriptor(
        uniqueId = id,
        caption = caption,
        uri = uri.toString(),
        date = shotDate,
    )
}

fun FootageEntity.toImageDescriptor(): ImageDescriptor = ImageDescriptor(
    id = id,
    uri = Uri.parse(uri),
    width = width,
    height = height,
    orientation = orientation,
    caption = caption,
    shotDate = date,
    lat = lat,
    lon = lon,
    camera = camera,
    focalLength = focalLength,
    flash = flash,
)

fun ImageDescriptor.toFootageEntity(): FootageEntity = FootageEntity(
    id = id,
    uri = uri.toString(),
    width = width,
    height = height,
    orientation = orientation,
    caption = caption,
    date = shotDate,
    lat = lat,
    lon = lon,
    camera = camera,
    focalLength = focalLength,
    flash = flash,
)

fun ImageDescriptor.toCollectionEntity(): CollectionEntity = CollectionEntity(
    id = id,
    uri = uri.toString(),
    width = width,
    height = height,
    orientation = orientation,
    caption = caption,
    date = shotDate,
    lat = lat,
    lon = lon,
    camera = camera,
    focalLength = focalLength,
    flash = flash,
)

fun CollectionEntity.toImageDescriptor(): ImageDescriptor = ImageDescriptor(
    id = id,
    uri = Uri.parse(uri),
    width = width,
    height = height,
    orientation = orientation,
    caption = caption,
    shotDate = date,
    lat = lat,
    lon = lon,
    camera = camera,
    focalLength = focalLength,
    flash = flash,
)
