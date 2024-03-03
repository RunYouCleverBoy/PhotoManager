package com.photomanager.photomanager.main.home

import android.net.Uri
import com.photomanager.photomanager.main.home.repository.ImagesRepo

data class ImageUIDescriptor(val uri: Uri, val caption: String)

data class HomeState(
    val pickerOn: Boolean = false,
    val footage: List<ImageUIDescriptor> = emptyList(),
    val collection: List<ImageUIDescriptor> = emptyList(),
    val footageSearchCriteria: ImagesRepo.SearchCriteria = ImagesRepo.SearchCriteria(),
    val collectionSearchCriteria: ImagesRepo.SearchCriteria = ImagesRepo.SearchCriteria()
)

sealed class HomeEvent {
    data class OnImageClicked(val uri: Uri) : HomeEvent()
    data class OnImagesPicked(val uris: List<Uri>) : HomeEvent()
    data class OnApproachingFootageWindowEnd(val index: Int) : HomeEvent()
    data class OnApproachingCollectionWindowEnd(val index: Int) : HomeEvent()
}

sealed class HomeAction {
    data class OpenImage(val uri: Uri) : HomeAction()
}