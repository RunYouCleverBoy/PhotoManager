package com.photomanager.photomanager.main.nav

import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.hilt.navigation.compose.hiltViewModel
import androidx.navigation.NavHostController
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.navArgument
import com.photomanager.photomanager.main.MainAction
import com.photomanager.photomanager.main.MainEvent
import com.photomanager.photomanager.main.MainViewModel
import com.photomanager.photomanager.main.home.ui.HomeScreen
import com.photomanager.photomanager.main.photoscreen.PhotoDetailScreen
import com.photomanager.photomanager.splash.PhotoSplashScreen

@Composable
fun MainNav(navController: NavHostController, startDestination: String) {
    val mainViewModel = hiltViewModel<MainViewModel>()
    LaunchedEffect(Unit) {
        mainViewModel.action.collect { action -> navController.renderAction(action) }
    }
    NavHost(navController = navController, startDestination = startDestination, builder = {
        composable(MainNavPath.Splash.navTemplate) {
            PhotoSplashScreen { mainViewModel.dispatchEvent(MainEvent.OnSplashComplete) }
        }
        composable(MainNavPath.Home.navTemplate) {
            HomeScreen(onClicked = { uri ->
                mainViewModel.dispatchEvent(MainEvent.OnPhotoOpenRequest(uri))
            })
        }
        composable(MainNavPath.PhotoDetail.navTemplate, arguments = listOf(
            navArgument("uri") { defaultValue = "http://NoArg" }
        )) {
            val uri = it.arguments?.getString("uri") ?: "http://NULLLLLLL"
            PhotoDetailScreen(uri = uri)
        }
    })
}

private fun NavHostController.renderAction(
    action: MainAction
) {
    when (action) {
        is MainAction.NavigateTo -> navigate(action.path) {
            action.popupToPath?.let { popUpTo(it) { inclusive = true } }
        }
    }
}