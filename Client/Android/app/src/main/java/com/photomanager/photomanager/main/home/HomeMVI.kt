package com.photomanager.photomanager.main.home

import android.net.Uri

data class HomeState(val pickerOn: Boolean = false, val collections: List<Uri> = emptyList())
sealed class HomeEvent {
    data class OnImageClicked(val uri: Uri): HomeEvent()
    data class OnImagesPicked(val uris: List<Uri>): HomeEvent()
}

sealed class HomeAction {
    data class MarkDone(val uri: Uri): HomeAction()
}