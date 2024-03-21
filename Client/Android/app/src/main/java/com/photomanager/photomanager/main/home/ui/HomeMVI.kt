package com.photomanager.photomanager.main.home.ui

import android.net.Uri
import com.photomanager.photomanager.main.home.LazyBulk
import com.photomanager.photomanager.main.home.model.SearchCriteria
import com.photomanager.photomanager.main.home.model.WorkflowStage

sealed class ImageUIDescriptor {
    data class Data(val id: String, val uri: Uri, val caption: String) : ImageUIDescriptor()
    data object Loading : ImageUIDescriptor()
}

data class HomeState(
    val isBusy: Boolean = false,
    val pickerOn: Boolean = false,
    val tabDescriptors: List<HomeTabRepo.TabDescriptor> = emptyList(),
    val currentMode: WorkflowStage = WorkflowStage.FOOTAGE,
    val footage: LazyBulk<ImageUIDescriptor> = LazyBulk(cachedSize = 1000) { ImageUIDescriptor.Loading },
    val collection: LazyBulk<ImageUIDescriptor> = LazyBulk(cachedSize = 1000) { ImageUIDescriptor.Loading },
    val footageSearchCriteria: SearchCriteria = SearchCriteria(stage = WorkflowStage.FOOTAGE),
    val collectionSearchCriteria: SearchCriteria = SearchCriteria(stage = WorkflowStage.COLLECTION)
)

sealed class HomeEvent {
    data class OnImageClicked(val uri: Uri) : HomeEvent()
    data class OnImagesPicked(val uris: List<Uri>) : HomeEvent()
    data class OnAddToCollection(val ids: List<String>) : HomeEvent()
    data class OnTabSelected(val stage: WorkflowStage) : HomeEvent()
    data class OnFootageAboutToMiss(val index: Int) : HomeEvent()
    data class OnCollectionAboutToMiss(val index: Int) : HomeEvent()
}

sealed class HomeAction {
    data class OpenImage(val uri: Uri) : HomeAction()
}