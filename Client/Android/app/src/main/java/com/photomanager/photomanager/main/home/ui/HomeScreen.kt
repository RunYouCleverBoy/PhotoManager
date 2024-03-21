package com.photomanager.photomanager.main.home.ui

import android.net.Uri
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.hilt.navigation.compose.hiltViewModel
import androidx.navigation.NavOptions
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.photomanager.photomanager.main.nav.MainNavPath

@Composable
fun HomeScreen(onClicked: (Uri) -> Unit) {
    val vm = hiltViewModel<HomeViewModel>()
    val navController = rememberNavController()

    LaunchedEffect(Unit) {
        vm.action.collect { action ->
            when (action) {
                is HomeAction.OpenImage -> onClicked(action.uri)
                is HomeAction.NavigateTo -> navController.navigate(action.path.navTemplate, navOptions = NavOptions.Builder().setPopUpTo(action.path.navTemplate, inclusive = true).build())
            }
        }
    }

    val state by vm.state.collectAsState()
    Column(modifier = Modifier.fillMaxSize(), horizontalAlignment = Alignment.CenterHorizontally) {
        NavHost(
            modifier = Modifier.weight(1f),
            navController = navController,
            startDestination = MainNavPath.HomeFootage.navTemplate
        ) {
            composable(MainNavPath.HomeFootage.navTemplate) {
                FootageScreen(state.footage, vm::dispatchEvent)
            }

            composable(MainNavPath.HomeCollection.navTemplate) {
                CollectionsScreen(state.collection, vm::dispatchEvent)
            }
        }

        HomeTabs(state.tabDescriptors, state.currentMode) {
            vm.dispatchEvent(HomeEvent.OnTabSelected(it))
        }

        LaunchedEffect(Unit) {
            vm.dispatchEvent(HomeEvent.OnTabSelected(state.currentMode))
        }
    }
}

