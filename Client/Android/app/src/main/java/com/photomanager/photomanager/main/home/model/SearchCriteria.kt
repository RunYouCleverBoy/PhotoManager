package com.photomanager.photomanager.main.home.model

import java.util.Date

open class SearchCriteria(
    val captionIncludes: String? = null,
    val afterDate: Date? = null,
    val beforeDate: Date? = null
)