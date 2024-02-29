package com.photomanager.photomanager.main

import android.net.Uri

class MainState
sealed class MainEvent {
    data class OnPhotoOpenRequest(val uri: Uri) : MainEvent()
    data object OnSplashComplete : MainEvent()
}

sealed class MainAction {
    data class NavigateTo(val path: String) : MainAction()
}