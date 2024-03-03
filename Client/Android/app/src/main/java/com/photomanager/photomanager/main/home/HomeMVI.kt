package com.photomanager.photomanager.main.home

import android.net.Uri
import com.photomanager.photomanager.main.home.model.SearchCriteria

data class ImageUIDescriptor(val id: String, val uri: Uri, val caption: String)

data class HomeState(
    val isBusy: Boolean = false,
    val pickerOn: Boolean = false,
    val footage: List<ImageUIDescriptor> = emptyList(),
    val collection: List<ImageUIDescriptor> = emptyList(),
    val footageSearchCriteria: SearchCriteria = SearchCriteria(),
    val collectionSearchCriteria: SearchCriteria = SearchCriteria()
)

sealed class HomeEvent {
    data class OnImageClicked(val uri: Uri) : HomeEvent()
    data class OnImagesPicked(val uris: List<Uri>) : HomeEvent()
    data class OnAddToCollection(val ids: List<String>) : HomeEvent()
    data object OnApproachingFootageWindowEnd : HomeEvent()
    data object OnApproachingCollectionWindowEnd : HomeEvent()
}

sealed class HomeAction {
    data class OpenImage(val uri: Uri) : HomeAction()
}