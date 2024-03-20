package com.photomanager.photomanager.main.home.api.mappers

import android.location.Location
import com.photomanager.photomanager.main.home.api.Geolocation
import com.photomanager.photomanager.main.home.api.model.PhotoSearchLocation
import com.photomanager.photomanager.main.home.api.model.PhotoSearchOptions
import com.photomanager.photomanager.main.home.api.model.PhotoSearchOwnedPhotoFilter
import com.photomanager.photomanager.main.home.api.model.UpvoteGradeRange
import com.photomanager.photomanager.main.home.api.model.WorkflowStage
import com.photomanager.photomanager.main.home.model.SearchCriteria
import com.photomanager.photomanager.main.home.model.WorkflowStage as ModelWorkFlowStage

fun geoRectToPolar(
    east: Double,
    north: Double,
    west: Double,
    south: Double
): Pair<Geolocation, Double> {
    val southWest = Location("SearchCriteria").apply {
        latitude = south
        longitude = west
    }
    val northEast = Location("SearchCriteria").apply {
        latitude = north
        longitude = east
    }

    val radius = southWest.distanceTo(northEast) / 2.0
    return Geolocation((south + north) / 2.0, (east + west) / 2.0) to radius
}

fun ModelWorkFlowStage.workFlowMapper(): WorkflowStage = when (this) {
    ModelWorkFlowStage.FOOTAGE -> WorkflowStage.FOOTAGE
    ModelWorkFlowStage.COLLECTION -> WorkflowStage.COLLECTION
}

fun IntRange.toRatingRange(): UpvoteGradeRange = UpvoteGradeRange(start, endInclusive)

fun SearchCriteria.toPhotoSearchOptions(): PhotoSearchOptions {
    val (geoLocation, radius) = geoRectToPolar(
        east = longitudeRange.endInclusive,
        north = latitudeRange.endInclusive,
        west = longitudeRange.start,
        south = latitudeRange.start
    )
    return PhotoSearchOptions(
        shotAfter = dateRange.first,
        shotBefore = dateRange.last,
        camera = camera,
        location = PhotoSearchLocation(geoLocation, radius),
        locationContains = locationNameContains,
        commentsContaining = commentsContaining,
        ownedPhotoFilter = PhotoSearchOwnedPhotoFilter(
            onlyMine = null, // No corresponding field in SearchCriteria
            isPublic = null, // No corresponding field in SearchCriteria
            upvoteGrade = rating?.takeUnless { it.isEmpty() }?.toRatingRange(),
            workflowStage = stage?.workFlowMapper()
        ),
        modifiedAround = null
    )
}