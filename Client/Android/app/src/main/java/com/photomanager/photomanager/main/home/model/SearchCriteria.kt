package com.photomanager.photomanager.main.home.model

open class SearchCriteria(
    val descriptionIncludes: String? = null,
    val dateRange: LongRange = LongRange(0, Long.MAX_VALUE),
    val locationNameContains: String? = null,
    val latitudeRange: ClosedRange<Double> = -90.0..90.0,
    val longitudeRange: ClosedRange<Double> = -180.0..180.0,
    val camera: String? = null,
    val commentsContaining: String? = null,
    val rating: IntRange? = null,
    val tag: String? = null,
    val stage: WorkflowStage? = null,
)