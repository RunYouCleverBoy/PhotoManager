package com.photomanager.photomanager.main.nav

sealed class MainNavPath(val navTemplate: String) {
    data object Splash : MainNavPath("splash")
    data object Home : MainNavPath("home")
    data object PhotoDetail : MainNavPath("photo_detail?uri={uri}")
}