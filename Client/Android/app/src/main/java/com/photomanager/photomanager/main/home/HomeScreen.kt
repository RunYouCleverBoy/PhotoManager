package com.photomanager.photomanager.main.home

import android.net.Uri
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.lazy.grid.GridCells
import androidx.compose.foundation.lazy.grid.LazyVerticalGrid
import androidx.compose.material3.ElevatedButton
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import coil.compose.AsyncImage
import coil.request.ImageRequest
import com.photomanager.photomanager.R
import com.photomanager.photomanager.main.components.ImagePicker
import timber.log.Timber

@Composable
fun HomeScreen(onClicked: (Uri) -> Unit) {
    val vm = hiltViewModel<HomeViewModel>()
    var pickerShown by remember { mutableStateOf(false) }
    LaunchedEffect(Unit) {
        vm.action.collect {
            when (it) {
                is HomeAction.MarkDone -> onClicked(it.uri)
            }
        }
    }

    val state by vm.state.collectAsState()
    Column(modifier = Modifier.fillMaxSize(), horizontalAlignment = Alignment.CenterHorizontally) {
        ElevatedButton(onClick = { pickerShown = !pickerShown }) {
            Text(text = stringResource(id = R.string.open_picker))
        }

        if (pickerShown) {
            ImagePicker { uris ->
                if (!uris.isNullOrEmpty()) {
                    vm.dispatchEvent(HomeEvent.OnImagesPicked(uris))
                }
                pickerShown = false
            }
        }

        PhotoGrid(state.collections) {vm.dispatchEvent(HomeEvent.OnImageClicked(it))}
    }
}

@Composable
private fun PhotoGrid(imagesForGrid: List<Uri>, onImageClicked: (Uri) -> Unit = {}) {
    LazyVerticalGrid(modifier = Modifier.fillMaxSize(), columns = GridCells.Fixed(3)) {
        items(imagesForGrid.size) { index ->
            val uri = imagesForGrid[index]
            AsyncImage(
                modifier = Modifier
                    .fillMaxWidth()
                    .height(200.dp).clickable {
                        Timber.d("Clicked on image $uri")
                        onImageClicked(uri)
                    },
                model = ImageRequest.Builder(LocalContext.current).data(uri).build(),
                contentDescription = stringResource(id = R.string.grid_content_description, index)
            )
        }
    }
}