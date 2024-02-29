package com.photomanager.photomanager.main

import com.photomanager.photomanager.main.nav.MainNavPath
import com.photomanager.photomanager.mvi.MVIViewModel
import dagger.hilt.android.lifecycle.HiltViewModel
import javax.inject.Inject

@HiltViewModel
class MainViewModel @Inject constructor() :
    MVIViewModel<MainState, MainEvent, MainAction>(MainState()) {
    override fun dispatchEvent(event: MainEvent) {
        when (event) {
            MainEvent.OnSplashComplete -> {
                emit(MainAction.NavigateTo(MainNavPath.Home.navTemplate, MainNavPath.Splash.navTemplate))
            }

            is MainEvent.OnPhotoOpenRequest -> {
                val path = MainNavPath.PhotoDetail.navTemplate.replace("{uri}", event.uri.toString())
                emit(MainAction.NavigateTo(path))
            }
        }
    }
}

