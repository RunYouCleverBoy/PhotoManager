package com.photomanager.photomanager.main.home

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
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.photomanager.photomanager.main.nav.MainNavPath

@Composable
fun HomeScreen(onClicked: (Uri) -> Unit) {
    val vm = hiltViewModel<HomeViewModel>()

    LaunchedEffect(Unit) {
        vm.action.collect {
            when (it) {
                is HomeAction.OpenImage -> onClicked(it.uri)
            }
        }
    }

    val state by vm.state.collectAsState()
    Column(modifier = Modifier.fillMaxSize(), horizontalAlignment = Alignment.CenterHorizontally) {
        NavHost(navController = rememberNavController(), startDestination = MainNavPath.HomeFootage.navTemplate) {
            composable(MainNavPath.HomeFootage.navTemplate) {
                FootageScreen(state.footage, vm::dispatchEvent)
            }

            composable(MainNavPath.HomeCollection.navTemplate) {
                FootageScreen(state.collection, vm::dispatchEvent)
            }
        }

    }
}

